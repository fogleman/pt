package pt

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func showProgress(start time.Time, rays uint64, i, h int) {
	pct := int(100 * float64(i) / float64(h))
	elapsed := time.Since(start)
	rps := float64(rays) / elapsed.Seconds()
	fmt.Printf("\r%4d / %d (%3d%%) [", i, h, pct)
	for p := 0; p < 100; p += 3 {
		if pct > p {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %s %s ", DurationString(elapsed), NumberString(rps))
}

func render(scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) []Color {
	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)
	scene.Compile()
	pixels := make([]Color, w*h)
	ch := make(chan int, h)
	absSamples := int(math.Abs(float64(samplesPerPixel)))
	fmt.Printf("%d x %d pixels, %d spp, %d cores\n", w, h, absSamples, ncpu)
	start := time.Now()
	scene.rays = 0
	for i := 0; i < ncpu; i++ {
		go func(i int) {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := i; y < h; y += ncpu {
				for x := 0; x < w; x++ {
					c := Color{}
					if samplesPerPixel <= 0 {
						// random subsampling
						for i := 0; i < absSamples; i++ {
							fu := rnd.Float64()
							fv := rnd.Float64()
							ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
							c = c.Add(sampler.Sample(scene, ray, rnd))
						}
						c = c.DivScalar(float64(absSamples))
					} else {
						// stratified subsampling
						n := int(math.Sqrt(float64(samplesPerPixel)))
						for u := 0; u < n; u++ {
							for v := 0; v < n; v++ {
								fu := (float64(u) + 0.5) / float64(n)
								fv := (float64(v) + 0.5) / float64(n)
								ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
								c = c.Add(sampler.Sample(scene, ray, rnd))
							}
						}
						c = c.DivScalar(float64(n * n))
					}
					pixels[y*w+x] = c
				}
				ch <- 1
			}
		}(i)
	}
	showProgress(start, scene.RayCount(), 0, h)
	for i := 0; i < h; i++ {
		<-ch
		showProgress(start, scene.RayCount(), i+1, h)
	}
	fmt.Println()
	return pixels
}

func Render(scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) image.Image {
	pixels := render(scene, camera, sampler, w, h, samplesPerPixel)
	return pixelsToImage(pixels, w, h, 1)
}

func pixelsToImage(pixels []Color, w, h int, scale float64) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			result.SetRGBA(x, y, pixels[y*w+x].MulScalar(scale).Pow(1/2.2).RGBA())
		}
	}
	return result
}

func onIteration(pathTemplate string, i, w, h int, pixels []Color, frame []Color) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			index := y*w + x
			pixels[index] = pixels[index].Add(frame[index])
		}
	}
	scale := 1 / float64(i)
	result := pixelsToImage(pixels, w, h, scale)
	path := pathTemplate
	if strings.Contains(path, "%") {
		path = fmt.Sprintf(pathTemplate, i)
	}
	if err := SavePNG(path, result); err != nil {
		panic(err)
	}
}

func IterativeRender(pathTemplate string, iterations int, scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) {
	scene.Compile()
	pixels := make([]Color, w*h)
	for i := 1; i <= iterations; i++ {
		fmt.Printf("\n[Iteration %d of %d]\n", i, iterations)
		frame := render(scene, camera, sampler, w, h, samplesPerPixel)
		go onIteration(pathTemplate, i, w, h, pixels, frame)
	}
}
