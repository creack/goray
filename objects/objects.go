package objects

import "image/color"

// ObjectConfig represent an object configuraton.
type ObjectConfig struct {
	Type     string
	Position Point
	Rotation Vector
	Color    color.Color
	R        int
}

// RegisterObject registers an object.
// Used by the underlying implementations.
func RegisterObject(name string, f NewObjectFct) {
	ObjectList[name] = f
}

// NewObjectFct is a typedef on an Object constructor.
type NewObjectFct func(ObjectConfig) (Object, error)

// ObjectList hold the available objects
// (registered by underlying implementation)
var ObjectList = map[string]NewObjectFct{}

// Point represents a point
type Point struct {
	X int
	Y int
	Z int
}

// Add adds the given point to the current one.
func (p *Point) Add(p2 Point) {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
}

// Sub substract the given point from the currentone.
func (p *Point) Sub(p2 Point) {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
}

// Vector represents a vector.
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Object is the Object's interface.
type Object interface {
	Color() color.Color
	Intersect(v Vector, eye Point) float64
	Parse(values ObjectConfig) (Object, error)
}
