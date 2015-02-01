package main

import "github.com/fogleman/pt"

func main() {
	scene := pt.Scene{}
	material := pt.Material{3, pt.Radians(30), 0}
	for x := 0; x < 10; x++ {
		for z := 0; z < 10; z++ {
			center := pt.Vector{float64(x) - 4.5, 0, float64(z) - 4.5}
			scene.AddShape(pt.NewSphere(center, 0.4, pt.HexColor(0xEFC94C), material, nil))
		}
	}
	scene.AddShape(pt.NewPlane(pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, pt.Color{1, 1, 1}, material))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 3, -1}, 0.5, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	camera := pt.LookAt(pt.Vector{0, 4, -8}, pt.Vector{0, 0, -2}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 1, 4, 8)
	if err := pt.Save("out.png", im); err != nil {
		panic(err)
	}
}
