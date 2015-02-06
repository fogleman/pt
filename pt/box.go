package pt

import (
	"math"
)

type Box struct {
	Min, Max Vector
}

func BoxForShapes(shapes []Shape) Box {
	if len(shapes) == 0 {
		return Box{}
	}
	box := shapes[0].Box()
	for _, shape := range shapes {
		box = box.Extend(shape.Box())
	}
	return box
}

func BoxForTriangles(shapes []*Triangle) Box {
	if len(shapes) == 0 {
		return Box{}
	}
	box := shapes[0].Box()
	for _, shape := range shapes {
		box = box.Extend(shape.Box())
	}
	return box
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

func (b *Box) Partition(axis Axis, point float64) (left, right bool) {
	switch axis {
	case AxisX:
		left = b.Min.X <= point
		right = b.Max.X >= point
	case AxisY:
		left = b.Min.Y <= point
		right = b.Max.Y >= point
	case AxisZ:
		left = b.Min.Z <= point
		right = b.Max.Z >= point
	}
	return
}
