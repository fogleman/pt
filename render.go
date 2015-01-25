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

func Render(scene *Scene, camera *Camera, w, h, samples, bounces int) image.Image {
	image := image.NewNRGBA(image.Rect(0, 0, w, h))
	n := runtime.GOMAXPROCS(0)
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		go func (i int) {
		    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := 0; y < h; y++ {
				if y % n != i {
					continue
				}
				if i == 0 {
					pct := 100 * float64(y) / float64(h - 1)
					fmt.Printf("\r%d / %d (%.1f%%)", y + 1, h, pct)
				}
				for x := 0; x < w; x++ {
					c := Color{}
					for n := 0; n < samples; n++ {
						c = c.Add(scene.RecursiveSample(camera.CastRay(x, y, w, h, rnd), bounces, rnd))
					}
					c = c.Div(float64(samples))
					r := uint8(math.Min(255, c.R * 255))
					g := uint8(math.Min(255, c.G * 255))
					b := uint8(math.Min(255, c.B * 255))
					image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
				}
			}
			c <- 1
		}(i)
	}
	for i := 0; i < n; i++ {
		<- c
	}
	fmt.Println()
	return image
}
