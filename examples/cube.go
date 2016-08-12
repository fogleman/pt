package main

import (
	"log"
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

func createMesh(material Material) Shape {
	mesh, err := LoadSTL("examples/cube.stl", material)
	if err != nil {
		log.Fatalln("LoadSTL error:", err)
	}
	mesh.FitInside(Box{Vector{0, 0, 0}, Vector{1, 1, 1}}, Vector{0.5, 0.5, 0.5})
	return mesh
}

func main() {
	scene := Scene{}
	meshes := []Shape{
		createMesh(GlossyMaterial(HexColor(0x3B596A), 1.5, Radians(20))),
		createMesh(GlossyMaterial(HexColor(0x427676), 1.5, Radians(20))),
		createMesh(GlossyMaterial(HexColor(0x3F9A82), 1.5, Radians(20))),
		createMesh(GlossyMaterial(HexColor(0xA1CD73), 1.5, Radians(20))),
		createMesh(GlossyMaterial(HexColor(0xECDB60), 1.5, Radians(20))),
	}
	for x := -8; x <= 8; x++ {
		for z := -12; z <= 12; z++ {
			fx := float64(x)
			fy := rand.Float64() * 2
			fz := float64(z)
			scene.Add(NewTransformedShape(meshes[rand.Intn(len(meshes))], Translate(Vector{fx, fy, fz})))
			scene.Add(NewTransformedShape(meshes[rand.Intn(len(meshes))], Translate(Vector{fx, fy - 1, fz})))
		}
	}
	scene.Add(NewSphere(Vector{8, 10, 0}, 3, LightMaterial(Color{1, 1, 1}, 30)))
	camera := LookAt(Vector{-10, 10, 0}, Vector{-2, 0, 0}, Vector{0, 1, 0}, 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
