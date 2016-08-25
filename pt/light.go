package pt

import "math"

type DirectionalLight struct {
	Color     Color
	Direction Vector
	Emission  float64
	Exponent  float64
}

func NewDirectionalLight(color Color, direction Vector, emission, exponent float64) *DirectionalLight {
	return &DirectionalLight{color, direction.Normalize(), emission, exponent}
}

func (d *DirectionalLight) ColorForRay(r Ray) Color {
	p := math.Max(0, math.Pow(r.Direction.Dot(d.Direction), d.Exponent))
	return d.Color.MulScalar(p * d.Emission)
}
