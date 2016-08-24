package pt

import "math"

type SDFShape struct {
}

type SDF interface {
	Evaluate(p Vector) float64
}

type SphereSDF struct {
	Radius float64
}

func (s *SphereSDF) Evaluate(p Vector) float64 {
	return p.Length() - s.Radius
}

type CubeSDF struct {
	Size Vector
}

func (s *CubeSDF) Evaluate(p Vector) float64 {
	d := p.Abs().Sub(s.Size)
	return math.Min(math.Max(d.X, math.Max(d.Y, d.Z)), 0) + d.Max(Vector{}).Length()
}

type UnionSDF struct {
	A, B SDF
}

func (s *UnionSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Min(a, b)
}

type DifferenceSDF struct {
	A, B SDF
}

func (s *DifferenceSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Max(-a, b)
}

type IntersectionSDF struct {
	A, B SDF
}

func (s *IntersectionSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Max(a, b)
}
