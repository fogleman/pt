package pt

import (
	"math"
	"math/rand"
)

type Scene struct {
	shapes []Shape
	lights []Shape
	tree   *Tree
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

func (s *Scene) Intersect(r Ray) Hit {
	return s.tree.Intersect(r)
}

func (s *Scene) Shadow(r Ray, light Shape, max float64) bool {
	hit := s.tree.Intersect(r)
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
	newRay, reflected := info.Ray.Bounce(r, info.Material, p, u, v)
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
			newRay, reflected := info.Ray.Bounce(r, info.Material, p, fu, fv)
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
