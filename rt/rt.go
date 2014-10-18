package rt

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

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
}

type workReponse struct {
	x, y int
	c    color.Color
}

// Compute .
func (rt *RT) ComputeChanResult(eye objects.Point, objs []objects.Object) {
	rc := make(chan []workReponse, rt.Height)
	c := make(chan int, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		go func() {
			wg.Add(1)
			for {
				y, ok := <-c
				if !ok {
					wg.Done()
					return
				}
				resp := make([]workReponse, rt.Width)
				for x := 0; x < rt.Width; x++ {
					if rt.Verbose && x == 0 && y%10 == 0 {
						fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.Height))*100+1))
					}
					resp[x].x, resp[x].y, resp[x].c = x, y, rt.calc(x, y, eye, objs)
				}
				rc <- resp
			}
		}()
	}
	for y := 0; y < rt.Height; y++ {
		c <- y
	}
	close(c)
	go func() {
		wg.Wait()
		close(rc)
	}()

	for ress := range rc {
		for _, res := range ress {
			rt.Img.Set(res.x, res.y, res.c)
		}
	}

	// rt.Img.Set(x, y,
	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}

// Compute .
func (rt *RT) Compute(eye objects.Point, objs []objects.Object) {
	c := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		go func() {
			wg.Add(1)
			for {
				y, ok := <-c
				if !ok {
					wg.Done()
					return
				}
				for x := 0; x < rt.Width; x++ {
					if rt.Verbose && x == 0 && y%10 == 0 {
						fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.Height))*100+1))
					}
					rt.Img.Set(x, y, rt.calc(x, y, eye, objs))
				}
			}
		}()
	}
	for y := 0; y < rt.Height; y++ {
		c <- y
	}
	// for i, total := 0, rt.Width*rt.Height; i < total; i++ {
	// 	wq.c <- i
	// }
	close(c)
	wg.Wait()

	// rt.Img.Set(x, y,
	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}

func (rt *RT) ComputeSemaphone(eye objects.Point, objs []objects.Object) {
	max := runtime.NumCPU()
	c := make(chan struct{}, max)
	for i := 0; i < max; i++ {
		c <- struct{}{}
	}
	for y := 0; y < rt.Height; y++ {
		<-c
		go func(y int) {
			for x := 0; x < rt.Width; x++ {
				if rt.Verbose && x == 0 && y%10 == 0 {
					fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.Height))*100+1))
				}
				rt.Img.Set(x, y, rt.calc(x, y, eye, objs))
			}
			c <- struct{}{}
		}(y)
	}
	for i := len(c); i < max; i++ {
		<-c
	}
	close(c)
	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}

func (rt *RT) ComputeOrigin(eye objects.Point, objs []objects.Object) {
	var (
		x int
		y int
	)

	for i, total := 0, rt.Width*rt.Height; i < total; i++ {
		x = i % rt.Width
		y = i / rt.Width
		if rt.Verbose && x == 0 && y%10 == 0 {
			fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.Height))*100+1))
		}
		rt.Img.Set(x, y, rt.calc(x, y, eye, objs))
	}
	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}
