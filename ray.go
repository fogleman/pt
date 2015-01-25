package pt

import (
	"math"
	"math/rand"
)

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) UniformBounce() Ray {
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

func (r Ray) WeightedBounce() Ray {
	u := rand.Float64()
	m := math.Sqrt(u)
	a := 2 * math.Pi * rand.Float64()
	s := r.Direction.Cross(Vector{0, 1, 0})
	if math.Abs(r.Direction.X) < 0.5 {
		s = r.Direction.Cross(Vector{1, 0, 0})
	}
	t := r.Direction.Cross(s)
	v := Vector{}
	v = v.Add(s.Mul(m * math.Cos(a)))
	v = v.Add(t.Mul(m * math.Sin(a)))
	v = v.Add(r.Direction.Mul(math.Sqrt(1 - u)))
	return Ray{r.Origin, v}
}
