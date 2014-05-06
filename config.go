package main

import (
	"fmt"
	"github.com/vaughan0/go-ini"
)

func parseConfig(filename string) (eye *Point, objects []Object, err error) {
	config, err := ini.LoadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	eye = &Point{}
	objects = []Object{}
	for section, values := range config {
		if section == "eye" {
			eye.Parse(values)
		} else {
			newObjFct, ok := objectList[section]
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
	return eye, objects, nil
}
