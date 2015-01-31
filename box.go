package pt

import (
	"math"
)

type Box struct {
	Min, Max Vector
}

func (a Box) Extend(b Box) Box {
	return Box{a.Min.Min(b.Min), a.Max.Max(b.Max)}
}

func (b *Box) Intersect(r Ray) (float64, float64) {
	n := b.Min.Sub(r.Origin).DivVector(r.Direction)
	f := b.Max.Sub(r.Origin).DivVector(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	return t0, t1
}
