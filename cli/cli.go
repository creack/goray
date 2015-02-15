package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/parser"
	_ "github.com/creack/goray/parser/yaml" // default parser
	"github.com/creack/goray/render"
	_ "github.com/creack/goray/render/x11" // default renderer
)

// Set defaults.
const (
	DefaultParser   = "yaml"
	DefaultRenderer = "x11"
)

// Config represent the RT configuration variables
type Config struct {
	Renderer  RendererCLI
	Parser    ParserCLI
	SceneFile string
	Verbose   bool
}

// Flags handle CLIs flags
func Flags() (*Config, error) {
	conf := &Config{}

	// Set Defaults/
	conf.Renderer.Set(DefaultRenderer)
	conf.Parser.Set(DefaultParser)

	// Use different name to differenciate set/unset.
	conf.Parser.name += "."

	// Get from command line
	flag.Var(&conf.Renderer, "renderer", "Renderer to use.")
	flag.Var(&conf.Parser, "parser", "Parser to use.")
	flag.StringVar(&conf.SceneFile, "scene", "", "Scene file to render")
	flag.BoolVar(&conf.Verbose, "v", false, "Verbose")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s -scene=scene_file\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nAvailable Parsers:\n")
		for p := range parser.Parsers {
			fmt.Fprintf(os.Stderr, "- %s\n", p)
		}
		fmt.Fprintf(os.Stderr, "\nAvailable Renderers:\n")
		for r := range render.Renderers {
			fmt.Fprintf(os.Stderr, "- %s\n", r)
		}
		fmt.Fprintf(os.Stderr, "\nAvailable Objects:\n")
		for o := range objects.ObjectList {
			fmt.Fprintf(os.Stderr, "- %s\n", o)
		}
	}

	flag.Parse()

	// Validate input
	if conf.SceneFile == "" {
		return nil, fmt.Errorf("Input scene file missing (-scene)")
	}

	// Autodetect parser if not set.
	if conf.Parser.name == DefaultParser+"." {
		name := DetectParser(conf.SceneFile)
		if name == "" {
			return nil, fmt.Errorf("Unkown scene format: %s", conf.SceneFile)
		}
		if name != DefaultParser {
			conf.Parser.Set(name)
		} else {
			conf.Parser.name = DefaultParser
		}

	}

	if conf.Verbose {
		fmt.Fprintf(os.Stderr, "Parser: %s\nRenderer: %s\nSceneFile: %s\n", conf.Parser.name, conf.Renderer.name, conf.SceneFile)
	}

	return conf, nil
}
