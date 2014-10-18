package main

import (
	"fmt"
	"image/color"
	"log"
	"runtime"
	"time"

	"github.com/creack/goray/cli"
	"github.com/creack/goray/objects"
	_ "github.com/creack/goray/parser/all"
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

func main() {
	// Process CLI flags
	cliConf, err := cli.Flags()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the scene file
	sceneConf, err := cliConf.Parser.Parser.Parse(cliConf.SceneFile)
	if err != nil {
		log.Fatal(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	// Process the image
	rtrace := rt.NewRT(sceneConf.Width, sceneConf.Height)
	rtrace.Verbose = cliConf.Verbose

	start := time.Now().UTC()
	rtrace.Compute(sceneConf.Eye.Position, sceneConf.Objects)
	fmt.Printf("%0.6fs\n", time.Since(start).Seconds())

	// Render the image
	if err := cliConf.Renderer.Renderer.Render(rtrace, sceneConf.Eye, sceneConf.Objects); err != nil {
		log.Fatal(err)
	}
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
