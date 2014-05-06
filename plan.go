package main

import (
	"image/color"
	"strconv"
)

func init() {
	objectList["plan"] = NewPlan
}

type Plan struct {
	color color.Color
	z     int
}

func NewPlan(values map[string]string) (Object, error) {
	return (&Plan{}).Parse(values)
}

func (p *Plan) Color() color.Color {
	return p.color
}

func (p *Plan) Parse(values map[string]string) (Object, error) {
	if p == nil {
		p = &Plan{}
	}
	if z, err := strconv.Atoi(values["z"]); err == nil {
		p.z = z
	}
	color, err := DecodeColor(values["color"])
	if err != nil {
		return nil, err
	}
	p.color = color
	return p, nil
}

func (p *Plan) Intersect(v *Vector, eye *Point) float64 {
	if v.z == 0 {
		return 0
	}
	return -float64(eye.z+p.z) / v.z
}
