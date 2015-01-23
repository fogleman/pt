package pt

import (
	"math"
)

type Scene struct {
	shapes []Shape
}

func (s *Scene) Add(shape Shape) {
	s.shapes = append(s.shapes, shape)
}

func (s *Scene) Intersect(r Ray) float64 {
	t := INF
	for _, shape := range s.shapes {
		t = math.Min(t, shape.Intersect(r))
	}
	return t
}
