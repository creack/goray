package main

import "image/color"

func init() {
	objectList["sphere"] = NewSphere
}

type Sphere struct {
	color    color.Color
	R        int
	position Point
}

func NewSphere(obj ObjectConfig) (Object, error) {
	return (&Sphere{}).Parse(obj)
}

func (s *Sphere) Color() color.Color {
	return s.color
}

func (s *Sphere) Parse(obj ObjectConfig) (Object, error) {
	if s == nil {
		s = &Sphere{}
	}
	position, err := s.position.Parse(obj)
	if err != nil {
		return nil, err
	}
	color := RgbIntToColor(uint32(obj.Color))
	s.position, s.R, s.color = position, obj.R, color
	return s, nil
}

func (s *Sphere) Intersect(v *Vector, eye *Point) float64 {
	eye.Sub(s.position)
	defer eye.Add(s.position)

	var (
		a = v.x*v.x + v.y*v.y + v.z*v.z
		b = 2*float64(eye.x)*v.x + float64(eye.y)*v.y + float64(eye.z)*v.z
		c = float64(eye.x*eye.x - s.R*s.R)
	)
	return SecondDegree(a, b, c)
}
