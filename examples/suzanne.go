package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	material := pt.DiffuseMaterial()
	scene.AddLight(pt.NewSphere(pt.Vector{0.5, 1, 3}, 0.5, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	scene.AddLight(pt.NewSphere(pt.Vector{1.5, 1, 3}, 0.5, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-5, -5, -2}, pt.Vector{5, 5, -1}, pt.HexColor(0x334D5C), material))
	mesh := pt.NewMesh(pt.HexColor(0xEFC94C), pt.Material{2, 0, 0})
	mesh.LoadOBJ("examples/suzanne.obj")
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{1, -0.45, 4}, pt.Vector{1, -0.6, 0.4}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 1, 4, 4)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
