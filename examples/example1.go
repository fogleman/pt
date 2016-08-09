package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	scene.Add(NewSphere(V(1.5, 1.25, 0), 1.25, SpecularMaterial(HexColor(0x334D5C), 1.3)))
	scene.Add(NewSphere(V(-1, 1, 2), 1, SpecularMaterial(HexColor(0xEFC94C), 1.3)))
	scene.Add(NewSphere(V(-2.5, 0.75, 0), 0.75, SpecularMaterial(HexColor(0xEFC94C), 10000)))
	scene.Add(NewSphere(V(-0.75, 0.75, -1), 0.5, TransparentMaterial(Color{}, 1.333, 0, 0)))
	scene.Add(NewCube(V(-100, -1, -100), V(100, 0, 100), DiffuseMaterial(Color{1, 1, 1})))
	scene.Add(NewSphere(V(-1.5, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 100, LinearAttenuation(1))))
	// scene.Add(NewSphere(V(0, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 100, LinearAttenuation(1))))
	// scene.Add(NewSphere(V(1.5, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 100, LinearAttenuation(1))))
	camera := LookAt(V(0, 2, -5), V(0, 0, 3), V(0, 1, 0), 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
