package main

import (
	"github.com/fogleman/pt/pt"
	"math/rand"
)

func main() {
	n := 20.0
	material := pt.Material{3, pt.Radians(30), 0}
	scene := pt.Scene{}
	scene.AddLight(pt.NewSphere(pt.Vector{0, 10, 0}, 2, pt.Color{2, 2, 2}, pt.DiffuseMaterial(), nil))
	scene.AddLight(pt.NewSphere(pt.Vector{n, 10, n}, 2, pt.Color{2, 2, 2}, pt.DiffuseMaterial(), nil))
	scene.AddLight(pt.NewSphere(pt.Vector{-n, 10, -n}, 2, pt.Color{2, 2, 2}, pt.DiffuseMaterial(), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-n, -1, -n}, pt.Vector{n, 0, n}, pt.HexColor(0x334D5C), pt.DiffuseMaterial()))
	for _, point := range pt.PoissonDisc(-n, -n, n, n, 1, 32) {
		point = pt.Vector{point.X, 0, point.Y}
		r := rand.Float64()*0.6 + 0.2
		sphere := pt.NewSphere(point, r, pt.HexColor(0xEFC94C), material, nil)
		scene.AddShape(sphere)
	}
	camera := pt.LookAt(pt.Vector{0, 5, -n}, pt.Vector{0, 0, -n / 2}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/2, 1440/2, 16, 16, 4)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
