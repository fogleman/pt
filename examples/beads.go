package main

import (
	"math"
	"math/rand"

	"github.com/fogleman/pt/pt"
)

func main() {
	materials := []pt.Material{
		pt.GlossyMaterial(pt.HexColor(0xFFFFFF), 1e4, 0),
		pt.TransparentMaterial(pt.HexColor(0xFFFFFF), 1.33, 0, 0),
		pt.GlossyMaterial(pt.HexColor(0x167F39), 1.3, pt.Radians(20)),
		pt.GlossyMaterial(pt.HexColor(0x45BF55), 1.3, pt.Radians(20)),
		pt.GlossyMaterial(pt.HexColor(0x96ED89), 1.3, pt.Radians(20)),
	}
	rand.Seed(1211)
	eye := pt.Vector{4, 2, 8}
	center := pt.Vector{0, 0, 0}
	up := pt.Vector{0, 0, 1}
	scene := pt.Scene{}
	for a := 0; a < 80; a++ {
		material := materials[rand.Intn(len(materials))]
		n := 400
		xs := LowPassNoise(n, 0.25, 4)
		ys := LowPassNoise(n, 0.25, 4)
		zs := LowPassNoise(n, 0.25, 4)
		position := pt.Vector{}
		for i := 0; i < n; i++ {
			sphere := pt.NewSphere(position, 0.1, material)
			scene.Add(sphere)
			v := pt.Vector{xs[i], ys[i], zs[i]}.Normalize().MulScalar(0.1)
			position = position.Add(v)
		}
	}
	scene.Add(pt.NewSphere(pt.Vector{4, 4, 20}, 2, pt.LightMaterial(pt.HexColor(0xFFFFFF), 1, pt.NoAttenuation)))
	fovy := 40.0
	camera := pt.LookAt(eye, center, up, fovy)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560, 1440, -1, 4, 4)
}

func Normalize(values []float64, a, b float64) []float64 {
	result := make([]float64, len(values))
	lo := values[0]
	hi := values[0]
	for _, x := range values {
		lo = math.Min(lo, x)
		hi = math.Max(hi, x)
	}
	for i, x := range values {
		p := (x - lo) / (hi - lo)
		result[i] = a + p*(b-a)
	}
	return result
}

func LowPass(values []float64, alpha float64) []float64 {
	result := make([]float64, len(values))
	var y float64
	for i, x := range values {
		y -= alpha * (y - x)
		result[i] = y
	}
	return result
}

func LowPassNoise(n int, alpha float64, iterations int) []float64 {
	result := make([]float64, n)
	for i := range result {
		result[i] = rand.Float64()*2 - 1
	}
	for i := 0; i < iterations; i++ {
		result = LowPass(result, alpha)
	}
	result = Normalize(result, -1, 1)
	return result
}
