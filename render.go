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
	ch := make(chan int, h)
	fmt.Printf("%d x %d pixels, %d x %d = %d samples, %d bounces, %d cores\n",
		w, h, cameraSamples, hitSamples, cameraSamples*hitSamples, depth, ncpu)
	start := time.Now()
	for i := 0; i < ncpu; i++ {
		go func(i int) {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := i; y < h; y += ncpu {
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
					r := uint8(math.Min(255, c.R*255))
					g := uint8(math.Min(255, c.G*255))
					b := uint8(math.Min(255, c.B*255))
					image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
				}
				ch <- 1
			}
		}(i)
	}
	for i := 0; i < h; i++ {
		<-ch
		pct := int(100 * float64(i) / float64(h-1))
		elapsed := time.Since(start)
		hr := int(elapsed.Hours())
		min := int(elapsed.Minutes()) % 60
		sec := int(elapsed.Seconds()) % 60
		fmt.Printf("\r%4d / %d (%3d%%) [", i+1, h, pct)
		for p := 0; p < 100; p += 2 {
			if pct >= p {
				fmt.Printf("=")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("] %d:%02d:%02d", hr, min, sec)
	}
	fmt.Println()
	return image
}
