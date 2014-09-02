package rt

import (
	"fmt"
	"image"
	"image/color"

	"github.com/creack/goray/objects"
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

func NewRT(x, y int) *RT {
	return &RT{
		Img:    image.NewRGBA(image.Rect(0, 0, x, y)),
		Width:  x,
		Height: y,
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

func (rt *RT) FillImage(eye objects.Point, objs []objects.Object) {
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
