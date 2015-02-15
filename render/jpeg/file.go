package jpeg

import (
	"flag"
	"image/jpeg"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/render"
	"github.com/creack/goray/scene"
)

func init() {
	render.RegisterRenderer("jpeg", &Renderer{})
}

// Renderer represent the JPEG renderer.
type Renderer struct {
	file string
}

// Render renders the given scene (`rt`) with the given object list
// From the `eye` perspective.
// Renders to a JPEG file.
func (r *Renderer) Render(s *scene.Scene, eye scene.Eye, objs []objects.Object) error {
	f, err := os.Create(r.file)
	if err != nil {
		return err
	}
	if err := jpeg.Encode(f, s.Img, nil); err != nil {
		return err
	}
	return nil
}

// Flags extends the CLI with JPEF specific flags.
func (r *Renderer) Flags() {
	flag.StringVar(&r.file, "file", "out.jpeg", "File to create with jpeg renderer")
	// TODO: Add flags for jpg options
}
