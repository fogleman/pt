package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	material := DiffuseMaterial(HexColor(0x334D5C))
	scene.Add(NewSphere(V(0.5, 1, 3), 1, LightMaterial(White, 4)))
	scene.Add(NewSphere(V(1.5, 1, 3), 1, LightMaterial(White, 4)))
	scene.Add(NewCube(V(-5, -5, -2), V(5, 5, -1), material))
	mesh, err := LoadOBJ("examples/suzanne.obj", SpecularMaterial(HexColor(0xEFC94C), 1.3))
	if err != nil {
		panic(err)
	}
	scene.Add(mesh)
	camera := LookAt(V(1, -0.45, 4), V(1, -0.6, 0.4), V(0, 1, 0), 40)
	sampler := NewSampler(16, 8)
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}
