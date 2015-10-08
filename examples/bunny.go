package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xF2EBC7), 1.5, pt.Radians(0))
	mesh, err := pt.LoadOBJ("examples/bunny.obj", material)
	if err != nil {
		panic(err)
	}
	mesh.SmoothNormalsThreshold(pt.Radians(20))
	mesh.FitInside(pt.Box{pt.Vector{-1, 0, -1}, pt.Vector{1, 2, 1}}, pt.Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	floor := pt.GlossyMaterial(pt.HexColor(0x33332D), 1.2, pt.Radians(20))
	scene.Add(pt.NewCube(pt.Vector{-10000, -10000, -10000}, pt.Vector{10000, 0, 10000}, floor))
	scene.Add(pt.NewSphere(pt.Vector{0, 5, 0}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{4, 5, 4}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-1, 2, 3}, pt.Vector{0, 0.75, 0}, pt.Vector{0, 1, 0}, 50)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560, 1440, -1, 4, 4)
}
