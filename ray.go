package pt

import (
	"math"
)

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) UniformBounce(u, v float64) Ray {
	rx := u * 2 * math.Pi
	ry := v * 2 * math.Pi
	x := math.Sin(rx) * math.Sin(ry)
	y := math.Sin(rx) * math.Cos(ry)
	z := math.Cos(rx)
	d := Vector{x, y, z}
	if d.Dot(r.Direction) < 0 {
		d = d.Mul(-1)
	}
	return Ray{r.Origin, d}
}

func (r Ray) WeightedBounce(u, v float64) Ray {
	m := math.Sqrt(u)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(Vector{0, 1, 0})
	if math.Abs(r.Direction.X) < 0.5 {
		s = r.Direction.Cross(Vector{1, 0, 0})
	}
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.Mul(m * math.Cos(a)))
	d = d.Add(t.Mul(m * math.Sin(a)))
	d = d.Add(r.Direction.Mul(math.Sqrt(1 - u)))
	return Ray{r.Origin, d}
}
