package main

type Point struct {
	x, y, z int
}

func (p *Point) Add(p2 Point) {
	p.x += p2.x
	p.y += p2.y
	p.z += p2.z
}

func (p *Point) Sub(p2 Point) {
	p.x -= p2.x
	p.y -= p2.y
	p.z -= p2.z
}

func (p *Point) Parse(obj ObjectConfig) (Point, error) {
	if p == nil {
		p = &Point{}
	}
	p.x, p.y, p.z = obj.X, obj.Y, obj.Z
	return *p, nil
}
