package main

import (
	"strconv"
)

type Point struct {
	x, y, z int
}

func (p *Point) Add(p2 *Point) {
	p.x += p2.x
	p.y += p2.y
	p.z += p2.z
}

func (p *Point) Sub(p2 *Point) {
	p.x -= p2.x
	p.y -= p2.y
	p.z -= p2.z
}

func (p *Point) Parse(values map[string]string) (*Point, error) {
	if p == nil {
		p = &Point{}
	}
	xStr, yStr, zStr := values["x"], values["y"], values["z"]
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return nil, err
	}
	z, err := strconv.Atoi(zStr)
	if err != nil {
		return nil, err
	}
	p.x, p.y, p.z = x, y, z
	return p, nil
}
