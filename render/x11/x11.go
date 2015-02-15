package x11

import (
	"flag"
	"image"
	"image/draw"
	"os"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	"github.com/creack/goray/objects"
	"github.com/creack/goray/render"
	"github.com/creack/goray/rt"
	"github.com/creack/goray/utils"
)

// Renderer represent the X11 renderer.
type Renderer struct {
	init bool
}

func init() {
	render.RegisterRenderer("x11", &Renderer{})
}

// Render renders the given scene (`rt`) with the given object list
// From the `eye` perspective.
// Renders on X11.
func (r *Renderer) Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error {
	w, err := x11.NewWindow()
	if err != nil {
		return err
	}
	fct := func() {
		rt.Compute(eye.Position, objs)
		draw.Draw(w.Screen(), w.Screen().Bounds(), rt.Img, image.ZP, draw.Src)
		w.FlushImage()
	}
	fct()
	for e := range w.EventChan() {
		switch e := e.(type) {
		case ui.KeyEvent:
			switch utils.KeyListInt[e.Key] {
			case " ", "<esc>", "\n", "q":
				return nil
			case "<up>":
				eye.Position.X += 10
			case "<down>":
				eye.Position.X -= 10
			case "<left>":
				eye.Position.Y += 10
			case "<right>":
				eye.Position.Y -= 10
			case "a":
				eye.Position.Z += 10
			case "z":
				eye.Position.Z -= 10
			}
			fct()
		}
	}
	return nil
}

// Flags extends the CLI with X11 specific flags.
func (r *Renderer) Flags() {
	if r.init {
		return
	}
	r.init = true

	display := os.Getenv("DISPLAY")
	if display == "" {
		display = ":0"
	}
	flag.StringVar(&display, "display", display, "Display to use")
}
