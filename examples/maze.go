package main

import (
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

func main() {
	scene := Scene{}
	floor := GlossyMaterial(HexColor(0x7E827A), 1.1, Radians(30))
	material := GlossyMaterial(HexColor(0xE3CDA4), 1.1, Radians(30))
	scene.Add(NewCube(Vector{-10000, -10000, -10000}, Vector{10000, 10000, 0}, floor))
	n := 24
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			if rand.Float64() > 0.8 {
				min := Vector{float64(x) - 0.5, float64(y) - 0.5, 0}
				max := Vector{float64(x) + 0.5, float64(y) + 0.5, 1}
				cube := NewCube(min, max, material)
				scene.Add(cube)
			}
		}
	}
	a := NoAttenuation // QuadraticAttenuation(0.25)
	scene.Add(NewSphere(Vector{0, 0, 2.25}, 0.25, LightMaterial(Color{1, 1, 1}, 1, a)))
	camera := LookAt(Vector{1, 0, 30}, Vector{0, 0, 0}, Vector{0, 0, 1}, 35)
	IterativeRender("out%03d.png", 1000, &scene, &camera, 2560, 1440, -1, 4, 4)
}
