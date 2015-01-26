package main

import (
	"github.com/fogleman/pt"
	"image"
	"image/png"
	"os"
	"runtime"
)

func scene1() (pt.Scene, pt.Camera) {
	scene, camera := pt.Scene{}, pt.Camera{}
	scene.AddShape(&pt.Sphere{pt.Vector{0, 0, 0}, 1, pt.HexColor(0x334D5C), pt.Material{}})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, -2}, 1, pt.HexColor(0x45B29D), pt.Material{}})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, 2}, 1, pt.HexColor(0xEFC94C), pt.Material{}})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, -2}, 1, pt.HexColor(0xE27A3F), pt.Material{}})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, 2}, 1, pt.HexColor(0xDF5A49), pt.Material{}})
	scene.AddShape(&pt.Cube{pt.Vector{-10, -2, -10}, pt.Vector{10, -1, 10}, pt.Color{1, 1, 1}, pt.Material{}})
	scene.AddLight(&pt.Sphere{pt.Vector{-1, 4, -1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}})
	scene.AddLight(&pt.Sphere{pt.Vector{1, 4, -1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}})
	scene.AddLight(&pt.Sphere{pt.Vector{-1, 4, 1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}})
	scene.AddLight(&pt.Sphere{pt.Vector{1, 4, 1}, 0.25, pt.Color{1, 1, 1}, pt.Material{}})
	camera.LookAt(pt.Vector{0, 6, -8}, pt.Vector{0, 0, -1}, pt.Vector{0, 1, 0}, 45)
	return scene, camera
}

func save(path string, im image.Image) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, im)
	if err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	scene, camera := scene1()
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 16, 16, 4)
	save("out.png", im)
}
