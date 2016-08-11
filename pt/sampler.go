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
compute probability based on cos a, cos a?
combine
*/

func NewSampler(firstHitSamples, maxBounces int) Sampler {
	return &DefaultSampler{firstHitSamples, maxBounces, true}
}

type DefaultSampler struct {
	FirstHitSamples int
	MaxBounces      int
	DirectLighting  bool
}

func (s *DefaultSampler) Sample(scene *Scene, ray Ray, rnd *rand.Rand) Color {
	return s.sample(scene, ray, true, s.FirstHitSamples, s.MaxBounces, rnd)
}

func (s *DefaultSampler) sample(scene *Scene, ray Ray, emission bool, samples, depth int, rnd *rand.Rand) Color {
	if depth < 0 {
		return Color{}
	}
	hit := scene.Intersect(ray)
	if !hit.Ok() {
		if scene.Texture != nil {
			d := ray.Direction
			u := math.Atan2(d.Z, d.X) + scene.TextureAngle
			v := math.Atan2(d.Y, Vector{d.X, 0, d.Z}.Length())
			u = (u + math.Pi) / (2 * math.Pi)
			v = (v + math.Pi/2) / math.Pi
			return scene.Texture.Sample(u, v)
		}
		return scene.Color
	}
	info := hit.Info(ray)
	result := Color{}
	emittance := info.Material.Emittance
	if emittance > 0 {
		if s.DirectLighting && !emission {
			return Color{}
		}
		result = result.Add(info.Color.MulScalar(emittance * float64(samples)))
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
				direct := Color{}
				if s.DirectLighting {
					direct = s.directLight(scene, info.Ray, rnd)
				}
				result = result.Add(info.Color.Mul(direct.Add(indirect)))
			}
		}
	}
	return result.DivScalar(float64(n * n))
}

func (s *DefaultSampler) directLight(scene *Scene, n Ray, rnd *rand.Rand) Color {
	nLights := len(scene.Lights)
	if nLights == 0 {
		return Color{}
	}

	// pick a random light
	light := scene.Lights[rand.Intn(nLights)]

	// get bounding sphere center and radius
	var center Vector
	var radius float64
	switch t := light.(type) {
	case *Sphere:
		radius = t.Radius
		center = t.Center
	default:
		// get bounding sphere from bounding box
		box := t.BoundingBox()
		radius = box.OuterRadius()
		center = box.Center()
	}

	// get random point on sphere surface
	// TODO: use disk instead, this is biased?
	point := RandomUnitVector(rnd).MulScalar(radius).Add(center)

	// construct ray toward light point
	ray := Ray{n.Origin, point.Sub(n.Origin).Normalize()}

	// get cosine term
	diffuse := ray.Direction.Dot(n.Direction)
	if diffuse <= 0 {
		return Color{}
	}

	// check for light visibility
	hit := scene.Intersect(ray)
	if !hit.Ok() || hit.Shape != light {
		return Color{}
	}

	// get material properties from light
	material := light.MaterialAt(point)
	color := light.ColorAt(point)
	emittance := material.Emittance

	// compute solid angle (hemisphere coverage)
	hyp := center.Sub(n.Origin).Length()
	opp := radius
	// TODO: fix issue where hyp < opp (point inside sphere)
	theta := math.Asin(opp / hyp)
	adj := opp / math.Tan(theta)
	d := math.Cos(theta) * adj
	r := math.Sin(theta) * adj
	coverage := (r * r) / (d * d)

	m := diffuse * emittance * coverage * float64(nLights)
	return color.MulScalar(m)
}
