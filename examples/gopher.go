package main

import (
	"log"

	"github.com/fogleman/pt/pt"
)

func main() {
	scene := pt.Scene{}
	wall := pt.SpecularMaterial(pt.HexColor(0xFCFAE1), 2)
	scene.Add(pt.NewSphere(pt.Vector{4, 7, 3}, 2, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{-8, 10, 30}, wall))
	scene.Add(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0.376662, 30}, wall))
	material := pt.GlossyMaterial(pt.Color{}, 1.5, pt.Radians(30))
	mesh, err := pt.LoadOBJ("examples/gopher.obj", material)
	if err != nil {
		log.Fatalln("LoadOBJ error:", err)
	}
	mesh.SmoothNormals()
	scene.Add(mesh)
	camera := pt.LookAt(pt.Vector{8, 3, 0.5}, pt.Vector{-1, 2.5, 0.5}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 10, &scene, &camera, 2560/4, 1440/4, -1, 16, 4)
}
