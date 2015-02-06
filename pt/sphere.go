package pt

import (
	"image"
	"math"
	"math/rand"
)

type Sphere struct {
	center   Vector
	radius   float64
	material Material
	texture  image.Image
	box      Box
}

func NewSphere(center Vector, radius float64, material Material, texture image.Image) Shape {
	min := Vector{center.X - radius, center.Y - radius, center.Z - radius}
	max := Vector{center.X + radius, center.Y + radius, center.Z + radius}
	box := Box{min, max}
	return &Sphere{center, radius, material, texture, box}
}

func (s *Sphere) Compile() {
}

func (s *Sphere) Box() Box {
	return s.box
}

func (s *Sphere) Intersect(r Ray) Hit {
	to := r.Origin.Sub(s.center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - s.radius*s.radius
	d := b*b - c
	if d > 0 {
		t := -b - math.Sqrt(d)
		if t > 0 {
			return Hit{s, t}
		}
	}
	return NoHit
}

func (s *Sphere) Color(p Vector) Color {
	if s.texture == nil {
		return s.material.Color
	}
	// TODO: make a Texture interface
	size := s.texture.Bounds().Max
	u := math.Atan2(p.Z, p.X)
	v := math.Atan2(p.Y, Vector{p.X, 0, p.Z}.Length())
	u = (u + math.Pi) / (2 * math.Pi)
	v = 1 - (v+math.Pi/2)/math.Pi
	x := int(u * float64(size.X))
	y := int(v * float64(size.Y))
	r, g, b, _ := s.texture.At(x, y).RGBA()
	return Color{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535}
}

func (s *Sphere) Material(p Vector) Material {
	return s.material
}

func (s *Sphere) Normal(p Vector) Vector {
	return p.Sub(s.center).Normalize()
}

func (s *Sphere) RandomPoint(rnd *rand.Rand) Vector {
	for {
		x := rnd.Float64()*2 - 1
		y := rnd.Float64()*2 - 1
		z := rnd.Float64()*2 - 1
		v := Vector{x, y, z}
		if v.Length() <= 1 {
			return v.Mul(s.radius).Add(s.center)
		}
	}
}
