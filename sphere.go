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
		a = v.X*v.X + v.Y*v.Y + v.Z*v.Z
		b = 2*float64(eye.X)*v.X + float64(eye.Y)*v.Y + float64(eye.Z)*v.Z
		c = float64(eye.X*eye.X - s.R*s.R)
	)
	return SecondDegree(a, b, c)
}
