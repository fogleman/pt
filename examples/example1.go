package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	scene.Add(pt.NewSphere(pt.Vector{1.5, 1.25, 0}, 1.25, pt.SpecularMaterial(pt.HexColor(0x334D5C), 1.3)))
	scene.Add(pt.NewSphere(pt.Vector{-1, 1, 2}, 1, pt.SpecularMaterial(pt.HexColor(0xEFC94C), 1.3)))
	scene.Add(pt.NewSphere(pt.Vector{-2.5, 0.75, 0}, 0.75, pt.SpecularMaterial(pt.HexColor(0xEFC94C), 10000)))
	scene.Add(pt.NewSphere(pt.Vector{-0.75, 0.5, -1}, 0.5, pt.TransparentMaterial(pt.Color{}, 1.333, 0, 0)))
	scene.Add(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
	scene.Add(pt.NewSphere(pt.Vector{-1.5, 4, 0}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 100, pt.LinearAttenuation(1))))
	// scene.Add(pt.NewSphere(pt.Vector{0, 4, 0}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 100, pt.LinearAttenuation(1))))
	// scene.Add(pt.NewSphere(pt.Vector{1.5, 4, 0}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 100, pt.LinearAttenuation(1))))
	camera := pt.LookAt(pt.Vector{0, 2, -5}, pt.Vector{0, 0, 3}, pt.Vector{0, 1, 0}, 45)
	sampler := pt.DefaultSampler{16, 8}
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 2560, 1440, -1)
}
