package cli

import (
	"flag"
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
	println("--> ok")
	return nil
}

func (rc RendererCLI) String() string {
	return rc.name
}

type CLIConfig struct {
	Renderer  RendererCLI
	SceneFile string
}

func Flags() (*CLIConfig, error) {
	conf := &CLIConfig{
		Renderer: RendererCLI{name: "x11", Renderer: render.Renderers["x11"]},
	}
	flag.Var(&conf.Renderer, "renderer", "Renderer to use.")
	flag.StringVar(&conf.SceneFile, "scene", "", "Schene file to render")
	flag.Parse()

	// Validate input
	if conf.SceneFile == "" {
		return nil, fmt.Errorf("Input scene file mandatory (-scene)")
	}
	return conf, nil
}
