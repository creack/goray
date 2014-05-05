package main

import (
	"image/color"
)

type Plan struct {
	color color.Color
	k     float64
	z     int
}

func (p *Plan) Color() color.Color {
	return p.color
}

func (p *Plan) Intersect(v *Vector, eye *Point) float64 {
	p.k = 0
	if v.z == 0 {
		return 0
	}
	p.k = -float64(eye.z+p.z) / v.z
	return p.k
}
