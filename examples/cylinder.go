package main

import (
	"log"

	. "github.com/fogleman/pt/pt"
)

func createMesh(material Material) Shape {
	mesh, err := LoadBinarySTL("examples/cylinder.stl", material)
	if err != nil {
		log.Fatalln("LoadBinarySTL error:", err)
	}
	mesh.FitInside(Box{Vector{-0.1, -0.1, 0}, Vector{1.1, 1.1, 100}}, Vector{0.5, 0.5, 0})
	mesh.SmoothNormalsThreshold(Radians(10))
	return mesh
}

func main() {
	scene := Scene{}
	meshes := []Shape{
		createMesh(GlossyMaterial(HexColor(0x730046), 1.6, Radians(45))),
		createMesh(GlossyMaterial(HexColor(0xBFBB11), 1.6, Radians(45))),
		createMesh(GlossyMaterial(HexColor(0xFFC200), 1.6, Radians(45))),
		createMesh(GlossyMaterial(HexColor(0xE88801), 1.6, Radians(45))),
		createMesh(GlossyMaterial(HexColor(0xC93C00), 1.6, Radians(45))),
	}
	for x := -6; x <= 3; x++ {
		mesh := meshes[(x+6)%len(meshes)]
		for y := -5; y <= 4; y++ {
			fx := float64(x) / 2
			fy := float64(y)
			fz := float64(x) / 2
			scene.Add(NewTransformedShape(mesh, Translate(Vector{fx, fy, fz})))
		}
	}
	scene.Add(NewSphere(Vector{1, 0, 10}, 3, LightMaterial(Color{1, 1, 1}, 20)))
	camera := LookAt(Vector{-5, 0, 5}, Vector{1, 0, 0}, Vector{0, 0, 1}, 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
