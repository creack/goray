package main

import (
	"fmt"
	"github.com/vaughan0/go-ini"
	"strings"
)

func parseConfig(filename string) (eye *Point, objects []Object, err error) {
	config, err := ini.LoadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	objects = []Object{}
	for section, values := range config {
		if section == "eye" || section == "camera" {
			if eye, err = (*Point).Parse(nil, values); err != nil {
				return nil, nil, err
			}
		} else {
			newObjFct, ok := objectList[strings.Split(section, ".")[0]]
			if !ok {
				return nil, nil, fmt.Errorf("Unkown section: %s", section)
			}
			obj, err := newObjFct(values)
			if err != nil {
				return nil, nil, err
			}
			objects = append(objects, obj)
		}
	}
	if eye == nil {
		return nil, nil, fmt.Errorf("no definition for the camera")
	}
	if len(objects) == 0 {
		return nil, nil, fmt.Errorf("no object definition")
	}
	return eye, objects, nil
}
