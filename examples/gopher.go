package main

import . "github.com/hborntraeger/pt/pt"

func main() {
	scene := Scene{}

	// create materials
	gopher := GlossyMaterial(Black, 1.2, Radians(30))
	wall := GlossyMaterial(HexColor(0xFCFAE1), 1.5, Radians(10))
	light := LightMaterial(White, 80)

	// add walls and lights
	scene.Add(NewCube(V(-10, -1, -10), V(-2, 10, 10), wall))
	scene.Add(NewCube(V(-10, -1, -10), V(10, 0, 10), wall))
	scene.Add(NewSphere(V(4, 10, 1), 1, light))

	// load and transform gopher mesh
	mesh, err := LoadOBJ("examples/gopher.obj", gopher)
	if err != nil {
		panic(err)
	}
	mesh.Transform(Rotate(V(0, 1, 0), Radians(-10)))
	mesh.SmoothNormals()
	mesh.FitInside(Box{V(-1, 0, -1), V(1, 2, 1)}, V(0.5, 0, 0.5))
	scene.Add(mesh)

	// position camera
	camera := LookAt(V(4, 1, 0), V(0, 0.9, 0), V(0, 1, 0), 40)

	// render the scene
	sampler := NewSampler(16, 16)
	renderer := NewRenderer(&scene, &camera, sampler, 1024, 1024)
	renderer.IterativeRender("out%03d.png", 1000)
}
