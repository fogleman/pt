package pt

type Shape interface {
	Intersect(r Ray) float64
	Color() Color
}
