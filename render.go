package pt

import (
	"image"
	"image/color"
)

func Render(scene *Scene, camera *Camera, w, h, samples int) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := Color{}
			for n := 0; n < samples; n++ {
				c = c.Add(scene.Sample(camera.CastRay(x, y, w, h)))
			}
			c = c.Div(float64(samples))
			r, g, b := uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255)
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
	return image
}
