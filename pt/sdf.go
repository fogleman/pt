package pt

import "math"

type SDFShape struct {
	SDF
	Material Material
}

func NewSDFShape(sdf SDF, material Material) Shape {
	return &SDFShape{sdf, material}
}

func (s *SDFShape) Compile() {
}

func (s *SDFShape) BoundingBox() Box {
	return s.SDF.BoundingBox()
}

func (s *SDFShape) Intersect(ray Ray) Hit {
	box := s.BoundingBox()
	t1, t2 := box.Intersect(ray)
	if t2 < t1 || t2 < 0 {
		return NoHit
	}
	epsilon := 0.0001
	t := math.Max(0.001, t1)
	for i := 0; i < 100; i++ {
		p := ray.Position(t)
		d := s.Evaluate(p)
		if d < epsilon {
			return Hit{s, t, nil}
		}
		t += d
		if t > t2 {
			return NoHit
		}
	}
	return NoHit
}

func (s *SDFShape) UV(p Vector) Vector {
	return Vector{}
}

func (s *SDFShape) NormalAt(p Vector) Vector {
	eps := 0.0001
	n := Vector{
		s.Evaluate(Vector{p.X - eps, p.Y, p.Z}) - s.Evaluate(Vector{p.X + eps, p.Y, p.Z}),
		s.Evaluate(Vector{p.X, p.Y - eps, p.Z}) - s.Evaluate(Vector{p.X, p.Y + eps, p.Z}),
		s.Evaluate(Vector{p.X, p.Y, p.Z - eps}) - s.Evaluate(Vector{p.X, p.Y, p.Z + eps}),
	}
	return n.Normalize()
}

func (s *SDFShape) MaterialAt(p Vector) Material {
	return s.Material
}

type SDF interface {
	Evaluate(p Vector) float64
	BoundingBox() Box
}

type SphereSDF struct {
	Radius float64
}

func (s *SphereSDF) Evaluate(p Vector) float64 {
	return p.Length() - s.Radius
}

func (s *SphereSDF) BoundingBox() Box {
	r := s.Radius
	return Box{Vector{-r, -r, -r}, Vector{r, r, r}}
}

type CubeSDF struct {
	Size Vector
}

func (s *CubeSDF) Evaluate(p Vector) float64 {
	d := p.Abs().Sub(s.Size)
	return math.Min(math.Max(d.X, math.Max(d.Y, d.Z)), 0) + d.Max(Vector{}).Length()
}

func (s *CubeSDF) BoundingBox() Box {
	x, y, z := s.Size.X, s.Size.Y, s.Size.Z
	return Box{Vector{-x, -y, -z}, Vector{x, y, z}}
}

type CylinderSDF struct {
	Height float64
}

func (s *CylinderSDF) Evaluate(p Vector) float64 {
	d := Vector{Vector{p.X, p.Z, 0}.Length(), p.Y, 0}.Abs().SubScalar(s.Height)
	return math.Min(math.Max(d.X, d.Y), 0) + d.Max(Vector{}).Length()
}

func (s *CylinderSDF) BoundingBox() Box {
	h := s.Height
	return Box{Vector{-1, -h, -1}, Vector{1, h, 1}}
}

type CapsuleSDF struct {
	A, B   Vector
	Radius float64
}

func (s *CapsuleSDF) Evaluate(p Vector) float64 {
	pa := p.Sub(s.A)
	ba := s.B.Sub(s.A)
	h := math.Max(0, math.Min(1, pa.Dot(ba)/ba.Dot(ba)))
	return pa.Sub(ba.MulScalar(h)).Length() - s.Radius
}

func (s *CapsuleSDF) BoundingBox() Box {
	a, b := s.A.Min(s.B), s.A.Max(s.B)
	return Box{a.SubScalar(s.Radius), b.AddScalar(s.Radius)}
}

type TransformSDF struct {
	SDF
	Matrix Matrix
}

func (s *TransformSDF) Evaluate(p Vector) float64 {
	// TODO: precompute inverse
	q := s.Matrix.Inverse().MulPosition(p)
	return s.SDF.Evaluate(q)
}

func (s *TransformSDF) BoundingBox() Box {
	return s.Matrix.MulBox(s.SDF.BoundingBox())
}

type UnionSDF struct {
	A, B SDF
}

func (s *UnionSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Min(a, b)
}

func (s *UnionSDF) BoundingBox() Box {
	a := s.A.BoundingBox()
	b := s.B.BoundingBox()
	return a.Extend(b)
}

type DifferenceSDF struct {
	A, B SDF
}

func (s *DifferenceSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Max(a, -b)
}

func (s *DifferenceSDF) BoundingBox() Box {
	return s.A.BoundingBox()
}

type IntersectionSDF struct {
	A, B SDF
}

func (s *IntersectionSDF) Evaluate(p Vector) float64 {
	a := s.A.Evaluate(p)
	b := s.B.Evaluate(p)
	return math.Max(a, b)
}

func (s *IntersectionSDF) BoundingBox() Box {
	// TODO: intersect boxes
	a := s.A.BoundingBox()
	b := s.B.BoundingBox()
	return a.Extend(b)
}
