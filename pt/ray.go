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

func (n Ray) Refract(i Ray, n1, n2 float64) Ray {
	return Ray{n.Origin, n.Direction.Refract(i.Direction, n1, n2)}
}

func (n Ray) Reflectance(i Ray, n1, n2 float64) float64 {
	return n.Direction.Reflectance(i.Direction, n1, n2)
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

func (i Ray) Bounce(info *HitInfo, p, u, v float64) (Ray, bool) {
	n := info.Ray
	n1, n2 := 1.0, info.Material.Index
	if info.Inside {
		n1, n2 = n2, n1
	}
	if p < n.Reflectance(i, n1, n2) {
		reflected := n.Reflect(i)
		return reflected.ConeBounce(info.Material.Gloss, u, v), true
	} else if info.Material.Transparent {
		refracted := n.Refract(i, n1, n2)
		return refracted.ConeBounce(info.Material.Gloss, u, v), true
	} else {
		return n.WeightedBounce(u, v), false
	}
}
