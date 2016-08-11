package main

import (
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

func main() {
	scene := Scene{}
	floor := GlossyMaterial(HexColor(0x7E827A), 1.1, Radians(30))
	material := GlossyMaterial(HexColor(0xE3CDA4), 1.1, Radians(30))
	scene.Add(NewCube(V(-10000, -10000, -10000), V(10000, 10000, 0), floor))
	n := 24
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			if rand.Float64() > 0.8 {
				min := V(float64(x)-0.5, float64(y)-0.5, 0)
				max := V(float64(x)+0.5, float64(y)+0.5, 1)
				cube := NewCube(min, max, material)
				scene.Add(cube)
			}
		}
	}
	scene.Add(NewSphere(V(0, 0, 2.25), 0.25, LightMaterial(Color{1, 1, 1}, 500)))
	camera := LookAt(V(1, 0, 30), V(0, 0, 0), V(0, 0, 1), 35)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
