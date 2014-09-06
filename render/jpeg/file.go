package jpeg

import (
	"flag"
	"image"
	"image/jpeg"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/render"
	"github.com/creack/goray/rt"
)

func init() {
	render.RegisterRenderer("jpeg", &JpegRenderer{})
}

type JpegRenderer struct {
	file string
}

func (pr *JpegRenderer) jpegRender(img image.Image) error {
	f, err := os.Create(pr.file)
	if err != nil {
		return err
	}
	if err := jpeg.Encode(f, img, nil); err != nil {
		return err
	}
	return nil
}

func (pr *JpegRenderer) Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error {
	return pr.jpegRender(rt.Img)
}

func (pr *JpegRenderer) Flags() {
	flag.StringVar(&pr.file, "file", "out.jpeg", "File to create with jpeg renderer")
	// TODO: Add flags for jpg options
}
