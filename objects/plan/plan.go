package plan

import (
	"image/color"

	"github.com/creack/goray/objects"
)

func init() {
	objects.RegisterObject("plan", NewPlan)
}

// Plan is the object's implemetation for a Plan.
type Plan struct {
	color color.Color
	z     int
}

// NewPlan instanciate the Plan object.
func NewPlan(obj objects.ObjectConfig) (objects.Object, error) {
	return (&Plan{}).Parse(obj)
}

// Color returns the Object's color
func (p *Plan) Color() color.Color {
	return p.color
}

// Parse populates the Plan's values from the given configuration object.
// If the plan is nil, instanciate it.
func (p *Plan) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if p == nil {
		p = &Plan{}
	}
	p.z = obj.Position.Z
	p.color = obj.Color
	return p, nil
}

// Intersect calculates the distance between the eye and the Object.
func (p *Plan) Intersect(v objects.Vector, eye objects.Point) float64 {
	if v.Z == 0 {
		return 0
	}
	return -float64(eye.Z+p.z) / v.Z
}
