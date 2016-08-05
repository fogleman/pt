package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	. "github.com/fogleman/pt/pt"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run volume.go DIRECTORY")
		return
	}
	dirname := args[0]
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	var images []image.Image
	for _, info := range infos {
		filename := path.Join(dirname, info.Name())
		im, err := LoadPNG(filename)
		if err != nil {
			panic(err)
		}
		images = append(images, im)
	}

	scene := Scene{}
	scene.SetColor(Color{1, 1, 1})

	windows := []VolumeWindow{
		// VolumeWindow{0.18, 0.20, GlossyMaterial(HexColor(0x334D5C), 1.3, Radians(20))},
		// VolumeWindow{0.33, 0.38, GlossyMaterial(HexColor(0xEFC94C), 1.3, Radians(20))},
		VolumeWindow{0.10, 0.12, DiffuseMaterial(Color{0.2, 0.2, 0.2})},
		VolumeWindow{0.20, 0.22, DiffuseMaterial(Color{0.3, 0.3, 0.3})},
		VolumeWindow{0.30, 0.32, DiffuseMaterial(Color{0.4, 0.4, 0.4})},
		VolumeWindow{0.40, 0.42, DiffuseMaterial(Color{0.5, 0.5, 0.5})},
		VolumeWindow{0.50, 0.52, DiffuseMaterial(Color{0.6, 0.6, 0.6})},
		VolumeWindow{0.60, 0.62, DiffuseMaterial(Color{0.7, 0.7, 0.7})},
	}
	scene.Add(NewVolume(images, 2, windows))

	// light := LightMaterial(Color{1, 1, 1}, 3, NoAttenuation)
	// scene.Add(NewCube(V(-2, -2, 3), V(2, 2, 3.1), light))

	camera := LookAt(V(2, -2, 0), V(0, -0.25, 0), V(0, 0, 1), 32)
	sampler := DefaultSampler{1, 4}
	IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 512, 512, -1)
}
