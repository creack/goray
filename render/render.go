package render

import (
	"github.com/creack/goray/objects"
	"github.com/creack/goray/scene"
)

// Rendereres hold all the available Renderer.
// This gets populated from the `init()` of the various
// implementations.
var Renderers = map[string]Renderer{}

// RegisterRenderer registers the given Renderer with the given name.
func RegisterRenderer(name string, renderer Renderer) {
	Renderers[name] = renderer
}

// Renderer is the interface that wraps the Renderer.
type Renderer interface {
	// Render renders the scene (`rt`) with the object list
	// From the `eye` perspective.
	Render(s *scene.Scene, eye scene.Eye, objs []objects.Object) error
	// Flags extends the CLI with implementation specific flags.
	Flags()
}
