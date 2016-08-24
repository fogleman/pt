package pt

import "math"

type Plane struct {
	Point    Vector
	Normal   Vector
	Material Material
}

func NewPlane(point, normal Vector, material Material) *Plane {
	normal = normal.Normalize()
	return &Plane{point, normal, material}
}

func (p *Plane) Compile() {
}

func (p *Plane) BoundingBox() Box {
	return Box{Vector{-INF, -INF, -INF}, Vector{INF, INF, INF}}
}

func (p *Plane) Intersect(ray Ray) Hit {
	d := p.Normal.Dot(ray.Direction)
	if math.Abs(d) < EPS {
		return NoHit
	}
	a := p.Point.Sub(ray.Origin)
	t := a.Dot(p.Normal) / d
	if t < EPS {
		return NoHit
	}
	return Hit{p, t, nil}
}

func (p *Plane) UV(a Vector) Vector {
	return Vector{}
}

func (p *Plane) MaterialAt(a Vector) Material {
	return p.Material
}

func (p *Plane) NormalAt(a Vector) Vector {
	return p.Normal
}
