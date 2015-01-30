package main

import "github.com/fogleman/pt"

func main() {
	white := pt.Color{0.740, 0.742, 0.734}
	red := pt.Color{0.366, 0.037, 0.042}
	green := pt.Color{0.163, 0.409, 0.083}
	light := pt.Color{0.780, 0.780, 0.776}
	scene := pt.Scene{}
	scene.AddShape(pt.NewPlane(pt.Vector{0, -10, 0}, pt.Vector{0, 1, 0}, white, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewPlane(pt.Vector{0, 10, 0}, pt.Vector{0, -1, 0}, white, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewPlane(pt.Vector{0, 0, 10}, pt.Vector{0, 0, -1}, white, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewPlane(pt.Vector{0, 0, -10}, pt.Vector{0, 0, 1}, white, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewPlane(pt.Vector{-10, 0, 0}, pt.Vector{1, 0, 0}, red, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewPlane(pt.Vector{10, 0, 0}, pt.Vector{-1, 0, 0}, green, pt.DiffuseMaterial()))
	scene.AddShape(pt.NewSphere(pt.Vector{3, -7, -3}, 3, white, pt.RefractiveMaterial(3), nil))
	cube := pt.NewCube(pt.Vector{-3, -4, -3}, pt.Vector{3, 4, 3}, light, pt.DiffuseMaterial())
	transform := pt.Rotate(pt.Vector{0, 1, 0}, pt.Radians(30)).Translate(pt.Vector{-3, -6, 4})
	scene.AddShape(pt.NewTransformedShape(cube, transform))
	scene.AddLight(pt.NewCube(pt.Vector{-2, 9.8, -2}, pt.Vector{2, 10, 2}, light, pt.DiffuseMaterial()))
	camera := pt.LookAt(pt.Vector{0, 0, -20}, pt.Vector{0, 0, 1}, pt.Vector{0, 1, 0}, 65)
	im := pt.Render(&scene, &camera, 512, 512, 4, 16, 3)
	if err := pt.Save("out.png", im); err != nil {
		panic(err)
	}
}
