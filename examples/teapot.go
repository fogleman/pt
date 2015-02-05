package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	scene.AddLight(pt.NewSphere(pt.Vector{-2, 5, -3}, 2, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	// scene.AddLight(pt.NewSphere(pt.Vector{5, 5, -3}, 0.5, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, pt.HexColor(0x334D5C), pt.DiffuseMaterial()))
	mesh, _ := pt.LoadOBJ("examples/teapot.obj")
	mesh.SetColor(pt.HexColor(0xEFC94C))
	mesh.SetMaterial(pt.DiffuseMaterial())//pt.Material{1.5, pt.Radians(0), 0})
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{2, 5, -6}, pt.Vector{0.5, 1, 0}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/2, 1440/2, 1, 4, 4)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
