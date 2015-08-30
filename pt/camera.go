package pt

import (
	"math"
)

type Camera struct {
	p, u, v, w Vector
	m          float64
}

func LookAt(eye, look, up Vector, fovy float64) Camera {
	c := Camera{}
	c.p = eye
	c.w = look.Sub(eye).Normalize()
	c.u = up.Cross(c.w).Normalize()
	c.v = c.w.Cross(c.u).Normalize()
	c.m = 1 / math.Tan(fovy*math.Pi/360)
	return c
}

func (c *Camera) CastRay(x, y, w, h int, u, v float64) Ray {
	aspect := float64(w) / float64(h)
	px := ((float64(x)+u-0.5)/(float64(w)-1))*2 - 1
	py := ((float64(y)+v-0.5)/(float64(h)-1))*2 - 1
	d := Vector{}
	d = d.Add(c.u.MulScalar(-px * aspect))
	d = d.Add(c.v.MulScalar(-py))
	d = d.Add(c.w.MulScalar(c.m))
	d = d.Normalize()
	return Ray{c.p, d}
}
