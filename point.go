package main

type Point struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
	Z int `yaml:"z"`
}

func (p *Point) Add(p2 Point) {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
}

func (p *Point) Sub(p2 Point) {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
}

func (p *Point) Parse(obj ObjectConfig) (Point, error) {
	if p == nil {
		p = &Point{}
	}
	p.X, p.Y, p.Z = obj.Position.X, obj.Position.Y, obj.Position.Z
	return *p, nil
}
