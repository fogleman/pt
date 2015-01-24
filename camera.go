package pt

import (
	"math"
	"math/rand"
)

type Camera struct {
	Position Vector
	U, V, W Vector
	Scale float64
}

func (c *Camera) LookAt(eye, look, up Vector, fovy float64) {
	c.Position = eye
	c.W = look.Sub(eye).Normalize()
	c.U = up.Cross(c.W).Normalize()
	c.V = c.W.Cross(c.U).Normalize()
	c.Scale = 1 / math.Tan(fovy * math.Pi / 360)
}

func (c *Camera) CastRay(x, y, w, h int) Ray {
	aspect := float64(w) / float64(h)
	px := ((float64(x) + rand.Float64() - 0.5) / (float64(w) - 1)) * 2 - 1
	py := ((float64(y) + rand.Float64() - 0.5) / (float64(h) - 1)) * 2 - 1
	d := Vector{}
	d = d.Add(c.U.Mul(px * aspect))
	d = d.Add(c.V.Mul(-py))
	d = d.Add(c.W.Mul(c.Scale))
	return Ray{c.Position, d}
}
