package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}

	var material Material
	var mesh *Mesh

	material = DiffuseMaterial(White)
	mesh, _ = LoadOBJ("examples/veach_scene/backdrop.obj", material)
	scene.Add(mesh)

	material = MetallicMaterial(White, Radians(20), 0)
	mesh, _ = LoadOBJ("examples/veach_scene/bar0.obj", material)
	scene.Add(mesh)

	material = MetallicMaterial(White, Radians(15), 0)
	mesh, _ = LoadOBJ("examples/veach_scene/bar1.obj", material)
	scene.Add(mesh)

	material = MetallicMaterial(White, Radians(10), 0)
	mesh, _ = LoadOBJ("examples/veach_scene/bar2.obj", material)
	scene.Add(mesh)

	material = MetallicMaterial(White, Radians(5), 0)
	mesh, _ = LoadOBJ("examples/veach_scene/bar3.obj", material)
	scene.Add(mesh)

	material = MetallicMaterial(White, Radians(0), 0)
	mesh, _ = LoadOBJ("examples/veach_scene/bar4.obj", material)
	scene.Add(mesh)

	scene.Add(NewSphere(Vector{3.75, 4.281, 0}, 1.8/2, LightMaterial(White, 3)))
	scene.Add(NewSphere(Vector{1.25, 4.281, 0}, 0.6/2, LightMaterial(White, 9)))
	scene.Add(NewSphere(Vector{-1.25, 4.281, 0}, 0.2/2, LightMaterial(White, 27)))
	scene.Add(NewSphere(Vector{-3.75, 4.281, 0}, 0.066/2, LightMaterial(White, 81.803)))

	scene.Add(NewSphere(V(0, 10, 4), 1, LightMaterial(White, 50)))

	camera := LookAt(Vector{0, 5, 12}, Vector{0, 1, 0}, Vector{0, 1, 0}, 50)
	sampler := NewSampler(4, 8)
	sampler.SpecularMode = SpecularModeAll
	sampler.LightMode = LightModeAll
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
