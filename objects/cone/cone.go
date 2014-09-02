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

type Cone struct {
	color    color.Color
	R        int
	position objects.Point
}

func NewCone(obj objects.ObjectConfig) (objects.Object, error) {
	return (&Cone{}).Parse(obj)
}

func (cc *Cone) Color() color.Color {
	return cc.color
}

func (cc *Cone) Parse(obj objects.ObjectConfig) (objects.Object, error) {
	if cc == nil {
		cc = &Cone{}
	}
	cc.position, cc.R, cc.color = obj.Position, obj.R, obj.Color
	return cc, nil
}

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
