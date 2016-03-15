package pt

import "math/rand"

type Shape interface {
	Compile()
	Box() Box
	Intersect(Ray) Hit
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

func (s *TransformedShape) Box() Box {
	return s.matrix.MulBox(s.Shape.Box())
}

func (s *TransformedShape) Intersect(r Ray) Hit {
	shapeRay := s.inverse.MulRay(r)
	hit := s.Shape.Intersect(shapeRay)
	if !hit.Ok() {
		return hit
	}
	shape := hit.Shape
	shapePosition := shapeRay.Position(hit.T)
	shapeNormal := shape.Normal(shapePosition)
	position := s.matrix.MulPosition(shapePosition)
	normal := s.inverse.Transpose().MulDirection(shapeNormal)
	color := shape.Color(shapePosition)
	material := shape.Material(shapePosition)
	inside := false
	if shapeNormal.Dot(shapeRay.Direction) > 0 {
		normal = normal.MulScalar(-1)
		inside = true
	}
	ray := Ray{position, normal}
	info := HitInfo{shape, position, normal, ray, color, material, inside}
	hit.T = position.Sub(r.Origin).Length()
	hit.HitInfo = &info
	return hit
}
