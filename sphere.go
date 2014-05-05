package main

import (
	"image/color"
	"math"
)

type Sphere struct {
	color          color.Color
	k              float64
	R              int
	position       *Point
	a, b, c, delta float64
}

func (s *Sphere) Color() color.Color {
	return s.color
}

func (s *Sphere) Intersect(v *Vector, eye *Point) float64 {
	eye.x -= s.position.x
	eye.y -= s.position.y
	eye.z -= s.position.z

	s.a = v.x*v.x + v.y*v.y + v.z*v.z
	s.b = 2*float64(eye.x)*v.x + float64(eye.y)*v.y + float64(eye.z)*v.z
	s.c = float64(eye.x*eye.x - s.R*s.R)
	// delta = b^2 - 4ac
	s.delta = s.b*s.b - 4*s.a*s.c
	s.k = 0

	eye.x += s.position.x
	eye.y += s.position.y
	eye.z += s.position.z

	if s.delta < 0 {
		return 0
	}
	// Two solution: (-b + sqrt(delta)) / 2a and (-b - sqrl(delta)) / 2a
	k1 := (-s.b + math.Sqrt(s.delta)) / (2 * s.a)
	k2 := (-s.b - math.Sqrt(s.delta)) / (2 * s.a)
	if k1 > 0 && k1 < k2 {
		s.k = k1
		return k1
	}
	s.k = k2
	return k2
}
