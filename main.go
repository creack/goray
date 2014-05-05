package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
)

type Object interface {
	Color() color.Color
	Intersect(v *Vector, eye *Point) float64
}

type Point struct {
	x, y, z int
}

type Vector struct {
	x, y, z float64
}

type RT struct {
	img *image.RGBA
}

func (rt *RT) calc(X, Y, x, y int, eye *Point, objs []Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = &Vector{
			x: 100,
			y: float64(X/2 - x),
			z: float64(Y/2 - y),
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
		x, y int
		X    = rt.img.Bounds().Dx()
		Y    = rt.img.Bounds().Dy()
	)

	for i := 0; i < X*Y; i++ {
		x = i % X
		y = i / X
		rt.img.Set(x, y, rt.calc(X, Y, x, y, eye, objs))
	}
}

func main() {
	w, err := x11.NewWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	rt := &RT{
		img: image.NewRGBA(image.Rect(0, 0, 800, 600)),
	}
	eye := &Point{
		x: -300,
		y: 0,
		z: 100,
	}

	objs := []Object{
		&Plan{
			color: rgbToColor(0xFF0000),
			z:     0,
		},
		&Sphere{
			position: &Point{
				x: 100,
				y: 0,
				z: 0,
			},
			R:     100,
			color: rgbToColor(0xFFFF00),
		},
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
			switch e.Key {
			case KeyList[" "], KeyList["\\e"], KeyList["\n"]:
				return
			case KeyList["<up>"]:
				eye.x += 10
			case KeyList["<down>"]:
				eye.x -= 10
			case KeyList["<left>"]:
				eye.y += 10
			case KeyList["<right>"]:
				eye.y -= 10
			case 'a':
				eye.z += 10
			case 'z':
				eye.z -= 10
			default:
				fmt.Printf("%#v\n", e.Key)
			}
			fct()
		}
	}
}
