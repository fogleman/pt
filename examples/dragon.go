package main

import "github.com/fogleman/pt/pt"

// http://graphics.cs.williams.edu/data/meshes/dragon.zip

func main() {
	scene := pt.Scene{}
	scene.SetColor(pt.HexColor(0xFEE7E0))
	// material := pt.GlossyMaterial(pt.HexColor(0x5C832F), 1.5, pt.Radians(20))
	material := pt.TransparentMaterial(pt.HexColor(0xFFFFFF), 2, pt.Radians(20), 0)
	mesh, err := pt.LoadOBJ("examples/dragon.obj", material)
	if err != nil {
		panic(err)
	}
	mesh.FitInside(pt.Box{pt.Vector{-1, 0, -1}, pt.Vector{1, 2, 1}}, pt.Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	floor := pt.GlossyMaterial(pt.HexColor(0xD8CAA8), 1.2, pt.Radians(20))
	scene.Add(pt.NewCube(pt.Vector{-1000, -1000, -1000}, pt.Vector{1000, 0, 1000}, floor))
	scene.Add(pt.NewSphere(pt.Vector{0, 10, 0}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{-3, 2, -1}, pt.Vector{0, 0.5, 0}, pt.Vector{0, 1, 0}, 35)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/4, 1440/4, -1, 4, 4)
}
