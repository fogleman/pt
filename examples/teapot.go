package main

import "fmt"
import "github.com/fogleman/pt"

func main() {
	scene := pt.Scene{}
	scene.AddLight(pt.NewSphere(pt.Vector{-2, 5, -3}, 2, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	// scene.AddLight(pt.NewSphere(pt.Vector{5, 5, -3}, 0.5, pt.Color{1, 1, 1}, pt.DiffuseMaterial(), nil))
	scene.AddShape(pt.NewCube(pt.Vector{-30, -1, -30}, pt.Vector{30, 0, 30}, pt.HexColor(0x334D5C), pt.DiffuseMaterial()))
	shapes, _ := pt.LoadOBJ("examples/teapot.obj")
	fmt.Println(len(shapes))
	for _, shape := range shapes {
		scene.AddShape(shape)
	}
	camera := pt.LookAt(pt.Vector{2, 5, -6}, pt.Vector{0.5, 1, 0}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 16, 16, 8)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
