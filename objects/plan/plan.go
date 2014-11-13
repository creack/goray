package plan

import (
	"image/color"

	"github.com/creack/goray/objects"
)

func init() {
	objects.RegisterObject("plan", NewPlan)
}

type Plan struct {
	color color.Color
	z     int
}

func NewPlan(obj objects.ObjectConfig) (objects.Object, error) {
	return (&Plan{}).Parse(obj)
}

func (p *Plan) Color() color.Color {
	return p.color
}

func (p *Plan) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if p == nil {
		p = &Plan{}
	}
	p.z = obj.Position.Z
	p.color = obj.Color
	return p, nil
}

func (p *Plan) Intersect(v objects.Vector, eye objects.Point) float64 {
	if v.Z == 0 {
		return 0
	}
	return -float64(eye.Z-p.z) / v.Z
}
