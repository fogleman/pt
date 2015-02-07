package main

import (
	"log"

	"github.com/fogleman/pt/pt"
)

func main() {
	scene := pt.Scene{}
	gopher := pt.Material{pt.Color{}, 3, pt.Radians(20), 0}
	wall := pt.Material{pt.HexColor(0xFCFAE1), 3, pt.Radians(20), 0}
	floor := pt.Material{pt.HexColor(0xFCFAE1), 3, pt.Radians(20), 0}
	scene.AddLight(pt.NewSphere(pt.Vector{10, 10, 10}, 2, pt.DiffuseMaterial(pt.Color{1, 1, 1}), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{-5, 10, 30}, wall))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, floor))
	mesh, err := pt.LoadOBJ("examples/gopher.obj", gopher)
	if err != nil {
		log.Fatalf("LoadOBJ error: %v", err)
	}
	mesh.SmoothNormals()
	mesh.MoveTo(pt.Vector{}, pt.Vector{0.5, 0, 0.5})
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{10, 3, 0}, pt.Vector{0, 2.5, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/4, 1440/4, 0, 16, 4)
}
