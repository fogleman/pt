package main

import (
	"log"
	"github.com/fogleman/pt/pt"
)

func main() {
	scene := pt.Scene{}
	scene.AddLight(pt.NewSphere(pt.Vector{-2, 5, -3}, 2, pt.DiffuseMaterial(pt.Color{1.5, 1.5, 1.5}), nil))
	scene.AddLight(pt.NewSphere(pt.Vector{5, 5, -3}, 0.5, pt.DiffuseMaterial(pt.Color{1.5, 1.5, 1.5}), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, pt.Material{pt.HexColor(0xFCFAE1), 2, 0, 0}))
	mesh, err := pt.LoadOBJ("examples/teapot.obj", pt.Material{pt.HexColor(0xB9121B), 2, 0, 0})
	if err != nil {
		log.Fatalf("LoadOBJ error: %v", err)
	}
	scene.AddShape(mesh)
	camera := pt.LookAt(pt.Vector{2, 5, -6}, pt.Vector{0.5, 1, 0}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 16, 16, 4)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
