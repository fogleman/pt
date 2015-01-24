package pt

import (
	"math"
)

type Scene struct {
	Shapes []Shape
	Lights []Shape
}

func (s *Scene) AddShape(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) AddLight(light Shape) {
	s.Lights = append(s.Lights, light)
}

func (s *Scene) Intersect(r Ray) (Hit, bool) {
	hit := Hit{}
	u := INF
	for _, shape := range s.Shapes {
		t := shape.Intersect(r)
		if t < u {
			u = t
			p := r.Origin.Add(r.Direction.Mul(t))
			n := shape.Normal(p)
			hit = Hit{shape, Ray{p, n}}
		}
	}
	ok := u < INF
	return hit, ok
}

func (s *Scene) Shadow(r Ray) bool {
	// TODO: ignore objects behind the light source
	for _, shape := range s.Shapes {
		t := shape.Intersect(r)
		if t < INF {
			return true
		}
	}
	return false
}

func (s *Scene) Light(r Ray) Color {
	color := Color{}
	for _, light := range s.Lights {
		lr := Ray{r.Origin, light.RandomPoint().Sub(r.Origin).Normalize()}
		if (s.Shadow(lr)) {
			continue
		}
		n := math.Max(0, lr.Direction.Dot(r.Direction))
		color = color.Add(light.Color().Mul(n))
	}
	return color
}

func (s *Scene) Sample(r Ray) Color {
	if hit, ok := s.Intersect(r); ok {
		return hit.Shape.Color().MulColor(s.Light(hit.Ray))
	}
	return Color{}
}
