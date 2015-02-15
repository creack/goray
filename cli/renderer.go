package cli

import (
	"fmt"

	"github.com/creack/goray/render"
)

// RendererCLI wraps a renderer with a name.
// Used to select the renderer from the CLI.
type RendererCLI struct {
	name     string
	Renderer render.Renderer
}

// Set will set and initialize the RendererCLI.
func (rc *RendererCLI) Set(value string) error {
	renderer, ok := render.Renderers[value]
	if !ok {
		possible := make([]string, 0, len(render.Renderers))
		for k := range render.Renderers {
			possible = append(possible, k)

		}
		return fmt.Errorf("Unkown renderer: %q. Possible values: %v", value, possible)
	}
	renderer.Flags()
	rc.name = value
	rc.Renderer = renderer
	return nil
}

// String returns the renderer's name.
func (rc RendererCLI) String() string {
	return rc.name
}
