package pt

import (
	"math"
	"math/rand"
)

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
	// http://dev.theomader.com/transform-bounding-boxes/
	b := s.Shape.Box()
	minx, maxx := b.Min.X, b.Max.X
	miny, maxy := b.Min.Y, b.Max.Y
	minz, maxz := b.Min.Z, b.Max.Z
	m := &s.matrix
	xa := m.x00*minx + m.x10*minx + m.x20*minx + m.x30*minx
	xb := m.x00*maxx + m.x10*maxx + m.x20*maxx + m.x30*maxx
	ya := m.x01*miny + m.x11*miny + m.x21*miny + m.x31*miny
	yb := m.x01*maxy + m.x11*maxy + m.x21*maxy + m.x31*maxy
	za := m.x02*minz + m.x12*minz + m.x22*minz + m.x32*minz
	zb := m.x02*maxz + m.x12*maxz + m.x22*maxz + m.x32*maxz
	minx, maxx = math.Min(xa, xb), math.Max(xa, xb)
	miny, maxy = math.Min(ya, yb), math.Max(ya, yb)
	minz, maxz = math.Min(za, zb), math.Max(za, zb)
	min := Vector{minx + m.x03, miny + m.x13, minz + m.x23}
	max := Vector{maxx + m.x03, maxy + m.x13, maxz + m.x23}
	return Box{min, max}
}

func (s *TransformedShape) Intersect(r Ray) Hit {
	hit := s.Shape.Intersect(s.inverse.MulRay(r))
	// if s.Shape is a Mesh, the hit.Shape will be a Triangle in the Mesh
	// we need to transform this Triangle, not the Mesh itself
	shape := &TransformedShape{hit.Shape, s.matrix, s.inverse}
	return Hit{shape, hit.T}
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
