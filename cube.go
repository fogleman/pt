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
}

func NewCube(min, max Vector, color Color, material Material) Shape {
	return &Cube{min, max, color, material}
}

func (c *Cube) Intersect(r Ray) float64 {
	a := c.min.Sub(r.Origin).DivVector(r.Direction)
	b := c.max.Sub(r.Origin).DivVector(r.Direction)
	t1 := a.Min(b)
	t2 := a.Max(b)
	n := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	f := math.Min(math.Min(t2.X, t2.Y), t2.Z)
	if n > 0 && n < f {
		return n
	}
	return INF
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
