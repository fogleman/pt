package pt

import (
	"math/rand"
)

type Shape interface {
	Intersect(Ray) float64
	Color() Color
	Material() Material
	Normal(Vector) Vector
	RandomPoint(*rand.Rand) Vector
}
