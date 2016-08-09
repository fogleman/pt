package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	wall := SpecularMaterial(HexColor(0xFCFAE1), 2)
	scene.Add(NewSphere(V(4, 7, 3), 2, LightMaterial(Color{1, 1, 1}, 10, NoAttenuation)))
	scene.Add(NewCube(V(-30, -1, -30), V(-8, 10, 30), wall))
	scene.Add(NewCube(V(-30, -1, -30), V(30, 0.376662, 30), wall))
	material := GlossyMaterial(Color{}, 1.5, Radians(30))
	mesh, err := LoadOBJ("examples/gopher.obj", material)
	if err != nil {
		panic(err)
	}
	mesh.SmoothNormals()
	scene.Add(mesh)
	camera := LookAt(V(8, 3, 0.5), V(-1, 2.5, 0.5), V(0, 1, 0), 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
