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

	scene.Add(NewVolume(images, 0.18, 0.20, GlossyMaterial(HexColor(0x334D5C), 1.3, Radians(0))))
	scene.Add(NewVolume(images, 0.33, 0.38, GlossyMaterial(HexColor(0xEFC94C), 1.3, Radians(0))))

	// light := LightMaterial(Color{1, 1, 1}, 3, NoAttenuation)
	// scene.Add(NewCube(V(-2, -2, 3), V(2, 2, 3.1), light))

	camera := LookAt(V(2, -2, 0), V(0, -0.25, 0), V(0, 0, 1), 32)
	sampler := DefaultSampler{1, 4}
	IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 1024, 1024, -1)

}
