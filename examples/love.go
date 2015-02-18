package main

import (
	"github.com/fogleman/pt/pt"
	"log"
)

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xF2F2F2), 1.5, pt.Radians(20))
	scene.Add(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, material))
	heart := pt.GlossyMaterial(pt.HexColor(0xF60A20), 1.5, pt.Radians(20))
	mesh, err := pt.LoadBinarySTL("examples/love.stl", heart)
	if err != nil {
		log.Fatalln("LoadBinarySTL error:", err)
	}
	mesh.FitInside(pt.Box{pt.Vector{-0.5, 0, -0.5}, pt.Vector{0.5, 1, 0.5}}, pt.Vector{0.5, 0, 0.5})
	scene.Add(mesh)
	scene.Add(pt.NewSphere(pt.Vector{-2, 10, -2}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{0, 10, -2}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{2, 10, -2}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{0, 1.5, -2}, pt.Vector{0, 0.5, 0}, pt.Vector{0, 1, 0}, 35)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 1440, 900, -1, 16, 4)
}
