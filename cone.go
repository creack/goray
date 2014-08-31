package main

import (
	"image/color"
	"math"
)

func init() {
	objectList["cone"] = NewCone
}

type Cone struct {
	color    color.Color
	R        int
	position Point
}

func NewCone(obj ObjectConfig) (Object, error) {
	return (&Cone{}).Parse(obj)
}

func (cc *Cone) Color() color.Color {
	return cc.color
}

func (cc *Cone) Parse(obj ObjectConfig) (Object, error) {
	if cc == nil {
		cc = &Cone{}
	}
	position, err := cc.position.Parse(obj)
	if err != nil {
		return nil, err
	}
	color := RgbIntToColor(uint32(obj.Color))

	cc.position, cc.R, cc.color = position, obj.R, color
	return cc, nil
}

func (cc *Cone) Intersect(v *Vector, eye *Point) float64 {
	eye.Sub(cc.position)
	defer eye.Add(cc.position)

	var (
		r = math.Tan(float64(cc.R) / math.Pi * 180)
		a = (v.X*v.X + v.Y*v.Y - v.Z*v.Z) / (r * r)
		b = (2 * (v.X*float64(eye.X) + v.Y*float64(eye.Y) - v.Z*float64(eye.Z))) / (r * r)
		c = float64(eye.X*eye.X+eye.Y*eye.Y-eye.Z*eye.Z) / (r * r)
	)
	return SecondDegree(a, b, c)
}
