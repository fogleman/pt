package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	wall := pt.Material{pt.HexColor(0xFCFAE1), 2, 0, 0}
	scene.AddLight(pt.NewSphere(pt.Vector{2, 8, 0.5}, 2, pt.DiffuseMaterial(pt.Color{1, 1, 1}), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{-8, 10, 30}, wall))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0.376662, 30}, wall))
	material := pt.Material{pt.Color{}, 1.5, pt.Radians(30), 0}
	mesh, _ := pt.LoadOBJ("examples/gopher.obj", material)
	mesh.SmoothNormals()
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{8, 3, 0.5}, pt.Vector{-1, 2.5, 0.5}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/4, 1440/4, 0, 16, 4)
}
