package rt

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

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

func (rt *RT) Compute(eye objects.Point, objs []objects.Object) {
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
