package file

import (
	"flag"
	"image"
	"image/png"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/render"
	"github.com/creack/goray/rt"
)

func init() {
	render.RegisterRenderer("png", &PngRenderer{})
}

type PngRenderer struct {
	file string
}

func (pr *PngRenderer) pngRender(img image.Image) error {
	f, err := os.Create(pr.file)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}

func (pr *PngRenderer) Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error {
	return pr.pngRender(rt.Img)
}

func (pr *PngRenderer) Flags() {
	flag.StringVar(&pr.file, "file", "out.png", "File to create with png renderer")
}
