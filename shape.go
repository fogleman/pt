package pt

import (
	"math/rand"
)

type Shape interface {
	Intersect(Ray) float64
	Color() Color
	Normal(Vector) Vector
	RandomPoint(*rand.Rand) Vector
}
