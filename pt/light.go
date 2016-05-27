package pt

import "math"

type DirectionalLight struct {
	Direction Vector
	Color     Color
}

func (d *DirectionalLight) ColorForRay(r Ray) Color {
	return d.Color.MulScalar(math.Max(0, r.Direction.Dot(d.Direction)))
}
