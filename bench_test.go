package main

import (
	"testing"
)

func BenchmarkFillImage(b *testing.B) {
	eye := &Point{
		x: -300,
		y: 0,
		z: 100,
	}
	objs := []Object{
		&Plan{
			color: RgbIntToColor(0xFF0000),
			z:     0,
		},
		&Sphere{
			position: &Point{
				x: 100,
				y: 0,
				z: 0,
			},
			R:     100,
			color: RgbIntToColor(0xFFFF00),
		},
	}
	rt := NewRT(800, 600)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rt.fillImage(eye, objs)
	}
}

func BenchmarkPlan(b *testing.B) {
	eye := &Point{
		x: -300,
		y: 0,
		z: 100,
	}
	p := &Plan{
		color: RgbIntToColor(0xFF0000),
		z:     0,
	}
	v := &Vector{}
	for i := 0; i < b.N; i++ {
		p.Intersect(v, eye)
	}
}

func BenchmarkSphere(b *testing.B) {
	eye := &Point{
		x: -300,
		y: 0,
		z: 100,
	}
	s := &Sphere{
		position: &Point{
			x: 100,
			y: 0,
			z: 0,
		},
		R:     100,
		color: RgbIntToColor(0xFFFF00),
	}
	v := &Vector{}
	for i := 0; i < b.N; i++ {
		s.Intersect(v, eye)
	}
}

func BenchmarkSecondDegree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecondDegree(1, 2, 3)
	}
}
