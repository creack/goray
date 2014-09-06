package cli

import (
	"fmt"

	"github.com/creack/goray/render"
)

type RendererCLI struct {
	name     string
	Renderer render.Renderer
}

func (rc *RendererCLI) Set(value string) error {
	renderer, ok := render.Renderers[value]
	if !ok {
		possible := make([]string, len(render.Renderers))
		i := 0
		for k := range render.Renderers {
			possible[i] = k
			i++
		}
		return fmt.Errorf("Unkown renderer: %s. Possible values: %v", value, possible)
	}
	renderer.Flags()
	rc.name = value
	rc.Renderer = renderer
	return nil
}

func (rc RendererCLI) String() string {
	return rc.name
}
