package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	scene.Add(pt.NewSphere(pt.Vector{1.5, 1, 0}, 1, pt.SpecularMaterial(pt.HexColor(0x334D5C), 2)))
	scene.Add(pt.NewSphere(pt.Vector{-1, 1, 2}, 1, pt.SpecularMaterial(pt.HexColor(0xEFC94C), 2)))
	scene.Add(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
	scene.Add(pt.NewSphere(pt.Vector{-1, 4, -1}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 3, pt.LinearAttenuation(1))))
	camera := pt.LookAt(pt.Vector{0, 2, -5}, pt.Vector{0, 0, 3}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 10, &scene, &camera, 2560/4, 1440/4, 4, 64, 8)
}
