package pt

import (
	"math"
	"math/rand"
	"sync/atomic"
)

type Scene struct {
	shapes []Shape
	lights []Shape
	tree   *Tree
	rays   uint64
}

func (s *Scene) Compile() {
	for _, shape := range s.shapes {
		if mesh, ok := shape.(*Mesh); ok {
			mesh.Compile()
		}
	}
	if s.tree == nil {
		s.tree = NewTree(s.shapes)
	}
}

func (s *Scene) Add(shape Shape) {
	s.shapes = append(s.shapes, shape)
	if shape.Material(Vector{}).Emittance > 0 {
		s.lights = append(s.lights, shape)
	}
}

func (s *Scene) RayCount() int {
	return int(atomic.LoadUint64(&s.rays))
}

func (s *Scene) Intersect(r Ray) Hit {
	atomic.AddUint64(&s.rays, 1)
	return s.tree.Intersect(r)
}

func (s *Scene) Shadow(r Ray, light Shape, max float64) bool {
	hit := s.Intersect(r)
	return hit.Shape != light && hit.T < max
}

func (s *Scene) DirectLight(n Ray, rnd *rand.Rand) Color {
	color := Color{}
	for _, light := range s.lights {
		p := light.RandomPoint(rnd)
		d := p.Sub(n.Origin)
		lr := Ray{n.Origin, d.Normalize()}
		diffuse := lr.Direction.Dot(n.Direction)
		if diffuse <= 0 {
			continue
		}
		distance := d.Length()
		if s.Shadow(lr, light, distance) {
			continue
		}
		material := light.Material(p)
		emittance := material.Emittance
		attenuation := material.Attenuation.Compute(distance)
		color = color.Add(light.Color(p).MulScalar(diffuse * emittance * attenuation))
	}
	return color.DivScalar(float64(len(s.lights)))
}

func (s *Scene) RecursiveSample(r Ray, reflected bool, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := s.Intersect(r)
	if !hit.Ok() {
		return Color{}
	}
	info := hit.Info(r)
	result := Color{}
	if reflected {
		result = result.Add(info.Color.MulScalar(info.Material.Emittance))
	}
	p, u, v := rnd.Float64(), rnd.Float64(), rnd.Float64()
	newRay, reflected := s.Bounce(r, info.Ray, &info, p, u, v)
	indirect := s.RecursiveSample(newRay, reflected, depth-1, rnd)
	if reflected {
		tinted := indirect.Mix(info.Color.Mul(indirect), info.Material.Tint)
		result = result.Add(tinted)
	} else {
		direct := s.DirectLight(info.Ray, rnd)
		result = result.Add(info.Color.Mul(direct.Add(indirect)))
	}
	return result
}

func (s *Scene) Sample(r Ray, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := s.Intersect(r)
	if !hit.Ok() {
		return Color{}
	}
	info := hit.Info(r)
	result := Color{}
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			p := rnd.Float64()
			fu := (float64(u) + rnd.Float64()) / float64(n)
			fv := (float64(v) + rnd.Float64()) / float64(n)
			newRay, reflected := s.Bounce(r, info.Ray, &info, p, fu, fv)
			indirect := s.RecursiveSample(newRay, reflected, depth-1, rnd)
			if reflected {
				tinted := indirect.Mix(info.Color.Mul(indirect), info.Material.Tint)
				result = result.Add(tinted)
			} else {
				direct := s.DirectLight(info.Ray, rnd)
				result = result.Add(info.Color.Mul(direct.Add(indirect)))
			}
			result = result.Add(info.Color.MulScalar(info.Material.Emittance))
		}
	}
	return result.DivScalar(float64(n * n))
}

func (s *Scene) Bounce(i, n Ray, info *HitInfo, p, u, v float64) (Ray, bool) {
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
