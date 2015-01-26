package pt

import (
	"math"
	"math/rand"
)

type Scene struct {
	Shapes []Shape
	Lights []Shape
}

func (s *Scene) AddShape(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) AddLight(light Shape) {
	s.Lights = append(s.Lights, light)
}

func (s *Scene) Intersect(r Ray) (Hit, bool) {
	hit := Hit{}
	u := INF
	for _, shape := range s.Shapes {
		t := shape.Intersect(r)
		if t < u {
			u = t
			p := r.Origin.Add(r.Direction.Mul(t))
			n := shape.Normal(p)
			hit = Hit{shape, Ray{p, n}}
		}
	}
	ok := u < INF
	return hit, ok
}

func (s *Scene) Shadow(r Ray) bool {
	// TODO: ignore objects behind the light source
	for _, shape := range s.Shapes {
		t := shape.Intersect(r)
		if t < INF {
			return true
		}
	}
	return false
}

func (s *Scene) Light(r Ray, rnd *rand.Rand) Color {
	color := Color{}
	for _, light := range s.Lights {
		lr := Ray{r.Origin, light.RandomPoint(rnd).Sub(r.Origin).Normalize()}
		if (s.Shadow(lr)) {
			continue
		}
		diffuse := math.Max(0, lr.Direction.Dot(r.Direction))
		color = color.Add(light.Color().Mul(diffuse))
	}
	return color.Div(float64(len(s.Lights)))
}

func (s *Scene) RecursiveSample(r Ray, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit, ok := s.Intersect(r)
	if !ok {
		return Color{}
	}
	color := hit.Shape.Color()
	direct := s.Light(hit.Ray, rnd).Mul(1 - hit.Shape.Material().Gloss)
	p, u, v := rnd.Float64(), rnd.Float64(), rnd.Float64()
	ray := hit.Ray.Bounce(r, hit.Shape.Material(), p, u, v)
	indirect := s.RecursiveSample(ray, depth - 1, rnd)
	return color.MulColor(direct.Add(indirect))
}

func (s *Scene) Sample(r Ray, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit, ok := s.Intersect(r)
	if !ok {
		return Color{}
	}
	result := Color{}
	color := hit.Shape.Color()
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			direct := s.Light(hit.Ray, rnd).Mul(1 - hit.Shape.Material().Gloss)
			p := rnd.Float64()
			fu := (float64(u) + rnd.Float64()) * (1 / float64(n))
			fv := (float64(v) + rnd.Float64()) * (1 / float64(n))
			ray := hit.Ray.Bounce(r, hit.Shape.Material(), p, fu, fv)
			indirect := s.RecursiveSample(ray, depth - 1, rnd)
			result = result.Add(color.MulColor(direct.Add(indirect)))
		}
	}
	return result.Div(float64(n * n))
}