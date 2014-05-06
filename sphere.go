package main

import (
	"image/color"
	"strconv"
)

func init() {
	objectList["sphere"] = NewSphere
}

type Sphere struct {
	color    color.Color
	R        int
	position *Point
}

func NewSphere(values map[string]string) (Object, error) {
	return (&Sphere{}).Parse(values)
}

func (s *Sphere) Color() color.Color {
	return s.color
}

func (s *Sphere) Parse(values map[string]string) (Object, error) {
	if s == nil {
		s = &Sphere{}
	}
	position, err := s.position.Parse(values)
	if err != nil {
		return nil, err
	}
	r, err := strconv.Atoi(values["R"])
	if err != nil {
		return nil, err
	}
	color, err := DecodeColor(values["color"])
	if err != nil {
		return nil, err
	}
	s.position, s.R, s.color = position, r, color
	return s, nil
}

func (s *Sphere) Intersect(v *Vector, eye *Point) float64 {
	eye.x -= s.position.x
	eye.y -= s.position.y
	eye.z -= s.position.z

	var (
		a = v.x*v.x + v.y*v.y + v.z*v.z
		b = 2*float64(eye.x)*v.x + float64(eye.y)*v.y + float64(eye.z)*v.z
		c = float64(eye.x*eye.x - s.R*s.R)
	)

	eye.x += s.position.x
	eye.y += s.position.y
	eye.z += s.position.z
	return SecondDegree(a, b, c)
}
