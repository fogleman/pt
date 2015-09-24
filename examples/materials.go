package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	r := 0.4
	var material pt.Material

	material = pt.DiffuseMaterial(pt.HexColor(0x334D5C))
	scene.Add(pt.NewSphere(pt.Vector{-2, r, 0}, r, material))

	material = pt.SpecularMaterial(pt.HexColor(0x334D5C), 2)
	scene.Add(pt.NewSphere(pt.Vector{-1, r, 0}, r, material))

	material = pt.GlossyMaterial(pt.HexColor(0x334D5C), 2, pt.Radians(50))
	scene.Add(pt.NewSphere(pt.Vector{0, r, 0}, r, material))

	material = pt.TransparentMaterial(pt.HexColor(0x334D5C), 2, pt.Radians(20), 1)
	scene.Add(pt.NewSphere(pt.Vector{1, r, 0}, r, material))

	material = pt.ClearMaterial(2, 0)
	scene.Add(pt.NewSphere(pt.Vector{2, r, 0}, r, material))

	material = pt.SpecularMaterial(pt.HexColor(0xFFFFFF), 1000)
	scene.Add(pt.NewSphere(pt.Vector{0, 1.5, -4}, 1.5, material))

	scene.Add(pt.NewCube(pt.Vector{-1000, -1, -1000}, pt.Vector{1000, 0, 1000}, pt.GlossyMaterial(pt.HexColor(0xFFFFFF), 1.4, pt.Radians(20))))
	scene.Add(pt.NewSphere(pt.Vector{0, 5, 0}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 3, pt.LinearAttenuation(0.4))))
	camera := pt.LookAt(pt.Vector{0, 3, 6}, pt.Vector{0, 1, 0}, pt.Vector{0, 1, 0}, 30)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, -1, 16, 16)
}
