package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	scene.AddShape(pt.NewSphere(pt.Vector{1.5, 1, 0}, 1, pt.SpecularMaterial(pt.HexColor(0x334D5C), 2)))
	scene.AddShape(pt.NewSphere(pt.Vector{-1, 1, 2}, 1, pt.SpecularMaterial(pt.HexColor(0xEFC94C), 2)))
	scene.AddShape(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 3, -1}, 0.5, pt.DiffuseMaterial(pt.Color{1, 1, 1}.MulScalar(60))))
	camera := pt.LookAt(pt.Vector{0, 2, -5}, pt.Vector{0, 0, 3}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/4, 1440/4, 0, 64, 8)
}
