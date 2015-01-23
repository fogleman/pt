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
			image.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 255})
			ray := camera.CastRay(x, y, w, h)
			t := scene.Intersect(ray)
			if t < INF {
				image.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
		}
	}
	if err = png.Encode(file, image); err != nil {
		return
	}
}
