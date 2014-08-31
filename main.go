package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
)

type NewObjectFct func(ObjectConfig) (Object, error)

var objectList = map[string]NewObjectFct{}

type Object interface {
	Color() color.Color
	Intersect(v *Vector, eye *Point) float64
	Parse(values ObjectConfig) (Object, error)
}

type Vector struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

type RT struct {
	img     *image.RGBA
	width   int
	height  int
	verbose bool
}

func NewRT(x, y int) *RT {
	return &RT{
		img:    image.NewRGBA(image.Rect(0, 0, x, y)),
		width:  x,
		height: y,
	}
}

func (rt *RT) calc(x, y int, eye *Point, objs []Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = &Vector{
			X: 100,
			Y: float64(rt.width/2 - x),
			Z: float64(rt.height/2 - y),
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

func (rt *RT) fillImage(eye *Point, objs []Object) {
	var (
		x int
		y int
	)

	for i, total := 0, rt.width*rt.height; i < total; i++ {
		x = i % rt.width
		y = i / rt.width
		if rt.verbose && x == 0 && y%10 == 0 {
			fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.height))*100+1))
		}
		rt.img.Set(x, y, rt.calc(x, y, eye, objs))
	}
	if rt.verbose {
		fmt.Println("\rProcessing: 100%")
	}
}

func main() {
	eye, objs, err := parseConfigYaml("rt.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	rt := NewRT(800, 600)
	rt.verbose = true
	rt.fillImage(eye, objs)

	if true {
		w, err := x11.NewWindow()
		if err != nil {
			fmt.Println(err)
			return
		}
		fct := func() {
			rt.fillImage(eye, objs)
			draw.Draw(w.Screen(), w.Screen().Bounds(), rt.img, image.ZP, draw.Src)
			w.FlushImage()
		}
		fct()
		for e := range w.EventChan() {
			switch e := e.(type) {
			case ui.KeyEvent:
				switch KeyListInt[e.Key] {
				case " ", "<esc>", "\n", "q":
					return
				case "<up>":
					eye.X += 10
				case "<down>":
					eye.X -= 10
				case "<left>":
					eye.Y += 10
				case "<right>":
					eye.Y -= 10
				case "a":
					eye.Z += 10
				case "z":
					eye.Z -= 10
				}
				fct()
			}
		}
	}
}
