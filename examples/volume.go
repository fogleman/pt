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

	// images = append(images[138:], images[:138]...)
	// images = images[len(images)/2:]

	scene := Scene{}
	scene.SetColor(Color{1, 1, 1})

	colors := []Color{
		HexColor(0x004358),
		HexColor(0x1F8A70),
		HexColor(0xBEDB39),
		HexColor(0xFFE11A),
		HexColor(0xFD7400),
	}
	const (
		start = 0.1
		size  = 0.001
		step  = 0.1
	)
	var windows []VolumeWindow
	for i := 0; i < len(colors); i++ {
		lo := start + step*float64(i)
		hi := lo + size
		material := GlossyMaterial(colors[i], 1.2, Radians(20))
		w := VolumeWindow{lo, hi, material}
		windows = append(windows, w)
	}
	volume := NewVolume(images, 6.5/0.429689, windows)
	scene.Add(volume)

	fmt.Println(volume.W, volume.H, volume.D)

	camera := LookAt(V(0, 0, 3), V(0, 0, 0), V(0, -1, 0), 40)
	sampler := DefaultSampler{4, 4}
	IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 2048, 2048, -1)
}
