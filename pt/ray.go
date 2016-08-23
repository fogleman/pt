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
	radius := math.Sqrt(u)
	theta := 2 * math.Pi * v
	s := r.Direction.Cross(RandomUnitVector(rnd)).Normalize()
	t := r.Direction.Cross(s)
	d := Vector{}
	d = d.Add(s.MulScalar(radius * math.Cos(theta)))
	d = d.Add(t.MulScalar(radius * math.Sin(theta)))
	d = d.Add(r.Direction.MulScalar(math.Sqrt(1 - u)))
	return Ray{r.Origin, d}
}

func (r Ray) ConeBounce(theta, u, v float64, rnd *rand.Rand) Ray {
	return Ray{r.Origin, Cone(r.Direction, theta, u, v, rnd)}
}

func (i Ray) Bounce(info *HitInfo, u, v float64, mode int, rnd *rand.Rand) (Ray, bool, float64) {
	n := info.Ray
	material := info.Material
	n1, n2 := 1.0, material.Index
	if info.Inside {
		n1, n2 = n2, n1
	}
	var p float64
	if material.Reflectivity >= 0 {
		p = material.Reflectivity
	} else {
		p = n.Reflectance(i, n1, n2)
	}
	var reflect bool
	switch mode {
	case -1:
		reflect = rnd.Float64() < p
	case 0:
		reflect = false
	case 1:
		reflect = true
	}
	if reflect {
		reflected := n.Reflect(i)
		return reflected.ConeBounce(material.Gloss, u, v, rnd), true, p
	} else if material.Transparent {
		refracted := n.Refract(i, n1, n2)
		return refracted.ConeBounce(material.Gloss, u, v, rnd), true, 1 - p
	} else {
		return n.WeightedBounce(u, v, rnd), false, 1 - p
	}
}
