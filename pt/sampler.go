package pt

import (
	"math"
	"math/rand"
)

type Sampler interface {
	Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color
}

/*
pick a random light
cast from camera N times, store hits
cast from light M times, store hits
cast shadow ray for all pairs
compute probability based on cos a, cos a
combine
*/

type DefaultSampler struct {
	HitSamples int
	Bounces    int
}

func (s *DefaultSampler) Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color {
	return s.sample(scene, ray, true, s.HitSamples, s.Bounces, rnd)
}

func (s *DefaultSampler) sample(scene *Scene, ray Ray, emission bool, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := scene.Intersect(ray)
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
	if emission && emittance > 0 {
		attenuation := info.Material.Attenuation.Compute(hit.T)
		result = result.Add(info.Color.MulScalar(emittance * attenuation * float64(samples)))
	}
	n := int(math.Sqrt(float64(samples)))
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			fu := (float64(u) + rnd.Float64()) / float64(n)
			fv := (float64(v) + rnd.Float64()) / float64(n)
			newRay, reflected := ray.Bounce(&info, fu, fv, rnd)
			indirect := s.sample(scene, newRay, reflected, 1, depth-1, rnd)
			if reflected {
				tinted := indirect.Mix(info.Color.Mul(indirect), info.Material.Tint)
				result = result.Add(tinted)
			} else {
				direct := s.directLight(scene, info.Ray, rnd)
				result = result.Add(info.Color.Mul(direct.Add(indirect)))
			}
		}
	}
	return result.DivScalar(float64(n * n))
}

func (s *DefaultSampler) directLight(scene *Scene, n Ray, rnd *rand.Rand) Color {
	nLights := len(scene.lights)
	if nLights == 0 {
		return Color{}
	}
	light := scene.lights[rand.Intn(nLights)]
	// TODO: get bounding sphere, get random point, see if hits shape
	p := light.RandomPoint(rnd)
	d := p.Sub(n.Origin)
	lr := Ray{n.Origin, d.Normalize()}
	diffuse := lr.Direction.Dot(n.Direction)
	if diffuse <= 0 {
		return Color{}
	}
	// distance := light.(*Sphere).center.Sub(n.Origin).Length()
	distance := light.Intersect(lr).T
	// distance := d.Length()
	if scene.Shadow(lr, light, distance) {
		return Color{}
	}
	material := light.Material(p)
	emittance := material.Emittance
	attenuation := material.Attenuation.Compute(distance)
	color := light.Color(p).MulScalar(diffuse * emittance * attenuation)

	// r := light.Box().Radius()
	r := light.(*Sphere).radius
	a := math.Atan2(r, distance)
	f := 1 - math.Cos(a)
	return color.MulScalar(f * float64(nLights))
}
