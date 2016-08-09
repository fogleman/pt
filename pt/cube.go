package pt

import "math"

type Cube struct {
	Min      Vector
	Max      Vector
	Material Material
	Box      Box
}

func NewCube(min, max Vector, material Material) Shape {
	box := Box{min, max}
	return &Cube{min, max, material, box}
}

func (c *Cube) Compile() {
}

func (c *Cube) BoundingBox() Box {
	return c.Box
}

func (c *Cube) Intersect(r Ray) Hit {
	n := c.Min.Sub(r.Origin).Div(r.Direction)
	f := c.Max.Sub(r.Origin).Div(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 > 0 && t0 < t1 {
		return Hit{c, t0, nil}
	}
	return NoHit
}

func (c *Cube) ColorAt(p Vector) Color {
	if c.Material.Texture == nil {
		return c.Material.Color
	}
	p = p.Sub(c.Min).Div(c.Max.Sub(c.Min))
	return c.Material.Texture.Sample(p.X, p.Z)
}

func (c *Cube) MaterialAt(p Vector) Material {
	return c.Material
}

func (c *Cube) NormalAt(p Vector) Vector {
	switch {
	case p.X < c.Min.X+EPS:
		return Vector{-1, 0, 0}
	case p.X > c.Max.X-EPS:
		return Vector{1, 0, 0}
	case p.Y < c.Min.Y+EPS:
		return Vector{0, -1, 0}
	case p.Y > c.Max.Y-EPS:
		return Vector{0, 1, 0}
	case p.Z < c.Min.Z+EPS:
		return Vector{0, 0, -1}
	case p.Z > c.Max.Z-EPS:
		return Vector{0, 0, 1}
	}
	return Vector{0, 1, 0}
}
