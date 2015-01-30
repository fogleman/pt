package pt

import (
	"math"
	"math/rand"
)

type Scene struct {
	shapes []Shape
	lights []Shape
}

func (s *Scene) AddShape(shape Shape) {
	s.shapes = append(s.shapes, shape)
}

func (s *Scene) AddLight(shape Shape) {
	s.lights = append(s.lights, shape)
}

func (s *Scene) Intersect(r Ray, shapes []Shape) (Hit, bool) {
	hit := Hit{}
	u := INF
	for _, shape := range shapes {
		t := shape.Intersect(r)
		if t < u {
			u = t
			p := r.Position(t)
			n := shape.Normal(p)
			hit = Hit{shape, Ray{p, n}}
		}
	}
	ok := u < INF
	return hit, ok
}

func (s *Scene) Shadow(r Ray) bool {
	// TODO: ignore objects behind the light source
	for _, shape := range s.shapes {
		t := shape.Intersect(r)
		if t < INF {
			return true
		}
	}
	return false
}

func (s *Scene) DirectLight(i, n Ray, rnd *rand.Rand) Color {
	color := Color{}
	for _, light := range s.lights {
		p := light.RandomPoint(rnd)
		lr := Ray{n.Origin, p.Sub(n.Origin).Normalize()}
		if s.Shadow(lr) {
			continue
		}
		diffuse := math.Max(0, lr.Direction.Dot(n.Direction))
		color = color.Add(light.Color(p).Mul(diffuse))
	}
	return color.Div(float64(len(s.lights)))
}

func (s *Scene) RecursiveSample(r Ray, reflected bool, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	if reflected {
		hit, ok := s.Intersect(r, s.lights)
		if ok {
			return hit.Shape.Color(hit.Ray.Origin)
		}
	}
	hit, ok := s.Intersect(r, s.shapes)
	if !ok {
		return Color{}
	}
	shape := hit.Shape
	color := shape.Color(hit.Ray.Origin)
	material := shape.Material(hit.Ray.Origin)
	direct := s.DirectLight(r, hit.Ray, rnd)
	direct = direct.Mul(1 - material.Gloss)
	p, u, v := rnd.Float64(), rnd.Float64(), rnd.Float64()
	ray, reflected := hit.Ray.Bounce(r, material, p, u, v)
	indirect := s.RecursiveSample(ray, reflected, depth-1, rnd)
	if reflected {
		return indirect
	} else {
		return color.MulColor(direct.Add(indirect))
	}
}

func (s *Scene) Sample(r Ray, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit, ok := s.Intersect(r, s.shapes)
	if !ok {
		return Color{}
	}
	shape := hit.Shape
	color := shape.Color(hit.Ray.Origin)
	material := shape.Material(hit.Ray.Origin)
	result := Color{}
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			direct := s.DirectLight(r, hit.Ray, rnd)
			direct = direct.Mul(1 - material.Gloss)
			p := rnd.Float64()
			fu := (float64(u) + rnd.Float64()) * (1 / float64(n))
			fv := (float64(v) + rnd.Float64()) * (1 / float64(n))
			ray, reflected := hit.Ray.Bounce(r, material, p, fu, fv)
			indirect := s.RecursiveSample(ray, reflected, depth-1, rnd)
			if reflected {
				result = result.Add(indirect)
			} else {
				result = result.Add(color.MulColor(direct.Add(indirect)))
			}
		}
	}
	return result.Div(float64(n * n))
}
