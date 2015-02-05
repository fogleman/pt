package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	scene.AddLight(pt.NewSphere(pt.Vector{2, 8, 0.5}, 2, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0.376662, 30}, pt.HexColor(0xFCFAE1), pt.Material{1.5, 0, 0}))
	mesh := pt.NewMesh(pt.HexColor(0x5BCFDE), pt.Material{1.5, 0, 0})
	mesh.LoadOBJ("examples/gopher.obj")
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{8, 3, 0.5}, pt.Vector{-1, 2.5, 0.5}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 16, 64, 4)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
