package pt

import (
	"math"
	"math/rand"
	"sync/atomic"
)

type Scene struct {
	color      Color
	texture    Texture
	visibility float64
	shapes     []Shape
	lights     []Shape
	tree       *Tree
	rays       uint64
}

func (s *Scene) SetColor(color Color) {
	s.color = color
}

func (s *Scene) SetTexture(texture Texture) {
	s.texture = texture
}

func (s *Scene) SetVisibility(visibility float64) {
	s.visibility = visibility
}

func (s *Scene) Compile() {
	for _, shape := range s.shapes {
		shape.Compile()
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

func (s *Scene) RayCount() uint64 {
	return atomic.LoadUint64(&s.rays)
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
	if len(s.lights) == 0 {
		return Color{}
	}
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

func (s *Scene) Sample(r Ray, emission bool, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := s.Intersect(r)
	if s.visibility > 0 {
		t := math.Pow(rnd.Float64(), 0.5) * s.visibility
		if t < hit.T {
			x := rnd.Float64() - 0.5
			y := rnd.Float64() - 0.5
			z := rnd.Float64() - 0.5
			d := Vector{x, y, z}
			d = d.Normalize()
			o := r.Position(t)
			newRay := Ray{o, d}
			return s.Sample(newRay, false, 1, depth-1, rnd)
		}
	}
	if !hit.Ok() {
		if s.texture != nil {
			d := r.Direction
			u := math.Atan2(d.Z, d.X)
			v := math.Atan2(d.Y, Vector{d.X, 0, d.Z}.Length())
			u = (u + math.Pi) / (2 * math.Pi)
			v = (v + math.Pi/2) / math.Pi
			return s.texture.Sample(u, v).MulScalar(5)
		}
		return s.color
	}
	info := hit.Info(r)
	result := Color{}
	if emission {
		emittance := info.Material.Emittance
		if emittance > 0 {
			attenuation := info.Material.Attenuation.Compute(hit.T)
			result = result.Add(info.Color.MulScalar(emittance * attenuation * float64(samples)))
		}
	}
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			p := rnd.Float64()
			fu := (float64(u) + rnd.Float64()) / float64(n)
			fv := (float64(v) + rnd.Float64()) / float64(n)
			newRay, reflected := r.Bounce(&info, p, fu, fv)
			indirect := s.Sample(newRay, reflected, 1, depth-1, rnd)
			if reflected {
				tinted := indirect.Mix(info.Color.Mul(indirect), info.Material.Tint)
				result = result.Add(tinted)
			} else {
				direct := s.DirectLight(info.Ray, rnd)
				result = result.Add(info.Color.Mul(direct.Add(indirect)))
			}
		}
	}
	return result.DivScalar(float64(n * n))
}
