package pt

import (
	"math"
	"math/rand"
)

type Camera struct {
	p, u, v, w     Vector
	m              float64
	focalDistance  float64
	apertureRadius float64
}

func LookAt(eye, center, up Vector, fovy float64) Camera {
	c := Camera{}
	c.p = eye
	c.w = center.Sub(eye).Normalize()
	c.u = up.Cross(c.w).Normalize()
	c.v = c.w.Cross(c.u).Normalize()
	c.m = 1 / math.Tan(fovy*math.Pi/360)
	return c
}

func (c *Camera) SetFocus(focalPoint Vector, apertureRadius float64) {
	c.focalDistance = focalPoint.Sub(c.p).Length()
	c.apertureRadius = apertureRadius
}

func (c *Camera) CastRay(x, y, w, h int, u, v float64, rnd *rand.Rand) Ray {
	aspect := float64(w) / float64(h)
	px := ((float64(x)+u-0.5)/(float64(w)-1))*2 - 1
	py := ((float64(y)+v-0.5)/(float64(h)-1))*2 - 1
	d := Vector{}
	d = d.Add(c.u.MulScalar(-px * aspect))
	d = d.Add(c.v.MulScalar(-py))
	d = d.Add(c.w.MulScalar(c.m))
	d = d.Normalize()
	p := c.p
	if c.apertureRadius > 0 {
		focalPoint := c.p.Add(d.MulScalar(c.focalDistance))
		angle := rnd.Float64() * 2 * math.Pi
		radius := rnd.Float64() * c.apertureRadius
		p = p.Add(c.u.MulScalar(math.Cos(angle) * radius))
		p = p.Add(c.v.MulScalar(math.Sin(angle) * radius))
		d = focalPoint.Sub(p).Normalize()
	}
	return Ray{p, d}
}
