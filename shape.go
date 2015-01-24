package pt

type Shape interface {
	Intersect(Ray) float64
	Color() Color
	Normal(Vector) Vector
	RandomPoint() Vector
}
