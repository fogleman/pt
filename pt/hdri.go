package pt

import (
	"fmt"
	"os"

	"github.com/Opioid/rgbe"
)

func NewHDRITexture(path string, scale float64) (Texture, error) {
	fmt.Printf("Loading HDRI: %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	width, height, data, err := rgbe.Decode(file)
	if err != nil {
		return nil, err
	}

	colors := make([]Color, width*height)
	var sum float64
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := width*y + x
			j := i * 3
			r := float64(data[j])
			g := float64(data[j+1])
			b := float64(data[j+2])
			sum += r + g + b
			colors[i] = Color{r, g, b}
		}
	}

	mean := sum / 3 / float64(width*height)
	s := scale / mean
	for i, c := range colors {
		colors[i] = c.MulScalar(s)
	}

	return &ColorTexture{width, height, colors}, nil
}
