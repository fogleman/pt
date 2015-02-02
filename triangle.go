package pt

import (
	"math"
	"math/rand"
)

type Triangle struct {
	v1, v2, v3 Vector
	color      Color
	material   Material
	normal     Vector
	box        Box
}

func NewTriangle(v1, v2, v3 Vector, color Color, material Material) Shape {
	normal := v2.Sub(v1).Normalize().Cross(v3.Sub(v1).Normalize())
	min := v1.Min(v2).Min(v3)
	max := v1.Max(v2).Max(v3)
	box := Box{min, max}
	return &Triangle{v1, v2, v3, color, material, normal, box}
}

func (t *Triangle) Box() Box {
	return t.box
}

func (me *Triangle) Intersect(r Ray) float64 {
	edge1 := me.v2.Sub(me.v1)
	edge2 := me.v3.Sub(me.v1)
	pvec := r.Direction.Cross(edge2)
	det := edge1.Dot(pvec)
	if math.Abs(det) < EPS {
		return INF
	}
	inv := 1 / det
	tvec := r.Origin.Sub(me.v1)
	u := tvec.Dot(pvec) * inv
	if u < 0 || u > 1 {
		return INF
	}
	qvec := tvec.Cross(edge1)
	v := r.Direction.Dot(qvec) * inv
	if v < 0 || u+v > 1 {
		return INF
	}
	t := edge2.Dot(qvec) * inv
	if t < EPS {
		return INF
	}
	return t
}

func (t *Triangle) Color(v Vector) Color {
	return t.color
}

func (t *Triangle) Material(v Vector) Material {
	return t.material
}

func (t *Triangle) Normal(v Vector) Vector {
	return t.normal
}

func (t *Triangle) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{} // TODO: fix
}
