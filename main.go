package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
)

var objectList = map[string]func(map[string]string) (Object, error){}

type Object interface {
	Color() color.Color
	Intersect(v *Vector, eye *Point) float64
	Parse(values map[string]string) (Object, error)
}

type Point struct {
	x, y, z int
}

func (p *Point) Parse(values map[string]string) (*Point, error) {
	if p == nil {
		p = &Point{}
	}
	xStr, yStr, zStr := values["x"], values["y"], values["z"]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return nil, err
	}
	z, err := strconv.Atoi(zStr)
	if err != nil {
		return nil, err
	}
	p.x, p.y, p.z = x, y, z
	return p, nil
}

type Vector struct {
	x, y, z float64
}

type RT struct {
	img    *image.RGBA
	width  int
	height int
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
			x: 100,
			y: float64(rt.width/2 - x),
			z: float64(rt.height/2 - y),
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
		rt.img.Set(x, y, rt.calc(x, y, eye, objs))
	}
}

func main() {
	eye, objs, err := parseConfig("rt.config")
	if err != nil {
		fmt.Println(err)
		return
	}
	rt := NewRT(800, 600)
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
				eye.x += 10
			case "<down>":
				eye.x -= 10
			case "<left>":
				eye.y += 10
			case "<right>":
				eye.y -= 10
			case "a":
				eye.z += 10
			case "z":
				eye.z -= 10
			}
			fct()
		}
	}
}
