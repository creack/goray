package x11

import (
	"image"
	"image/draw"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	"github.com/creack/goray/objects"
	"github.com/creack/goray/render"
	"github.com/creack/goray/rt"
	"github.com/creack/goray/utils"
)

type X11Renderer struct {
}

func init() {
	render.RegisterRenderer("x11", &X11Renderer{})
}

func (xr *X11Renderer) Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error {
	w, err := x11.NewWindow()
	if err != nil {
		return err
	}
	fct := func() {
		rt.FillImage(eye.Position, objs)
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

func (xr *X11Renderer) Flags() {
}
