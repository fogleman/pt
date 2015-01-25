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
	scene.AddShape(&pt.Sphere{pt.Vector{0, 0, 0}, 1, pt.HexColor(0x334D5C)})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, -2}, 1, pt.HexColor(0x45B29D)})
	scene.AddShape(&pt.Sphere{pt.Vector{-2, 0, 2}, 1, pt.HexColor(0xEFC94C)})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, -2}, 1, pt.HexColor(0xE27A3F)})
	scene.AddShape(&pt.Sphere{pt.Vector{2, 0, 2}, 1, pt.HexColor(0xDF5A49)})
	scene.AddShape(&pt.Cube{pt.Vector{-10, -2, -10}, pt.Vector{10, -1, 10}, pt.Color{1, 1, 1}})
	scene.AddLight(&pt.Sphere{pt.Vector{-1, 4, -1}, 0.25, pt.Color{0.8, 0.8, 0.8}})
	camera.LookAt(pt.Vector{0, 6, -8}, pt.Vector{0, 0, -1}, pt.Vector{0, 1, 0}, 45)
	return scene, camera
}

func scene2() (pt.Scene, pt.Camera) {
	scene, camera := pt.Scene{}, pt.Camera{}
	scene.AddShape(&pt.Cube{pt.Vector{-10, -1, -10}, pt.Vector{10, 0, 10}, pt.Color{1, 1, 1}})
	for x := -2; x <= 2; x++ {
		for z := -2; z <= 2; z++ {
			scene.AddShape(&pt.Sphere{pt.Vector{float64(x) * 2, 0.8, float64(z) * 2}, 0.8, pt.HexColor(0x334D5C)})
		}
	}
	scene.AddLight(&pt.Sphere{pt.Vector{-1, 4, -1}, 0.25, pt.Color{0.8, 0.8, 0.8}})
	camera.LookAt(pt.Vector{0, 6, -2}, pt.Vector{0, 0, -1}, pt.Vector{0, 1, 0}, 55)
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
	im := pt.Render(&scene, &camera, 2560, 1440, 16, 64, 4)
	save("out.png", im)
}
