package main

import . "github.com/hborntraeger/pt/pt"

func main() {
	scene := Scene{}
	material := GlossyMaterial(HexColor(0xF2F2F2), 1.5, Radians(20))
	scene.Add(NewCube(Vector{-100, -1, -100}, Vector{100, 0, 100}, material))
	heart := GlossyMaterial(HexColor(0xF60A20), 1.5, Radians(20))
	mesh, err := LoadSTL("examples/love.stl", heart)
	if err != nil {
		panic(err)
	}
	mesh.FitInside(Box{Vector{-0.5, 0, -0.5}, Vector{0.5, 1, 0.5}}, Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	scene.Add(NewSphere(Vector{-2, 10, 2}, 1, LightMaterial(White, 30)))
	scene.Add(NewSphere(Vector{0, 10, 2}, 1, LightMaterial(White, 30)))
	scene.Add(NewSphere(Vector{2, 10, 2}, 1, LightMaterial(White, 30)))
	camera := LookAt(Vector{0, 1.5, 2}, Vector{0, 0.5, 0}, Vector{0, 1, 0}, 35)
	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}
