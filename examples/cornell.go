package main

import . "github.com/hborntraeger/pt/pt"

func main() {
	scene := Scene{}
	material := DiffuseMaterial(Black)
	mesh, err := LoadOBJ("examples/CornellBox-Original.obj", material)
	if err != nil {
		panic(err)
	}
	for _, t := range mesh.Triangles {
		scene.Add(t)
	}
	camera := LookAt(V(0, 1, 3), V(0, 1, 0), V(0, 1, 0), 50)
	sampler := NewSampler(4, 8)
	renderer := NewRenderer(&scene, &camera, sampler, 512, 512)
	renderer.IterativeRender("out%03d.png", 1000)
}
