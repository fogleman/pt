package main

import (
	"github.com/fogleman/pt"
	"image/png"
	"os"
)

func main() {
	// place objects
	scene := pt.Scene{}
	scene.AddShape(&pt.Sphere{pt.Vector{0, 0, 0}, 1, pt.Color{1, 0.2, 0.2}})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, 2}, 1, pt.Color{0.2, 1, 0.2}})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, 2}, 1, pt.Color{0.2, 0.2, 1}})
	scene.AddShape(&pt.Cube{pt.Vector{-10, -2, -10}, pt.Vector{10, -1, 10}, pt.Color{1, 1, 1}})

	// place lights
	scene.AddLight(&pt.Sphere{pt.Vector{-4, 4, -4}, 0.1, pt.Color{0.8, 0.8, 0.8}})
	// scene.AddLight(&pt.Sphere{pt.Vector{5, 5, 0}, 0.1, pt.Color{1, 1, 1}})

	// place camera
	camera := pt.Camera{}
	camera.LookAt(pt.Vector{0, 0, -5}, pt.Vector{}, pt.Vector{0, 1, 0}, 45)

	// render image
	image := pt.Render(&scene, &camera, 800, 600, 256, 4)

	// save as png
	file, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, image)
	if err != nil {
		panic(err)
	}
}
