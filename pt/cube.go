package pt

import (
	"math"
	"math/rand"
)

type Cube struct {
	min      Vector
	max      Vector
	color    Color
	material Material
	box      Box
}

func NewCube(min, max Vector, color Color, material Material) Shape {
	box := Box{min, max}
	return &Cube{min, max, color, material, box}
}

func (c *Cube) Box() Box {
	return c.box
}

func (c *Cube) Intersect(r Ray) Hit {
	n := c.min.Sub(r.Origin).DivVector(r.Direction)
	f := c.max.Sub(r.Origin).DivVector(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 > 0 && t0 < t1 {
		return Hit{c, t0}
	}
	return NoHit
}

func (c *Cube) Color(p Vector) Color {
	return c.color
}

func (c *Cube) Material(p Vector) Material {
	return c.material
}

func (c *Cube) Normal(p Vector) Vector {
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

func (c *Cube) RandomPoint(rnd *rand.Rand) Vector {
	x := c.min.X + rnd.Float64()*(c.max.X-c.min.X)
	y := c.min.Y + rnd.Float64()*(c.max.Y-c.min.Y)
	z := c.min.Z + rnd.Float64()*(c.max.Z-c.min.Z)
	return Vector{x, y, z}
}
