package pt

import (
	"math"
	"math/rand"
)

type Triangle struct {
	material   *Material
	box        Box
	v1, v2, v3 Vector
	n1, n2, n3 Vector
	t1, t2, t3 Vector
}

func (t *Triangle) Compile() {
}

func (t *Triangle) Box() Box {
	return t.box
}

func (t *Triangle) Intersect(r Ray) Hit {
	edge1 := t.v2.Sub(t.v1)
	edge2 := t.v3.Sub(t.v1)
	pvec := r.Direction.Cross(edge2)
	det := edge1.Dot(pvec)
	if math.Abs(det) < EPS {
		return NoHit
	}
	inv := 1 / det
	tvec := r.Origin.Sub(t.v1)
	u := tvec.Dot(pvec) * inv
	if u < 0 || u > 1 {
		return NoHit
	}
	qvec := tvec.Cross(edge1)
	v := r.Direction.Dot(qvec) * inv
	if v < 0 || u+v > 1 {
		return NoHit
	}
	d := edge2.Dot(qvec) * inv
	if d < EPS {
		return NoHit
	}
	return Hit{t, d}
}

func (t *Triangle) Color(p Vector) Color {
	if t.material.Texture == nil {
		return t.material.Color
	}
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.t1.MulScalar(u))
	n = n.Add(t.t2.MulScalar(v))
	n = n.Add(t.t3.MulScalar(w))
	return t.material.Texture.Sample(n.X, n.Y)
}

func (t *Triangle) Material(p Vector) Material {
	return *t.material
}

func (t *Triangle) Normal(p Vector) Vector {
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.n1.MulScalar(u))
	n = n.Add(t.n2.MulScalar(v))
	n = n.Add(t.n3.MulScalar(w))
	n = n.Normalize() // needed?
	return n
}

func (t *Triangle) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{} // TODO: fix
}

func (t *Triangle) Area() float64 {
	e1 := t.v2.Sub(t.v1)
	e2 := t.v3.Sub(t.v1)
	n := e1.Cross(e2)
	return n.Length() / 2
}

func (t *Triangle) Barycentric(p Vector) (u, v, w float64) {
	v0 := t.v2.Sub(t.v1)
	v1 := t.v3.Sub(t.v1)
	v2 := p.Sub(t.v1)
	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)
	d := d00*d11 - d01*d01
	v = (d11*d20 - d01*d21) / d
	w = (d00*d21 - d01*d20) / d
	u = 1 - v - w
	return
}

func (t *Triangle) UpdateBox() {
	min := t.v1.Min(t.v2).Min(t.v3)
	max := t.v1.Max(t.v2).Max(t.v3)
	t.box = Box{min, max}
}

func (t *Triangle) FixNormals() {
	e1 := t.v2.Sub(t.v1)
	e2 := t.v3.Sub(t.v1)
	n := e1.Cross(e2).Normalize()
	zero := Vector{}
	if t.n1 == zero {
		t.n1 = n
	}
	if t.n2 == zero {
		t.n2 = n
	}
	if t.n3 == zero {
		t.n3 = n
	}
}
