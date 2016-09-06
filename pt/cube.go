package pt

import "math"

type Cube struct {
	Min      Vector
	Max      Vector
	Material Material
	Box      Box
}

func NewCube(min, max Vector, material Material) *Cube {
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

func (c *Cube) UV(p Vector) Vector {
	p = p.Sub(c.Min).Div(c.Max.Sub(c.Min))
	return Vector{p.X, p.Z, 0}
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

func (c *Cube) Mesh() *Mesh {
	a := c.Min
	b := c.Max
	z := Vector{}
	m := c.Material
	v000 := Vector{a.X, a.Y, a.Z}
	v001 := Vector{a.X, a.Y, b.Z}
	v010 := Vector{a.X, b.Y, a.Z}
	v011 := Vector{a.X, b.Y, b.Z}
	v100 := Vector{b.X, a.Y, a.Z}
	v101 := Vector{b.X, a.Y, b.Z}
	v110 := Vector{b.X, b.Y, a.Z}
	v111 := Vector{b.X, b.Y, b.Z}
	triangles := []*Triangle{
		NewTriangle(v000, v100, v110, z, z, z, m),
		NewTriangle(v000, v110, v010, z, z, z, m),
		NewTriangle(v001, v101, v111, z, z, z, m),
		NewTriangle(v001, v111, v011, z, z, z, m),
		NewTriangle(v000, v100, v101, z, z, z, m),
		NewTriangle(v000, v101, v001, z, z, z, m),
		NewTriangle(v010, v110, v111, z, z, z, m),
		NewTriangle(v010, v111, v011, z, z, z, m),
		NewTriangle(v000, v010, v011, z, z, z, m),
		NewTriangle(v000, v011, v001, z, z, z, m),
		NewTriangle(v100, v110, v111, z, z, z, m),
		NewTriangle(v100, v111, v101, z, z, z, m),
	}
	return NewMesh(triangles)
}
