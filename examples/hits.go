package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	material := DiffuseMaterial(Color{0.95, 0.95, 1})
	light := LightMaterial(Color{1, 1, 1}, 1300, LinearAttenuation(1))
	scene.Add(NewSphere(Vector{-0.75, -0.75, 5}, 0.25, light))
	scene.Add(NewCube(Vector{-1000, -1000, -1000}, Vector{1000, 1000, 0}, material))
	mesh, err := LoadSTL("examples/hits.stl", material)
	mesh.SmoothNormalsThreshold(Radians(10))
	mesh.FitInside(Box{V(-1, -1, 0), V(1, 1, 2)}, V(0.5, 0.5, 0))
	if err != nil {
		panic(err)
	}
	scene.Add(mesh)
	camera := LookAt(Vector{1.6, -3, 2}, Vector{-0.25, 0.5, 0.5}, Vector{0, 0, 1}, 50)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1500, 1500, -1)
}
