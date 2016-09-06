package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	scene.Add(NewSphere(Vector{-2, 5, -3}, 0.5, LightMaterial(White, 50)))
	scene.Add(NewSphere(Vector{5, 5, -3}, 0.5, LightMaterial(White, 50)))
	scene.Add(NewCube(Vector{-30, -1, -30}, Vector{30, 0, 30}, SpecularMaterial(HexColor(0xFCFAE1), 2)))
	mesh, err := LoadOBJ("examples/teapot.obj", SpecularMaterial(HexColor(0xB9121B), 2))
	if err != nil {
		panic(err)
	}
	scene.Add(mesh)
	camera := LookAt(Vector{2, 5, -6}, Vector{0.5, 1, 0}, Vector{0, 1, 0}, 45)
	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}
