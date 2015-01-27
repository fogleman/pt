package pt

import (
	"image"
	"math"
	"math/rand"
)

type Sphere struct {
	Center  Vector
	Radius  float64
	Col     Color
	Mat     Material
	Texture image.Image
}

func (s *Sphere) Intersect(r Ray) float64 {
	to := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * to.Dot(r.Direction)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - 4*a*c
	if d > 0 {
		t := (-b - math.Sqrt(d)) / (2 * a)
		if t > 0 {
			return t
		}
	}
	return INF
}

func (s *Sphere) Color(p Vector) Color {
	if s.Texture == nil {
		return s.Col
	}
	size := s.Texture.Bounds().Max
	u := math.Atan2(p.Z, p.X)
	v := math.Atan2(p.Y, Vector{p.X, 0, p.Z}.Length())
	u = (u + math.Pi) / (2 * math.Pi)
	v = 1 - (v+math.Pi/2)/math.Pi
	x := int(u * float64(size.X))
	y := int(v * float64(size.Y))
	r, g, b, _ := s.Texture.At(x, y).RGBA()
	return Color{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535}
}

func (s *Sphere) Material() Material {
	return s.Mat
}

func (s *Sphere) Normal(p Vector) Vector {
	return p.Sub(s.Center).Normalize()
}

func (s *Sphere) RandomPoint(rnd *rand.Rand) Vector {
	for {
		x := rnd.Float64()*2 - 1
		y := rnd.Float64()*2 - 1
		z := rnd.Float64()*2 - 1
		v := Vector{x, y, z}
		if v.Length() <= 1 {
			return v.Mul(s.Radius).Add(s.Center)
		}
	}
}
