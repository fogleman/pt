package main

import (
	"github.com/fogleman/pt"
)

func main() {
	scene := pt.Scene{}
	scene.AddShape(&pt.Sphere{pt.Vector{0, 0, 0}, 1, pt.Color{1, 0, 0}})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, 2}, 1, pt.Color{0, 1, 0}})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, 2}, 1, pt.Color{0, 0, 1}})
	scene.AddLight(&pt.Sphere{pt.Vector{0, 5, -2}, 0.1, pt.Color{1, 1, 1}})

	camera := pt.Camera{}
	camera.LookAt(pt.Vector{0, 0, -5}, pt.Vector{}, pt.Vector{0, 1, 0}, 45)

	pt.Render("out.png", 800, 600, &scene, &camera)
}
