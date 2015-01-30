package pt

import (
	"math/rand"
)

type Shape interface {
	Intersect(Ray) float64
	Color(Vector) Color
	Material(Vector) Material
	Normal(Vector) Vector
	RandomPoint(*rand.Rand) Vector
}

type TransformedShape struct {
	Shape
	matrix  Matrix
	inverse Matrix
}

func NewTransformedShape(s Shape, m Matrix) Shape {
	return &TransformedShape{s, m, m.Inverse()}
}

func (s *TransformedShape) Intersect(r Ray) float64 {
	return s.Shape.Intersect(s.inverse.MulRay(r))
}

func (s *TransformedShape) Color(p Vector) Color {
	return s.Shape.Color(s.inverse.MulVector(p))
}

func (s *TransformedShape) Material(p Vector) Material {
	return s.Shape.Material(s.inverse.MulVector(p))
}

func (s *TransformedShape) Normal(p Vector) Vector {
	return s.matrix.MulDirection(s.Shape.Normal(s.inverse.MulVector(p)))
}

func (s *TransformedShape) RandomPoint(rnd *rand.Rand) Vector {
	return s.matrix.MulVector(s.Shape.RandomPoint(rnd))
}
