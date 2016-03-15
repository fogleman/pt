package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	wall := pt.GlossyMaterial(pt.HexColor(0xFCFAE1), 1.333, pt.Radians(30))
	scene.Add(pt.NewSphere(pt.Vector{10, 10, 10}, 2, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	scene.Add(pt.NewCube(pt.Vector{-100, -100, -100}, pt.Vector{-6, 100, 100}, wall))
	scene.Add(pt.NewCube(pt.Vector{-100, -100, -100}, pt.Vector{100, -1, 100}, wall))
	material := pt.GlossyMaterial(pt.HexColor(0x167F39), 1.333, pt.Radians(30))
	sphere := pt.NewSphere(pt.Vector{}, 1, material)
	for i := 0; i < 180; i += 30 {
		m := pt.Identity()
		m = m.Scale(pt.Vector{0.3, 1, 5})
		m = m.Rotate(pt.Vector{0, 1, 0}, pt.Radians(float64(i)))
		shape := pt.NewTransformedShape(sphere, m)
		scene.Add(shape)
	}
	camera := pt.LookAt(pt.Vector{8, 8, 0}, pt.Vector{1, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, -1, 4, 4)
}
