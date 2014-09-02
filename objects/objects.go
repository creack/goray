package objects

import "image/color"

type ObjectConfig struct {
	Type     string
	Position Point
	Rotation Vector
	Color    color.Color
	R        int
}

func RegisterObject(name string, f NewObjectFct) {
	ObjectList[name] = f
}

type NewObjectFct func(ObjectConfig) (Object, error)

var ObjectList = map[string]NewObjectFct{}

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) Add(p2 Point) {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
}

func (p *Point) Sub(p2 Point) {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
}

type Vector struct {
	X float64
	Y float64
	Z float64
}

type Object interface {
	Color() color.Color
	Intersect(v Vector, eye Point) float64
	Parse(values ObjectConfig) (Object, error)
}
