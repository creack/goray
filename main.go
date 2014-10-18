package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/bitly/go-nsq"
	"github.com/bitly/nsq/util"
	"github.com/creack/goray/cli"
	"github.com/creack/goray/objects"
	_ "github.com/creack/goray/parser/all"
	"github.com/creack/goray/parser/yaml"
	_ "github.com/creack/goray/render/all"
	"github.com/creack/goray/rt"
)

// WorkRequest .
type WorkRequest struct {
	x, y int
	eye  objects.Point
	objs []objects.Object
	fct  func(x, y int, eye objects.Point, objs []objects.Object) color.Color
}

var workerQueue chan chan WorkRequest

func a(nWorkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	workerQueue = make(chan chan WorkRequest, nWorkers)

	// Now, create all of our workers.
	for i := 0; i < nWorkers; i++ {
		go func() {
			c := make(chan WorkRequest)
			for {
				// Signal that we are ready for a new task
				workerQueue <- c
				// Wait for task dispatch
				wr := <-c
				wr.fct(wr.x, wr.y, wr.eye, wr.objs)
			}
		}()
	}
}

func setMaxProc() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	fmt.Printf("MaxProcs: %d\n", runtime.GOMAXPROCS(0))
}

var (
	tcpAddr  = "localhost:4150"
	httpAddr = "http://localhost:4151"
)

func nsqSubscribe(topicName, channelName string, hdlr nsq.HandlerFunc) error {
	println("Subscribe on", topicName, channelName)
	cfg := nsq.NewConfig()
	cfg.UserAgent = fmt.Sprintf("nsq_pubsub/%s go-nsq/%s", util.BINARY_VERSION, nsq.VERSION)
	cfg.MaxInFlight = 8
	r, err := nsq.NewConsumer(topicName, channelName, cfg)
	if err != nil {
		return err
	}
	r.AddHandler(hdlr)

	if err := r.ConnectToNSQD(tcpAddr); err != nil {
		return err
	}
	//	<-r.StopChan
	return nil
}

func nsqPublish(topicName string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := http.Post(httpAddr+"/pub?topic="+topicName, "application/json", bytes.NewReader(d))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error publishing: %d\n", resp.StatusCode)
	}
	return nil
}

func nsqMPublish(topicName string, data []interface{}) error {
	var body string
	for _, elem := range data {
		d, err := json.Marshal(elem)
		if err != nil {
			return err
		}
		body += string(d) + "\n"
	}
	body = strings.TrimRight(body, "\n")
	resp, err := http.Post(httpAddr+"/mpub?topic="+topicName, "application/json", strings.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error publishing: %d\n", resp.StatusCode)
	}
	return nil
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
}

func randomBuf(len int) ([]byte, error) {
	buf := make([]byte, len/2+len%2)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(buf)), nil
}

type job struct {
	ID     string           `json:"id"`
	Objs   []objects.Object `json:"-"`
	Eye    *rt.Eye          `json:"-"`
	Height int              `json:"-"`
	Width  int              `json:"-"`
	Scene  string           `json:"scene"`
}

type workRequest struct {
	Y int `json:"y"`
}

type workReponse struct {
	X int      `json:"x"`
	Y int      `json:"y"`
	C []uint32 `json:"c"`
}

func (j *job) process(msg *nsq.Message) error {
	// rand, _ := randomBuf(6)
	// println("Start process", j.ID, string(rand))
	// defer println("Finish Process", j.ID, string(rand))

	var data workRequest
	if err := json.Unmarshal(msg.Body, &data); err != nil {
		println("------------------> Error Unmarshal work request")
		return err
	}
	resp := make([]workReponse, j.Width)
	var c color.Color
	for x := 0; x < j.Width; x++ {
		resp[x].X, resp[x].Y, c = x, data.Y, rt.Calc(j.Height, j.Width, x, data.Y, j.Eye.Position, j.Objs)
		// fmt.Println(j.Height, j.Width, data.Y, j.Eye.Position, len(j.Objs))
		// fmt.Println(c.RGBA())
		r, g, b, a := c.RGBA()
		resp[x].C = []uint32{r, g, b, a}
	}
	if err := nsqPublish("job-"+j.ID+"-response", resp); err != nil {
		println("------------------> Error publish job")
		return err
	}
	return nil
}

func newWork(msg *nsq.Message) error {
	println("--- New Job Created ---")
	defer println("--- New Job Created --- end --")

	var j job
	if err := json.Unmarshal(msg.Body, &j); err != nil {
		println("------------------> Error unnarshal new job")
		return err
	}
	sceneConf, err := yaml.ParseBuffer([]byte(j.Scene))
	if err != nil {
		println("------------------> Error parse scene")
		return err
	}
	j.Eye = sceneConf.Eye
	j.Objs = sceneConf.Objects
	j.Height = sceneConf.Height
	j.Width = sceneConf.Width

	if err := nsqSubscribe("job-"+j.ID, "jobchan", j.process); err != nil {
		return err
	}
	return nil
}

func startWorker() {
	println("Subscribe for new Jobs")
	defer println("Subscribe for new Jobs -- end")

	rand, err := randomBuf(12)
	if err != nil {
		log.Fatal(err)
	}
	if err := nsqSubscribe("newwork", "newwork-"+string(rand), newWork); err != nil {
		log.Fatal(err)
	}
}

func compute(sceneFile string, rtt *rt.RT, eye objects.Point, objs []objects.Object) {
	println("Start compute()")
	defer println("End compute()")
	// Generate job ID
	rand, err := randomBuf(12)
	if err != nil {
		log.Fatal(err)
	}

	// Publish new job
	j := &job{
		ID:    string(rand),
		Scene: sceneFile,
	}
	println("pre pubblish newwork")
	if err := nsqPublish("newwork", j); err != nil {
		log.Fatal(err)
	}
	println("post pubblish newwork")

	// Receive all response
	rc := make(chan []workReponse, rtt.Height)
	var i int32
	nsqSubscribe("job-"+j.ID+"-response", "jobresponse", func(msg *nsq.Message) error {
		var resp []workReponse
		if err := json.Unmarshal(msg.Body, &resp); err != nil {
			log.Fatal(err)
		}
		rc <- resp
		if atomic.AddInt32(&i, 1) == int32(rtt.Height) {
			close(rc)
		}
		return nil
	})

	println("pre pubblish work load")

	body := make([]interface{}, rtt.Height)
	// Send all work request
	for y := 0; y < rtt.Height; y++ {
		body[y] = &workRequest{Y: y}
	}

	if err := nsqMPublish("job-"+j.ID, body); err != nil {
		log.Fatal(err)
	}

	println("post pubblish work load -- subscribe to response")

	done := make(chan struct{})
	go func() {
		// Draw image
		for ress := range rc {
			for _, res := range ress {
				c := color.RGBA{
					R: uint8(res.C[0]),
					G: uint8(res.C[1]),
					B: uint8(res.C[2]),
					A: uint8(res.C[3]),
				}
				rtt.Img.Set(res.X, res.Y, c)
			}
		}
		close(done)
	}()

	timeout := time.After(600 * time.Second)
	ticker := time.Tick(time.Second)
	defer println(i, rtt.Height)
	for {
		select {
		case <-done:
			return
		case <-ticker:
			ii := atomic.LoadInt32(&i)
			println(ii, rtt.Height)
		case <-timeout:
			println("timeout")
			return
		}
	}
}

func main() {
	setMaxProc()

	remote := flag.Bool("remote", false, "Remote")
	docker := flag.Bool("docker", false, "Docker")

	// Process CLI flags
	cliConf, err := cli.Flags()
	if err != nil {
		log.Fatal(err)
	}
	if *remote {
		tcpAddr = "nsq.1.apollo.cloudburrito.com:21"
		httpAddr = "http://nsq.1.apollo.cloudburrito.com:23"
	}
	if *docker {
		tcpAddr = os.Getenv("DOCKER_IP") + ":21"
		httpAddr = "http://" + os.Getenv("DOCKER_IP") + ":23"
	}

	if cliConf.Worker {
		startWorker()
	}

	if cliConf.SceneFile == "" {
		fmt.Printf("Worker only\n")
		<-make(chan struct{})
	}

	// Parse the scene file
	sceneConf, err := cliConf.Parser.Parser.Parse(cliConf.SceneFile)
	if err != nil {
		log.Fatal(err)
	}

	scene, err := ioutil.ReadFile(cliConf.SceneFile)
	if err != nil {
		log.Fatal(err)
	}

	// Process the image
	rtrace := rt.NewRT(sceneConf.Width, sceneConf.Height)
	rtrace.Verbose = cliConf.Verbose

	start := time.Now().UTC()
	compute(string(scene), rtrace, sceneConf.Eye.Position, sceneConf.Objects)
	fmt.Printf("%0.6fs\n", time.Since(start).Seconds())

	// Render the image
	if err := cliConf.Renderer.Renderer.Render(rtrace, sceneConf.Eye, sceneConf.Objects); err != nil {
		log.Fatal(err)
	}

	return
}

/*
	for name, fct := range map[string]func(eye objects.Point, objs []objects.Object){
		//		"origin":         rtrace.ComputeOrigin,
		"sema":           rtrace.ComputeSemaphone,
		"workerDirect":   rtrace.ComputeWorkerDirect,
		"workerIndirect": rtrace.Compute,
	} {
		start := time.Now().UTC()
		fct(sceneConf.Eye.Position, sceneConf.Objects)
		fmt.Printf("%s:\t%0.6fms\n", name, time.Since(start).Seconds()*100)
	}
*/
