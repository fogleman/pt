package pt

import (
	"math"
)

type Ray struct {
	Origin, Direction Vector
}

func (n Ray) Reflect(i Ray) Ray {
	d := n.Direction.Reflect(i.Direction)
	return Ray{n.Origin, d}
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
	m1 := math.Sqrt(u)
	m2 := math.Sqrt(1 - u)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(r.Direction.MinAxis())
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.Mul(m1 * math.Cos(a)))
	d = d.Add(t.Mul(m1 * math.Sin(a)))
	d = d.Add(r.Direction.Mul(m2))
	return Ray{r.Origin, d}
}

func (r Ray) ConeBounce(theta, u, v float64) Ray {
	// TODO: make weighted
	theta = math.Acos(math.Cos(theta) + (1 - math.Cos(theta)) * u)
	m1 := math.Sin(theta)
	m2 := math.Cos(theta)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(r.Direction.MinAxis())
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.Mul(m1 * math.Cos(a)))
	d = d.Add(t.Mul(m1 * math.Sin(a)))
	d = d.Add(r.Direction.Mul(m2))
	return Ray{r.Origin, d}
}

func (r Ray) Bounce(i Ray, material Material, p, u, v float64) Ray {
	if p < material.Gloss {
		reflected := r.Reflect(i)
		return reflected.ConeBounce(material.Cone, u, v)
	} else {
		return r.WeightedBounce(u, v)
	}
}
