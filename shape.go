package pt

import (
	"math/rand"
)

type Shape interface {
	Intersect(Ray) float64
	Color() Color
	Material() Material
	Normal(Vector) Vector
	RandomPoint(*rand.Rand) Vector
}

type TransformedShape struct {
	Shape   Shape
	Matrix  Matrix
	Inverse Matrix
}

func NewTransformedShape(s Shape, m Matrix) *TransformedShape {
	return &TransformedShape{s, m, m.Inverse()}
}

func (s *TransformedShape) Intersect(r Ray) float64 {
	return s.Shape.Intersect(s.Inverse.MulRay(r))
}

func (s *TransformedShape) Color() Color {
	return s.Shape.Color()
}

func (s *TransformedShape) Material() Material {
	return s.Shape.Material()
}

func (s *TransformedShape) Normal(v Vector) Vector {
	return s.Matrix.MulVector(s.Shape.Normal(s.Inverse.MulVector(v)))
}

func (s *TransformedShape) RandomPoint(rnd *rand.Rand) Vector {
	return s.Matrix.MulVector(s.Shape.RandomPoint(rnd))
}
