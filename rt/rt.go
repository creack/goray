package rt

import (
	"fmt"
	"image"
	"image/color"

	"sync"

	"github.com/creack/goray/objects"
	_ "github.com/creack/goray/objects/all"
)

type RT struct {
	Img     *image.RGBA
	Width   int
	Height  int
	Verbose bool
}

type Eye struct {
	Position objects.Point
	Rotation objects.Vector
}

type SceneConfig struct {
	Height  int
	Width   int
	Eye     *Eye
	Objects []objects.Object
}

func NewRT(w, h int) *RT {
	return &RT{
		Img:    image.NewRGBA(image.Rect(0, 0, w, h)),
		Width:  w,
		Height: h,
	}
}

func (rt *RT) calc(x, y int, eye objects.Point, objs []objects.Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = objects.Vector{
			X: 100,
			Y: float64(rt.Width/2 - x),
			Z: float64(rt.Height/2 - y),
		}
	)
	for _, obj := range objs {
		if tmp := obj.Intersect(v, eye); tmp > 0 && (k == -1 || tmp < k) {
			k = tmp
			col = obj.Color()
		}
	}
	return col
}

type workQueue struct {
	eye  objects.Point
	objs []objects.Object
	fct  func(x, y int, eye objects.Point, objs []objects.Object) color.Color
	h, w int
	c    chan int
	wg   sync.WaitGroup
	rc   chan []workReponse
}

type workReponse struct {
	x, y int
	c    color.Color
}

func newWorker(wq *workQueue) {
	wq.wg.Add(1)
	for {
		y, ok := <-wq.c
		if !ok {
			wq.wg.Done()
			return
		}
		resp := make([]workReponse, wq.w)
		for x := 0; x < wq.w; x++ {
			resp[x].x, resp[x].y, resp[x].c = x, y, wq.fct(x, y, wq.eye, wq.objs)
		}
		wq.rc <- resp
	}
}

func (wq *workQueue) startWorkers(nWorkers int) {
	for i := 0; i < nWorkers; i++ {
		go newWorker(wq)
	}
}

// Finish ends the queue. Non Blocking. Expect user to block on response chan.
func (wq *workQueue) Finish() {
	go func() {
		close(wq.c)
		wq.wg.Wait()
		close(wq.rc)
	}()
}

// Compute .
func (rt *RT) Compute(eye objects.Point, objs []objects.Object) {
	wq := &workQueue{
		rc:   make(chan []workReponse, rt.Height),
		c:    make(chan int, 100),
		w:    rt.Width,
		h:    rt.Height,
		fct:  rt.calc,
		objs: objs,
		eye:  eye,
	}
	wq.startWorkers(8)
	for y := 0; y < rt.Height; y++ {
		wq.c <- y
	}
	wq.Finish()

	for ress := range wq.rc {
		for _, res := range ress {
			rt.Img.Set(res.x, res.y, res.c)
		}
	}

	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}
