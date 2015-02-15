package yaml

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/creack/goray/objects"
	"github.com/creack/goray/parser"
	"github.com/creack/goray/scene"
	"github.com/creack/goray/utils"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	parser.RegisterParser("yaml", &Parser{})
}

// Parser implement the GoRay parser interface for YAML scene files
type Parser struct {
}

// Extensions declares the supported file types of the parser.
// Used for automatic parser detection.
func (yp *Parser) Extensions() []string {
	return []string{
		"yaml",
		"yml",
	}
}

// toObjectConfig maps the local configuration object to the
// common Object's one.
func toObjectConfig(in objectConfig) objects.ObjectConfig {
	out := objects.ObjectConfig{
		Type: in.Type,
		Position: objects.Point{
			X: in.Position.X,
			Y: in.Position.Y,
			Z: in.Position.Z,
		},
		Rotation: objects.Vector{
			X: in.Rotation.X,
			Y: in.Rotation.Y,
			Z: in.Rotation.Z,
		},
		R: in.R,
	}
	out.Color = utils.RGBIntToColor(in.Color)
	return out
}

// Parse reads the config from the given file (or stdin) and generate
// the configuration object.
func (yp *Parser) Parse(filename string) (*scene.Config, error) {
	var conf config

	var inputStream io.Reader
	if filename == "-" {
		inputStream = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		inputStream = file
		defer file.Close()
	}
	content, err := ioutil.ReadAll(inputStream)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, &conf); err != nil {
		return nil, err
	}
	eye := scene.Eye{
		Position: objects.Point{
			X: conf.Eye.Position.X,
			Y: conf.Eye.Position.Y,
			Z: conf.Eye.Position.Z,
		},
		Rotation: objects.Vector{
			X: conf.Eye.Rotation.X,
			Y: conf.Eye.Rotation.Y,
			Z: conf.Eye.Rotation.Z,
		},
	}
	objs := []objects.Object{}
	for _, obj := range conf.Objects {
		newObjFct, ok := objects.ObjectList[obj.Type]
		if !ok {
			log.Printf("Unkown section: %s, skipping", obj.Type)
			continue
		}
		obj, err := newObjFct(toObjectConfig(obj))
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}

	if conf.Window.Width == 0 {
		conf.Window.Width = 800
	}
	if conf.Window.Height == 0 {
		conf.Window.Height = 600
	}

	return &scene.Config{
		Height:  conf.Window.Height,
		Width:   conf.Window.Width,
		Eye:     eye,
		Objects: objs,
	}, nil
}

type config struct {
	Window struct {
		Height int `yaml:"h"`
		Width  int `yaml:"w"`
	} `yaml:"window"`
	Eye     objectConfig   `yaml:"eye"`
	Objects []objectConfig `yaml:"objects"`
}

type objectConfig struct {
	Type     string `yaml:"type"`
	Position struct {
		X int `yaml:"x"`
		Y int `yaml:"y"`
		Z int `yaml:"z"`
	} `yaml:"position"`
	Rotation struct {
		X float64 `yaml:"x"`
		Y float64 `yaml:"y"`
		Z float64 `yaml:"z"`
	} `yaml:"rotation"`
	Color uint32 `yaml:"color"`
	R     int    `yaml:"R"`
}
