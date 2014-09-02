package main

import (
	"fmt"
	"io/ioutil"
	"log"

	goyaml "gopkg.in/yaml.v1"

	"github.com/creack/goray/cli"
	"github.com/creack/goray/objects"
	_ "github.com/creack/goray/objects/cone"
	_ "github.com/creack/goray/objects/plan"
	_ "github.com/creack/goray/objects/sphere"
	_ "github.com/creack/goray/render/file"
	_ "github.com/creack/goray/render/x11"
	"github.com/creack/goray/rt"
)

type Config struct {
	Verbose bool                   `yaml:"verbose"`
	Eye     objects.ObjectConfig   `yaml:"eye"`
	Objects []objects.ObjectConfig `yaml:"objects"`
}

func parseConfigYaml(filename string) (*rt.Eye, []objects.Object, error) {
	content, err := ioutil.ReadFile("rt.yaml")
	if err != nil {
		return nil, nil, err
	}
	var conf Config
	if err := goyaml.Unmarshal(content, &conf); err != nil {
		return nil, nil, err
	}

	eye := &rt.Eye{
		Position: conf.Eye.Position,
		Rotation: conf.Eye.Rotation,
	}

	objs := []objects.Object{}
	for _, obj := range conf.Objects {
		newObjFct, ok := objects.ObjectList[obj.Type]
		if !ok {
			log.Printf("Unkown section: %s, skipping", obj.Type)
			continue
		}
		obj, err := newObjFct(obj)
		if err != nil {
			return nil, nil, err
		}
		objs = append(objs, obj)
	}

	if conf.Verbose {
		fmt.Printf("-> %T - %#v\n", eye, eye)
		for _, elem := range objs {
			fmt.Printf("-> %T - %#v\n", elem, elem)
		}
	}
	return eye, objs, nil
}

func main() {
	cliConf, err := cli.Flags()
	if err != nil {
		log.Fatal(err)
	}
	eye, objs, err := parseConfigYaml(cliConf.SceneFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	rtrace := rt.NewRT(800, 600)
	rtrace.Verbose = true
	rtrace.FillImage(eye.Position, objs)

	if err := cliConf.Renderer.Renderer.Render(rtrace, eye, objs); err != nil {
		log.Fatal(err)
	}
}
