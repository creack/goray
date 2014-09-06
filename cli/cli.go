package cli

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/creack/goray/parser/yaml" // default parser
	_ "github.com/creack/goray/render/x11"  // default renderer
)

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

	// Use different name to differenciate set/unset.
	conf.Parser.name = "yaml."

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

	// Autodetect parser if not set.
	if conf.Parser.name == "yaml." {
		name := DetectParser(conf.SceneFile)
		if name == "" {
			return nil, fmt.Errorf("Unkown scene format: %s", conf.SceneFile)
		}
		if name != "yaml" {
			conf.Parser.Set(name)
		} else {
			conf.Parser.name = "yaml"
		}

	}

	if conf.Verbose {
		fmt.Fprintf(os.Stderr, "Parser: %s\nRenderer: %s\nSceneFile: %s\n", conf.Parser.name, conf.Renderer.name, conf.SceneFile)
	}

	return conf, nil
}
