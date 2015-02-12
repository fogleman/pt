package main

import "github.com/fogleman/pt/pt"

// http://graphics.cs.williams.edu/data/meshes/dragon.zip

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xFCFAE1), 1.5, pt.Radians(20))
	mesh, _ := pt.LoadOBJ("examples/dragon.obj", material)
	mesh.FitInside(pt.Box{pt.Vector{-1, -1, -1}, pt.Vector{1, 1, 1}}, pt.Vector{0.5, 0.5, 0.5})
	scene.Add(mesh)
	scene.Add(pt.NewSphere(pt.Vector{0, 10, 0}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-3, 1, 0}, pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 512, 512, 0, 16, 4)
}
