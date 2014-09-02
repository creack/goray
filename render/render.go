package render

import (
	"github.com/creack/goray/objects"
	"github.com/creack/goray/rt"
)

type NewRendererFct func(rtrace *rt.RT, eye *rt.Eye, objs []objects.Object) error

var Renderers = map[string]Renderer{}

func RegisterRenderer(name string, renderer Renderer) {
	Renderers[name] = renderer
}

type Renderer interface {
	Render(rt *rt.RT, eye *rt.Eye, objs []objects.Object) error
	Flags()
}
