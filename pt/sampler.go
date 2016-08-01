package pt

import (
	"math"
	"math/rand"
)

type Sampler interface {
	Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color
}

type DefaultSampler struct {
	HitSamples int
	Bounces    int
}

func (s *DefaultSampler) Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color {
	return s.sample(scene, ray, s.HitSamples, s.Bounces, rnd)
}

func (s *DefaultSampler) sample(scene *Scene, ray Ray, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := scene.Intersect(ray)
	if scene.visibility > 0 {
		t := math.Pow(rnd.Float64(), 0.5) * scene.visibility
		if t < hit.T {
			d := RandomUnitVector(rnd)
			o := ray.Position(t)
			newRay := Ray{o, d}
			return s.sample(scene, newRay, 1, depth-1, rnd)
		}
	}
	if !hit.Ok() {
		if scene.texture != nil {
			d := ray.Direction
			u := math.Atan2(d.Z, d.X)
			v := math.Atan2(d.Y, Vector{d.X, 0, d.Z}.Length())
			u = (u + math.Pi) / (2 * math.Pi)
			v = (v + math.Pi/2) / math.Pi
			return scene.texture.Sample(u, v)
		}
		return scene.color
	}
	info := hit.Info(ray)
	result := Color{}
	emittance := info.Material.Emittance
	if emittance > 0 {
		attenuation := info.Material.Attenuation.Compute(hit.T)
		result = result.Add(info.Color.MulScalar(emittance * attenuation * float64(samples)))
	}
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			fu := (float64(u) + rnd.Float64()) / float64(n)
			fv := (float64(v) + rnd.Float64()) / float64(n)
			newRay, reflected := ray.Bounce(&info, fu, fv, rnd)
			indirect := s.sample(scene, newRay, 1, depth-1, rnd)
			if reflected {
				tinted := indirect.Mix(info.Color.Mul(indirect), info.Material.Tint)
				result = result.Add(tinted)
			} else {
				result = result.Add(info.Color.Mul(indirect))
			}
		}
	}
	return result.DivScalar(float64(n * n))
}
