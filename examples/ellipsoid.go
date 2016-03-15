package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	wall := pt.SpecularMaterial(pt.HexColor(0xFCFAE1), 2)
	scene.Add(pt.NewSphere(pt.Vector{10, 10, 10}, 2, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewCube(pt.Vector{-100, -100, -100}, pt.Vector{-2, 100, 100}, wall))
	scene.Add(pt.NewCube(pt.Vector{-100, -100, -100}, pt.Vector{100, -1, 100}, wall))
	material := pt.GlossyMaterial(pt.HexColor(0x167F39), 1.5, pt.Radians(30))
	sphere := pt.NewSphere(pt.Vector{}, 1, material)
	m := pt.Scale(pt.Vector{1, 1, 2})
	m = m.Rotate(pt.Vector{0, 1, 0}, pt.Radians(15))
	m = m.Translate(pt.Vector{0, 0, -2})
	shape := pt.NewTransformedShape(sphere, m)
	scene.Add(shape)
	scene.Add(pt.NewSphere(pt.Vector{0, 0, 2}, 1, material))
	camera := pt.LookAt(pt.Vector{8, 3, 0}, pt.Vector{}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/4, 1440/4, -1, 16, 4)
}
