package main

import . "github.com/fogleman/pt/pt"

// http://graphics.cs.williams.edu/data/meshes/dragon.zip

func main() {
	scene := Scene{}
	scene.Color = HexColor(0xFEE7E0)
	// material := GlossyMaterial(HexColor(0x5C832F), 1.5, Radians(20))
	material := TransparentMaterial(HexColor(0xFFFFFF), 2, Radians(20), 0)
	mesh, err := LoadOBJ("examples/dragon.obj", material)
	if err != nil {
		panic(err)
	}
	mesh.FitInside(Box{Vector{-1, 0, -1}, Vector{1, 2, 1}}, Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	floor := GlossyMaterial(HexColor(0xD8CAA8), 1.2, Radians(20))
	scene.Add(NewCube(Vector{-1000, -1000, -1000}, Vector{1000, 0, 1000}, floor))
	scene.Add(NewSphere(Vector{0, 10, 0}, 1, LightMaterial(Color{1, 1, 1}, 20)))
	camera := LookAt(Vector{-3, 2, -1}, Vector{0, 0.5, 0}, Vector{0, 1, 0}, 35)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/4, 1080/4, -1)
}
