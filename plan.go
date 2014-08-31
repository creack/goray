package main

import "image/color"

func init() {
	objectList["plan"] = NewPlan
}

type Plan struct {
	color color.Color
	z     int
}

func NewPlan(obj ObjectConfig) (Object, error) {
	return (&Plan{}).Parse(obj)
}

func (p *Plan) Color() color.Color {
	return p.color
}

func (p *Plan) Parse(obj ObjectConfig) (Object, error) {
	if p == nil {
		p = &Plan{}
	}
	p.z = obj.Z
	p.color = RgbIntToColor(uint32(obj.Color))
	return p, nil
}

func (p *Plan) Intersect(v *Vector, eye *Point) float64 {
	if v.z == 0 {
		return 0
	}
	return -float64(eye.z+p.z) / v.z
}
