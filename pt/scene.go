package pt

import (
	"math"
	"math/rand"
)

type Scene struct {
	shapes    []Shape
	lights    []Shape
	shapeTree *Tree
	lightTree *Tree
}

func (s *Scene) Compile() {
	if s.shapeTree == nil {
		s.shapeTree = NewTree(s.shapes)
	}
	if s.lightTree == nil {
		s.lightTree = NewTree(s.lights)
	}
}

func (s *Scene) AddShape(shape Shape) {
	s.shapes = append(s.shapes, shape)
}

func (s *Scene) AddLight(shape Shape) {
	s.lights = append(s.lights, shape)
}

func (s *Scene) IntersectShapes(r Ray) Hit {
	return s.shapeTree.Intersect(r)
}

func (s *Scene) IntersectLights(r Ray) Hit {
	hit := s.lightTree.Intersect(r)
	if hit.Ok() {
		shapeHit := s.shapeTree.Intersect(r)
		if shapeHit.T < hit.T {
			return NoHit
		}
	}
	return hit
}

func (s *Scene) Shadow(r Ray, max float64) bool {
	hit := s.shapeTree.Intersect(r)
	return hit.T < max
}

func (s *Scene) DirectLight(n Ray, rnd *rand.Rand) Color {
	color := Color{}
	for _, light := range s.lights {
		p := light.RandomPoint(rnd)
		d := p.Sub(n.Origin)
		lr := Ray{n.Origin, d.Normalize()}
		if s.Shadow(lr, d.Length()) {
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
		hit := s.IntersectLights(r)
		if hit.Ok() {
			info := hit.Info(r)
			return info.Color
		}
	}
	hit := s.IntersectShapes(r)
	if !hit.Ok() {
		return Color{}
	}
	info := hit.Info(r)
	p, u, v := rnd.Float64(), rnd.Float64(), rnd.Float64()
	newRay, reflected := info.Ray.Bounce(r, info.Material, p, u, v)
	indirect := s.RecursiveSample(newRay, reflected, depth-1, rnd)
	if reflected {
		tinted := indirect.Mix(info.Color.MulColor(indirect), info.Material.Tint)
		return tinted
	} else {
		direct := s.DirectLight(info.Ray, rnd)
		return info.Color.MulColor(direct.Add(indirect))
	}
}

func (s *Scene) Sample(r Ray, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := s.IntersectShapes(r)
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
				tinted := indirect.Mix(info.Color.MulColor(indirect), info.Material.Tint)
				result = result.Add(tinted)
			} else {
				direct := s.DirectLight(info.Ray, rnd)
				result = result.Add(info.Color.MulColor(direct.Add(indirect)))
			}
		}
	}
	return result.Div(float64(n * n))
}
