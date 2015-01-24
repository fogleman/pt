package pt

import (
	"image"
	"image/color"
	"math"
)

func Render(scene *Scene, camera *Camera, w, h, samples int) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := Color{}
			for n := 0; n < samples; n++ {
				c = c.Add(scene.RecursiveSample(camera.CastRay(x, y, w, h), 8))
			}
			c = c.Div(float64(samples))
			r := uint8(math.Min(255, c.R * 255))
			g := uint8(math.Min(255, c.G * 255))
			b := uint8(math.Min(255, c.B * 255))
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
	return image
}
