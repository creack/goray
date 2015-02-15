package rt

import (
	"fmt"
	"image"
	"image/color"

	"github.com/creack/goray/objects"
	// Load all Object modules
	_ "github.com/creack/goray/objects/all"
)

// RT represent the Scene.
// Hold the bondaries as well as the processed image buffer.
// Also holds various extra metadata.
type RT struct {
	Img     *image.RGBA
	Width   int
	Height  int
	Verbose bool
}

// Eye represent the scene's Camera
// It is a point with a vector.
type Eye struct {
	Position objects.Point
	Rotation objects.Vector
}

// SceneConfig represent the configuration for a Scene.
// This contains Scene sizes, the Camera and the Object list.
type SceneConfig struct {
	Height  int
	Width   int
	Eye     *Eye
	Objects []objects.Object
}

// NewRT instantiates a new Scene.
func NewRT(w, h int) *RT {
	return &RT{
		Img:    image.NewRGBA(image.Rect(0, 0, w, h)),
		Width:  w,
		Height: h,
	}
}

// calc iterates through the whole Object list and
// returns the closest point's Color,
func (rt *RT) calc(x, y int, eye objects.Point, objs []objects.Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = objects.Vector{
			X: 100,
			Y: float64(rt.Width/2 - x),
			Z: float64(rt.Height/2 - y),
		}
	)
	for _, obj := range objs {
		// If k == -1, it is our first pass, so if we have a solution, keep it.
		// After that, we check that the solution is smaller than the one we have.
		if tmp := obj.Intersect(v, eye); tmp > 0 && (k == -1 || tmp < k) {
			k = tmp
			col = obj.Color()
		}
	}
	return col
}

// Compute process the Scene with the given Camera (Eye)
// and the given Object list.
func (rt *RT) Compute(eye objects.Point, objs []objects.Object) {
	var (
		x int
		y int
	)

	for i, total := 0, rt.Width*rt.Height; i < total; i++ {
		x = i % rt.Width
		y = i / rt.Width
		if rt.Verbose && x == 0 && y%10 == 0 {
			fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.Height))*100+1))
		}
		rt.Img.Set(x, y, rt.calc(x, y, eye, objs))
	}
	if rt.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}
