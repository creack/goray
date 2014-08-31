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
		a = (v.x*v.x + v.y*v.y - v.z*v.z) / (r * r)
		b = (2 * (v.x*float64(eye.x) + v.y*float64(eye.y) - v.z*float64(eye.z))) / (r * r)
		c = float64(eye.x*eye.x+eye.y*eye.y-eye.z*eye.z) / (r * r)
	)
	return SecondDegree(a, b, c)
}
