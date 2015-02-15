package main

import (
	"log"

	"github.com/creack/goray/cli"
	_ "github.com/creack/goray/parser/all"
	_ "github.com/creack/goray/render/all"
	"github.com/creack/goray/scene"
)

func main() {
	// Process CLI flags
	cliConf, err := cli.Flags()
	if err != nil {
		log.Fatalf("flags: %s", err)
	}

	// Parse the scene file
	sceneConf, err := cliConf.Parser.Parser.Parse(cliConf.SceneFile)
	if err != nil {
		log.Fatalf("parse: %s", err)
	}

	// Process the image
	rtrace := scene.NewScene(sceneConf.Width, sceneConf.Height)
	rtrace.Verbose = cliConf.Verbose
	rtrace.Compute(sceneConf.Eye.Position, sceneConf.Objects)

	// Render the image
	if err := cliConf.Renderer.Renderer.Render(rtrace, sceneConf.Eye, sceneConf.Objects); err != nil {
		log.Fatalf("render: %s", err)
	}
}
