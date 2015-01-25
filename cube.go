package pt

import (
	"math"
	"math/rand"
)

type Cube struct {
	Min Vector
	Max Vector
	Col Color
}

func (c *Cube) Intersect(r Ray) float64 {
	a := c.Min.Sub(r.Origin).DivVector(r.Direction)
	b := c.Max.Sub(r.Origin).DivVector(r.Direction)
	t1 := a.Min(b)
	t2 := a.Max(b)
	n := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	f := math.Min(math.Min(t2.X, t2.Y), t2.Z)
	if n > 0 && n < f {
		return n
	}
	return INF
}

func (c *Cube) Color() Color {
	return c.Col
}

func (c *Cube) Normal(p Vector) Vector {
	switch {
	case p.X < c.Min.X + EPS:
		return Vector{-1, 0, 0}
	case p.X > c.Max.X - EPS:
		return Vector{1, 0, 0}
	case p.Y < c.Min.Y + EPS:
		return Vector{0, -1, 0}
	case p.Y > c.Max.Y - EPS:
		return Vector{0, 1, 0}
	case p.Z < c.Min.Z + EPS:
		return Vector{0, 0, -1}
	case p.Z > c.Max.Z - EPS:
		return Vector{0, 0, 1}
	}
	return Vector{0, 1, 0}
}

func (c *Cube) RandomPoint() Vector {
	x := c.Min.X + rand.Float64() * (c.Max.X - c.Min.X)
	y := c.Min.Y + rand.Float64() * (c.Max.Y - c.Min.Y)
	z := c.Min.Z + rand.Float64() * (c.Max.Z - c.Min.Z)
	return Vector{x, y, z}
}
