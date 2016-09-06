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
		im, err := LoadImage(filename)
		if err != nil {
			continue
			// panic(err)
		}
		images = append(images, im)
	}

	scene := Scene{}
	scene.Color = White

	colors := []Color{
		// HexColor(0xFFF8E3),

		HexColor(0x004358),
		HexColor(0x1F8A70),
		HexColor(0xBEDB39),
		HexColor(0xFFE11A),
		HexColor(0xFD7400),

		// HexColor(0xFFE11A),
		// HexColor(0xBEDB39),
		// HexColor(0x1F8A70),
		// HexColor(0x004358),

		// White,
		// White,
		// White,
		// White,
		// White,
		// White,
		// White,
		// White,
		// White,
	}
	const (
		start = 0.2
		size  = 0.01
		step  = 0.1
	)
	var windows []VolumeWindow
	for i := 0; i < len(colors); i++ {
		lo := start + step*float64(i)
		hi := lo + size
		material := GlossyMaterial(colors[i], 1.3, Radians(0))
		w := VolumeWindow{lo, hi, material}
		windows = append(windows, w)
	}
	box := Box{Vector{-1, -1, -0.2}, Vector{1, 1, 1}}
	volume := NewVolume(box, images, 3.4/0.9765625, windows)
	scene.Add(volume)

	// wall := GlossyMaterial(White, 1.1, Radians(20))
	// scene.Add(NewCube(V(-10, 0.65, -10), V(10, 10, 10), wall))

	// light := LightMaterial(White, 20)
	// scene.Add(NewSphere(V(1, -5, -1), 1, light))

	fmt.Println(volume.W, volume.H, volume.D)

	camera := LookAt(V(0, -3, -3), V(0, 0, 0), V(0, 0, -1), 35)
	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 512, 512)
	renderer.IterativeRender("out%03d.png", 1000)
}
