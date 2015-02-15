package cone

import (
	"image/color"
	"math"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/utils"
)

func init() {
	objects.RegisterObject("cone", NewCone)
}

// Cone is the Object's implemetation for a Cone.
type Cone struct {
	color    color.Color
	R        int
	position objects.Point
}

// NewCone instanciate the Cone object.
func NewCone(obj objects.ObjectConfig) (objects.Object, error) {
	return (&Cone{}).Parse(obj)
}

// Color returns the Object's color
func (cc *Cone) Color() color.Color {
	return cc.color
}

// Parse populates the Cone's values from the given configuration object.
// If the Cone is nil, instantiate it.
func (cc *Cone) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if cc == nil {
		cc = &Cone{}
	}
	cc.position, cc.R, cc.color = obj.Position, obj.R, obj.Color
	return cc, nil
}

// Intersect calculates the distance between the eye and the Object.
func (cc *Cone) Intersect(v objects.Vector, eye objects.Point) float64 {
	eye.Sub(cc.position)
	defer eye.Add(cc.position)

	var (
		r = math.Tan(float64(cc.R) / math.Pi * 180)
		a = (v.X*v.X + v.Y*v.Y - v.Z*v.Z) / (r * r)
		b = (2 * (v.X*float64(eye.X) + v.Y*float64(eye.Y) - v.Z*float64(eye.Z))) / (r * r)
		c = float64(eye.X*eye.X+eye.Y*eye.Y-eye.Z*eye.Z) / (r * r)
	)
	return utils.SecondDegree(a, b, c)
}
