package cli

import (
	"flag"
	"fmt"

	"github.com/creack/goray/parser"
	_ "github.com/creack/goray/parser/yaml" // default parser
	"github.com/creack/goray/render"
	_ "github.com/creack/goray/render/x11" // default renderer
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

type ParserCLI struct {
	name   string
	Parser parser.ParseFct
}

func (pc *ParserCLI) Set(value string) error {
	parse, ok := parser.Parsers[value]
	if !ok {
		possible := make([]string, len(parser.Parsers))
		i := 0
		for k := range parser.Parsers {
			possible[i] = k
			i++
		}
		return fmt.Errorf("Unkown parser: %s. Possible values: %v", value, possible)
	}
	pc.name = value
	pc.Parser = parse
	return nil
}

func (pc *ParserCLI) String() string {
	return pc.name
}

type CLIConfig struct {
	Renderer  RendererCLI
	Parser    ParserCLI
	SceneFile string
	Verbose   bool
}

func Flags() (*CLIConfig, error) {
	conf := &CLIConfig{}

	// Set Default
	conf.Renderer.Set("x11")
	conf.Parser.Set("yaml")

	// Get from command line
	flag.Var(&conf.Renderer, "renderer", "Renderer to use.")
	flag.Var(&conf.Parser, "parser", "Parser to use.")
	flag.StringVar(&conf.SceneFile, "scene", "", "Scene file to render")
	flag.BoolVar(&conf.Verbose, "v", false, "Verbose")
	flag.Parse()

	// Validate input
	if conf.SceneFile == "" {
		return nil, fmt.Errorf("Input scene file mandatory (-scene)")
	}
	return conf, nil
}
