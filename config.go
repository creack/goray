package main

import (
	"fmt"
	"io/ioutil"

	goyaml "gopkg.in/yaml.v1"
)

type ObjectConfig struct {
	Type     string `yaml:"type"`
	Position Point  `yaml:"position"`
	Color    int    `yaml:"color"`
	R        int    `yaml:"R"`
}

type Config struct {
	Eye     ObjectConfig   `yaml:"eye"`
	Objects []ObjectConfig `yaml:"objects"`
}

func parseConfigYaml(filename string) (eye *Point, objects []Object, err error) {
	content, err := ioutil.ReadFile("rt.yaml")
	if err != nil {
		return nil, nil, err
	}
	var conf Config
	if err := goyaml.Unmarshal(content, &conf); err != nil {
		return nil, nil, err
	}

	eye1, err := (*Point).Parse(nil, conf.Eye)
	if err != nil {
		return nil, nil, err
	}
	eye = &eye1
	fmt.Printf("-> %T - %#v\n", eye, eye)
	objects = []Object{}
	for _, obj := range conf.Objects {
		newObjFct, ok := objectList[obj.Type]
		if !ok {
			return nil, nil, fmt.Errorf("Unkown section: %s", obj.Type)
		}
		obj, err := newObjFct(obj)
		if err != nil {
			return nil, nil, err
		}
		objects = append(objects, obj)
	}
	for _, elem := range objects {
		fmt.Printf("-> %T - %#v\n", elem, elem)
	}
	return eye, objects, nil
}

func parseConfig(filename string) (eye *Point, objects []Object, err error) {
	// config, err := ini.LoadFile(filename)
	// if err != nil {
	// 	return nil, nil, err
	// }
	// objects = []Object{}
	// for section, values := range config {
	// 	if section == "eye" || section == "camera" {
	// 		if eye, err = (*Point).Parse(nil, values); err != nil {
	// 			return nil, nil, err
	// 		}
	// 	} else {
	// 		newObjFct, ok := objectList[strings.Split(section, ".")[0]]
	// 		if !ok {
	// 			return nil, nil, fmt.Errorf("Unkown section: %s", section)
	// 		}
	// 		obj, err := newObjFct(values)
	// 		if err != nil {
	// 			return nil, nil, err
	// 		}
	// 		objects = append(objects, obj)
	// 	}
	// }
	// if eye == nil {
	// 	return nil, nil, fmt.Errorf("no definition for the camera")
	// }
	// if len(objects) == 0 {
	// 	return nil, nil, fmt.Errorf("no object definition")
	// }
	// return eye, objects, nil
	return nil, nil, nil
}
