package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	material := pt.DiffuseMaterial(pt.HexColor(0xFCFAE1))
	scene.Add(pt.NewCube(pt.Vector{-1000, -1, -1000}, pt.Vector{1000, 0, 1000}, material))
	for x := -20; x <= 20; x++ {
		for z := -20; z <= 20; z++ {
			if (x+z)%2 == 0 {
				continue
			}
			s := 0.1
			min := pt.Vector{float64(x) - s, 0, float64(z) - s}
			max := pt.Vector{float64(x) + s, 2, float64(z) + s}
			scene.Add(pt.NewCube(min, max, material))
		}
	}
	scene.Add(pt.NewCube(pt.Vector{-5, 10, -5}, pt.Vector{5, 11, 5}, pt.LightMaterial(pt.Color{1, 1, 1}, 5, pt.QuadraticAttenuation(0.05))))
	camera := pt.LookAt(pt.Vector{20, 10, 0}, pt.Vector{8, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, 0, 4, 2)
}
