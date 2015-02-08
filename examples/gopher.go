package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	wall := pt.Material{pt.HexColor(0xFCFAE1), nil, 2, 0, 0}
	scene.AddLight(pt.NewSphere(pt.Vector{2, 30, 0.5}, 2, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{-8, 30, 30}, wall))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, wall))
	material := pt.Material{pt.Color{}, nil, 1.5, pt.Radians(30), 0}
	mesh, _ := pt.LoadOBJ("examples/gopher.obj", material)
	mesh.SmoothNormals()
	mesh.MoveTo(pt.Vector{}, pt.Vector{0.5, 0, 0.5})
	for i := -2; i <= 2; i++ {
		m := pt.Translate(pt.Vector{0, 0, float64(i) * 5})
		scene.AddShape(pt.NewTransformedShape(mesh, m))
	}
	camera := pt.LookAt(pt.Vector{20, 4, 0}, pt.Vector{-1, 2, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out.png", 1000, &scene, &camera, 2560/4, 1440/4, 0, 16, 4)
}
