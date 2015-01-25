package pt

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func Render(scene *Scene, camera *Camera, w, h, cameraSamples, hitSamples, depth int) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	ncpu := runtime.GOMAXPROCS(0)
	ch := make(chan int, ncpu)
	for i := 0; i < ncpu; i++ {
		go func (i int) {
		    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := i; y < h; y += ncpu {
				if i == 0 {
					pct := 100 * float64(y) / float64(h - 1)
					fmt.Printf("\r%d / %d (%.1f%%)", y + 1, h, pct)
				}
				for x := 0; x < w; x++ {
					c := Color{}
					n := int(math.Sqrt(float64(cameraSamples)))
					for u := 0; u < n; u++ {
						for v := 0; v < n; v++ {
							fu := (float64(u) + 0.5) * (1 / float64(n))
							fv := (float64(v) + 0.5) * (1 / float64(n))
							ray := camera.CastRay(x, y, w, h, fu, fv)
							c = c.Add(scene.Sample(ray, hitSamples, depth, rnd))
						}
					}
					c = c.Div(float64(n * n))
					r := uint8(math.Min(255, c.R * 255))
					g := uint8(math.Min(255, c.G * 255))
					b := uint8(math.Min(255, c.B * 255))
					image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
				}
			}
			ch <- 1
		}(i)
	}
	for i := 0; i < ncpu; i++ {
		<- ch
	}
	fmt.Println()
	return image
}
