package render

import (
	"github.com/creack/goray/objects"
	"github.com/creack/goray/rt"
)

var Renderers = map[string]Renderer{}

func RegisterRenderer(name string, renderer Renderer) {
	Renderers[name] = renderer
}

type Renderer interface {
	Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error
	Flags()
}
