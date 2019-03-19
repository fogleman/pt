package main

import (
	"fmt"
	"math"
	"math/rand"

	. "github.com/hborntraeger/pt/pt"
)

func frame(path string, t float64) {
	materials := []Material{
		GlossyMaterial(HexColor(0x167F39), 1.3, Radians(20)),
		GlossyMaterial(HexColor(0x45BF55), 1.3, Radians(20)),
		GlossyMaterial(HexColor(0x96ED89), 1.3, Radians(20)),
	}
	rand.Seed(1211)
	eye := V(4, 2, 8)
	center := V(0, 0, 0)
	up := V(0, 0, 1)
	scene := Scene{}
	for a := 0; a < 80; a++ {
		material := materials[rand.Intn(len(materials))]
		n := 400
		xs := LowPassNoise(n, 0.25, 4)
		ys := LowPassNoise(n, 0.25, 4)
		zs := LowPassNoise(n, 0.25, 4)
		position := Vector{}
		positions := make([]Vector, n)
		for i := 0; i < n; i++ {
			positions[i] = position
			v := V(xs[i], ys[i], zs[i]).Normalize().MulScalar(0.1)
			position = position.Add(v)
		}
		for i := 0; i < n-1; i++ {
			a := positions[i]
			b := positions[i+1]
			p := a.Add(b.Sub(a).MulScalar(t))
			sphere := NewSphere(p, 0.1, material)
			scene.Add(sphere)
		}
	}
	scene.Add(NewSphere(V(4, 4, 20), 2, LightMaterial(HexColor(0xFFFFFF), 30)))
	fovy := 40.0
	camera := LookAt(eye, center, up, fovy)
	sampler := NewSampler(4, 4)
	sampler.SpecularMode = SpecularModeFirst
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}

func main() {
	for i := 0; i < 30; i++ {
		t := float64(i) / 30
		path := fmt.Sprintf("out%03d.png", i)
		fmt.Println(path)
		frame(path, t)
	}
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
