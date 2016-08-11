package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	scene.Add(NewSphere(V(1.5, 1.25, 0), 1.25, SpecularMaterial(HexColor(0x004358), 1.2)))
	scene.Add(NewSphere(V(-1, 1, 2), 1, SpecularMaterial(HexColor(0xFFE11A), 1.2)))
	scene.Add(NewSphere(V(-2.5, 0.75, 0), 0.75, SpecularMaterial(HexColor(0xFD7400), 1.2)))
	scene.Add(NewSphere(V(-0.75, 0.5, -1), 0.5, SpecularMaterial(HexColor(0xBEDB39), 1.2)))
	scene.Add(NewCube(V(-100, -1, -100), V(100, 0, 100), GlossyMaterial(Color{1, 1, 1}, 1.1, Radians(10))))
	scene.Add(NewSphere(V(-1.5, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 30)))
	// scene.Add(NewSphere(V(0, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 50)))
	// scene.Add(NewSphere(V(1.5, 4, 0), 0.5, LightMaterial(Color{1, 1, 1}, 50)))
	camera := LookAt(V(0, 2.5, -5), V(0, 0, 3), V(0, 1, 0), 45)
	camera.SetFocus(V(-0.75, 1, -1), 0.15)
	sampler := NewSampler(4, 8)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
