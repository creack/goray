package objects

import "image/color"

func RegisterObject(name string, f NewObjectFct) {
	ObjectList[name] = f
}

type NewObjectFct func(ObjectConfig) (Object, error)

var ObjectList = map[string]NewObjectFct{}

type ObjectConfig struct {
	Type     string `yaml:"type"`
	Position Point  `yaml:"position"`
	Rotation Vector `yaml:"rotation"`
	Color    uint32 `yaml:"color"`
	R        int    `yaml:"R"`
}

type Point struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
	Z int `yaml:"z"`
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
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

type Object interface {
	Color() color.Color
	Intersect(v Vector, eye Point) float64
	Parse(values ObjectConfig) (Object, error)
}
