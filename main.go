package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	goyaml "gopkg.in/yaml.v1"

	"github.com/creack/goray/objects"
	_ "github.com/creack/goray/objects/cone"
	_ "github.com/creack/goray/objects/plan"
	_ "github.com/creack/goray/objects/sphere"
	"github.com/creack/goray/utils"
)

type Config struct {
	Verbose bool                   `yaml:"verbose"`
	Eye     objects.ObjectConfig   `yaml:"eye"`
	Objects []objects.ObjectConfig `yaml:"objects"`
}

type Eye struct {
	Position objects.Point
	Rotation objects.Vector
}

func parseConfigYaml(filename string) (*Eye, []objects.Object, error) {
	content, err := ioutil.ReadFile("rt.yaml")
	if err != nil {
		return nil, nil, err
	}
	var conf Config
	if err := goyaml.Unmarshal(content, &conf); err != nil {
		return nil, nil, err
	}

	eye := &Eye{
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

type RT struct {
	img     *image.RGBA
	width   int
	height  int
	verbose bool
}

func NewRT(x, y int) *RT {
	return &RT{
		img:    image.NewRGBA(image.Rect(0, 0, x, y)),
		width:  x,
		height: y,
	}
}

func (rt *RT) calc(x, y int, eye objects.Point, objs []objects.Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = objects.Vector{
			X: 100,
			Y: float64(rt.width/2 - x),
			Z: float64(rt.height/2 - y),
		}
	)
	for _, obj := range objs {
		if tmp := obj.Intersect(v, eye); tmp > 0 && (k == -1 || tmp < k) {
			k = tmp
			col = obj.Color()
		}
	}
	return col
}

func (rt *RT) fillImage(eye objects.Point, objs []objects.Object) {
	var (
		x int
		y int
	)

	for i, total := 0, rt.width*rt.height; i < total; i++ {
		x = i % rt.width
		y = i / rt.width
		if rt.verbose && x == 0 && y%10 == 0 {
			fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(rt.height))*100+1))
		}
		rt.img.Set(x, y, rt.calc(x, y, eye, objs))
	}
	if rt.verbose {
		fmt.Println("\rProcessing: 100%")
	}
}

func main() {
	eye, objs, err := parseConfigYaml("rt.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	rt := NewRT(800, 600)
	rt.verbose = true
	rt.fillImage(eye.Position, objs)
	w, err := x11.NewWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	fct := func() {
		rt.fillImage(eye.Position, objs)
		draw.Draw(w.Screen(), w.Screen().Bounds(), rt.img, image.ZP, draw.Src)
		w.FlushImage()
	}
	fct()
	for e := range w.EventChan() {
		switch e := e.(type) {
		case ui.KeyEvent:
			switch utils.KeyListInt[e.Key] {
			case " ", "<esc>", "\n", "q":
				return
			case "<up>":
				eye.Position.X += 10
			case "<down>":
				eye.Position.X -= 10
			case "<left>":
				eye.Position.Y += 10
			case "<right>":
				eye.Position.Y -= 10
			case "a":
				eye.Position.Z += 10
			case "z":
				eye.Position.Z -= 10
			}
			fct()
		}
	}
}
