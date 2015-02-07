package pt

import (
	"math"
)

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Position(t float64) Vector {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func (n Ray) Reflect(i Ray) Ray {
	return Ray{n.Origin, n.Direction.Reflect(i.Direction)}
}

func (n Ray) Reflectance(i Ray, n1, n2 float64) float64 {
	nr := n1 / n2
	cosI := -n.Direction.Dot(i.Direction)
	sinT2 := nr * nr * (1 - cosI*cosI)
	if sinT2 > 1 {
		return 1
	}
	cosT := math.Sqrt(1 - sinT2)
	rOrth := (n1*cosI - n2*cosT) / (n1*cosI + n2*cosT)
	rPar := (n2*cosI - n1*cosT) / (n2*cosI + n1*cosT)
	return (rOrth*rOrth + rPar*rPar) / 2
}

func (r Ray) WeightedBounce(u, v float64) Ray {
	m1 := math.Sqrt(u)
	m2 := math.Sqrt(1 - u)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(r.Direction.MinAxis())
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.MulScalar(m1 * math.Cos(a)))
	d = d.Add(t.MulScalar(m1 * math.Sin(a)))
	d = d.Add(r.Direction.MulScalar(m2))
	return Ray{r.Origin, d}
}

func (r Ray) ConeBounce(theta, u, v float64) Ray {
	if theta < EPS {
		return r
	}
	theta = theta * (1 - (2 * math.Acos(u) / math.Pi))
	m1 := math.Sin(theta)
	m2 := math.Cos(theta)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(r.Direction.MinAxis())
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.MulScalar(m1 * math.Cos(a)))
	d = d.Add(t.MulScalar(m1 * math.Sin(a)))
	d = d.Add(r.Direction.MulScalar(m2))
	return Ray{r.Origin, d}
}

func (r Ray) Bounce(i Ray, material Material, p, u, v float64) (Ray, bool) {
	if p < r.Reflectance(i, 1, material.Index) {
		reflected := r.Reflect(i)
		return reflected.ConeBounce(material.Gloss, u, v), true
	} else {
		return r.WeightedBounce(u, v), false
	}
}
