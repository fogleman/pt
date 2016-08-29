package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	scene.Color = White
	scene.Texture = GetTexture("examples/courtyard_ccby/courtyard_8k.png")
	material := GlossyMaterial(White, 2, Radians(0))
	material.Texture = GetTexture("examples/checker.png")
	scene.Add(NewSphere(V(0, 0, 0), 1, material))
	scene.Add(NewSphere(V(-2.5, 0, 0), 1, material))
	scene.Add(NewSphere(V(2.5, 0, 0), 1, material))
	scene.Add(NewSphere(V(0, 0, -2.5), 1, material))
	scene.Add(NewSphere(V(0, 0, 2.5), 1, material))
	material = GlossyMaterial(HexColor(0xEFECCA), 1.1, Radians(45))
	scene.Add(NewCube(V(-100, -100, -100), V(100, -1, 100), material))
	camera := LookAt(V(2, 3, 4), V(0, 0, 0), V(0, 1, 0), 40)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
