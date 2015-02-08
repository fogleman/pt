package main

import "fmt"
import "github.com/fogleman/pt/pt"

func render(angle int) {
	scene := pt.Scene{}
	wall := pt.Material{pt.HexColor(0xFCFAE1), 2, 0, 0}
	scene.AddLight(pt.NewSphere(pt.Vector{2, 8, 0.5}, 2, pt.DiffuseMaterial(pt.Color{1, 1, 1}), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{-8, 10, 30}, wall))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, wall))
	material := pt.Material{pt.Color{}, 1.5, pt.Radians(30), 0}
	mesh, _ := pt.LoadOBJ("examples/gopher.obj", material)
	mesh.SmoothNormals()
	mesh.MoveTo(pt.Vector{}, pt.Vector{0.5, 0, 0.5})
	m := pt.Rotate(pt.Vector{0, 1, 0}, pt.Radians(float64(angle)))
	scene.AddShape(pt.NewTransformedShape(mesh, m))
	camera := pt.LookAt(pt.Vector{10, 3, 0}, pt.Vector{-1, 2, 0}, pt.Vector{0, 1, 0}, 45)
	path := fmt.Sprintf("out%03d.png", angle)
	pt.IterativeRender(path, 1, &scene, &camera, 256, 256, 4, 16, 4)
}

func main() {
	for i := 0; i < 360; i += 5 {
		render(i)
	}
}