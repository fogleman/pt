package pt

import (
	"math/rand"
)

type Plane struct {
	center, normal Vector
	color          Color
	material       Material
}

func NewPlane(center, normal Vector, color Color, material Material) Shape {
	return &Plane{center, normal, color, material}
}

func (p *Plane) Box() Box {
	return Box{} // TODO: fix
}

func (p *Plane) Intersect(r Ray) float64 {
	k := r.Direction.Dot(p.normal)
	if k > EPS {
		return INF
	}
	t := p.center.Sub(r.Origin).Dot(p.normal) / k
	if t < EPS {
		return INF
	}
	return t
}

func (p *Plane) Color(v Vector) Color {
	return p.color
}

func (p *Plane) Material(v Vector) Material {
	return p.material
}

func (p *Plane) Normal(v Vector) Vector {
	return p.normal
}

func (p *Plane) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{} // TODO: fix
}
