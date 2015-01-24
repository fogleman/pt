package pt

import (
	"math"
	"math/rand"
)

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) CosineBounce() Ray {
	rx := rand.Float64() * 2 * math.Pi
	ry := rand.Float64() * 2 * math.Pi
	x := math.Sin(rx) * math.Sin(ry)
	y := math.Sin(rx) * math.Cos(ry)
	z := math.Cos(rx)
	v := Vector{x, y, z}
	if v.Dot(r.Direction) < 0 {
		v = v.Mul(-1)
	}
	return Ray{r.Origin, v}
}
