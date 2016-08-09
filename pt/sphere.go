package pt

import "math"

type Sphere struct {
	center   Vector
	radius   float64
	material Material
	box      Box
}

func NewSphere(center Vector, radius float64, material Material) Shape {
	min := Vector{center.X - radius, center.Y - radius, center.Z - radius}
	max := Vector{center.X + radius, center.Y + radius, center.Z + radius}
	box := Box{min, max}
	return &Sphere{center, radius, material, box}
}

func (s *Sphere) Compile() {
}

func (s *Sphere) BoundingBox() Box {
	return s.box
}

func (s *Sphere) Intersect(r Ray) Hit {
	to := r.Origin.Sub(s.center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - s.radius*s.radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > 0 {
			return Hit{s, t1, nil}
		}
		// t2 := -b + d
		// if t2 > 0 {
		// 	return Hit{s, t2}
		// }
	}
	return NoHit
}

func (s *Sphere) ColorAt(p Vector) Color {
	if s.material.Texture == nil {
		return s.material.Color
	}
	p = p.Sub(s.center)
	u := math.Atan2(p.Z, p.X)
	v := math.Atan2(p.Y, Vector{p.X, 0, p.Z}.Length())
	u = (u + math.Pi) / (2 * math.Pi)
	v = (v + math.Pi/2) / math.Pi
	return s.material.Texture.Sample(u, v)
}

func (s *Sphere) MaterialAt(p Vector) Material {
	return s.material
}

func (s *Sphere) NormalAt(p Vector) Vector {
	return p.Sub(s.center).Normalize()
}
