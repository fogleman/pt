package pt

type Camera struct {
	Position, U, V, W Vector
}

func (c *Camera) LookAt(eye, look, up Vector) {
	c.Position = eye
	c.W = look.Sub(eye).Normalize()
	c.U = up.Cross(c.W).Normalize()
	c.V = c.W.Cross(c.U).Normalize()
}

func (c *Camera) CastRay(x, y, w, h int) Ray {
	aspect := float64(w) / float64(h)
	px := (float64(x) / (float64(w) - 1)) * 2 - 1
	py := (float64(y) / (float64(h) - 1)) * 2 - 1
	d := Vector{}
	d = d.Add(c.U.Mul(px * aspect))
	d = d.Add(c.V.Mul(py))
	d = d.Add(c.W.Mul(2.5))
	return Ray{c.Position, d}
}
