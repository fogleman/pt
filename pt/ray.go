package pt

import (
	"math"
	"math/rand"
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

func (r Ray) WeightedBounce(u, v float64, rnd *rand.Rand) Ray {
	m1 := math.Sqrt(u)
	m2 := math.Sqrt(1 - u)
	a := v * 2 * math.Pi
	s := r.Direction.Cross(RandomUnitVector(rnd))
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.MulScalar(m1 * math.Cos(a)))
	d = d.Add(t.MulScalar(m1 * math.Sin(a)))
	d = d.Add(r.Direction.MulScalar(m2))
	return Ray{r.Origin, d.Normalize()}
}

func (r Ray) ConeBounce(theta, u, v float64, rnd *rand.Rand) Ray {
	return Ray{r.Origin, Cone(r.Direction, theta, u, v, rnd)}
}

func (i Ray) Bounce(info *HitInfo, u, v float64, rnd *rand.Rand) (Ray, bool) {
	n := info.Ray
	n1, n2 := 1.0, info.Material.Index
	if info.Inside {
		n1, n2 = n2, n1
	}
	if rnd.Float64() < n.Reflectance(i, n1, n2) {
		reflected := n.Reflect(i)
		return reflected.ConeBounce(info.Material.Gloss, u, v, rnd), true
	} else if info.Material.Transparent {
		refracted := n.Refract(i, n1, n2)
		return refracted.ConeBounce(info.Material.Gloss, u, v, rnd), true
	} else {
		return n.WeightedBounce(u, v, rnd), false
	}
}
