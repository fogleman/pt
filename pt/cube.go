package pt

import "math"

type Cube struct {
	min      Vector
	max      Vector
	material Material
	box      Box
}

func NewCube(min, max Vector, material Material) Shape {
	box := Box{min, max}
	return &Cube{min, max, material, box}
}

func (c *Cube) Compile() {
}

func (c *Cube) BoundingBox() Box {
	return c.box
}

func (c *Cube) Intersect(r Ray) Hit {
	n := c.min.Sub(r.Origin).Div(r.Direction)
	f := c.max.Sub(r.Origin).Div(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 > 0 && t0 < t1 {
		return Hit{c, t0, nil}
	}
	return NoHit
}

func (c *Cube) ColorAt(p Vector) Color {
	if c.material.Texture == nil {
		return c.material.Color
	}
	p = p.Sub(c.min).Div(c.max.Sub(c.min))
	return c.material.Texture.Sample(p.X, p.Z)
}

func (c *Cube) MaterialAt(p Vector) Material {
	return c.material
}

func (c *Cube) NormalAt(p Vector) Vector {
	switch {
	case p.X < c.min.X+EPS:
		return Vector{-1, 0, 0}
	case p.X > c.max.X-EPS:
		return Vector{1, 0, 0}
	case p.Y < c.min.Y+EPS:
		return Vector{0, -1, 0}
	case p.Y > c.max.Y-EPS:
		return Vector{0, 1, 0}
	case p.Z < c.min.Z+EPS:
		return Vector{0, 0, -1}
	case p.Z > c.max.Z-EPS:
		return Vector{0, 0, 1}
	}
	return Vector{0, 1, 0}
}
