package sphere

import (
	"image/color"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/utils"
)

func init() {
	objects.RegisterObject("sphere", NewSphere)
}

// Sphere is the object's implemetation for a Sphere.
type Sphere struct {
	color    color.Color
	R        int
	position objects.Point
}

// NewSphere instanciate the Sphere object.
func NewSphere(obj objects.ObjectConfig) (objects.Object, error) {
	return (&Sphere{}).Parse(obj)
}

// Color returns the Object's color.
func (s *Sphere) Color() color.Color {
	return s.color
}

// Parse populates the Sphere's values from the given configuration object.
// If the Sphere is nil, instantiate it.
func (s *Sphere) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if s == nil {
		s = &Sphere{}
	}
	s.position, s.R, s.color = obj.Position, obj.R, obj.Color
	return s, nil
}

// Intersect calculates the distance between the eye and the Object.
func (s *Sphere) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(s.position)
	defer eye.Add(s.position)

	var (
		a = v.X*v.X + v.Y*v.Y + v.Z*v.Z
		b = 2*float64(eye.X)*v.X + float64(eye.Y)*v.Y + float64(eye.Z)*v.Z
		c = float64(eye.X*eye.X - s.R*s.R)
	)
	return utils.SecondDegree(a, b, c)
}
