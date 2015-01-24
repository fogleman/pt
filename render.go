package pt

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Render(path string, w, h int, scene *Scene, camera *Camera) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := scene.Sample(camera.CastRay(x, y, w, h))
			r, g, b := uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255)
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
	if err = png.Encode(file, image); err != nil {
		return
	}
}
