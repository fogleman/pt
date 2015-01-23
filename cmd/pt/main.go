package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"github.com/fogleman/pt"
)

func main() {
	sphere := pt.Sphere{pt.Vector{}, 1}
	camera := pt.Camera{}
	camera.LookAt(pt.Vector{0, 0, -5}, pt.Vector{}, pt.Vector{0, 1, 0}, 45)
	w, h := 512, 512

	f, err := os.OpenFile("out.png", os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			image.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 255})
			ray := camera.CastRay(x, y, w, h)
			t := sphere.Intersect(ray)
			if t < pt.INF {
				image.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
		}
	}
	if err = png.Encode(f, image); err != nil {
		return
	}
}
