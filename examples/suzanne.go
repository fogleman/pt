package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	material := DiffuseMaterial(HexColor(0x334D5C))
	scene.Add(NewSphere(Vector{0.5, 1, 3}, 1, LightMaterial(Color{1, 1, 1}, 10)))
	// scene.Add(NewSphere(Vector{1.5, 1, 3}, 1, LightMaterial(Color{1, 1, 1}, 10)))
	scene.Add(NewCube(Vector{-5, -5, -2}, Vector{5, 5, -1}, material))
	mesh, err := LoadOBJ("examples/suzanne.obj", SpecularMaterial(HexColor(0xEFC94C), 1.3))
	if err != nil {
		panic(err)
	}
	scene.Add(mesh)
	camera := LookAt(Vector{1, -0.45, 4}, Vector{1, -0.6, 0.4}, Vector{0, 1, 0}, 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
