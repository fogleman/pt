package pt

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"sync"
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

func render(scene *Scene, camera *Camera, sampler Sampler, samplesPerPixel int, buf *Buffer) {
	w, h := buf.W, buf.H
	ncpu := runtime.NumCPU()
	runtime.GOMAXPROCS(ncpu)
	scene.Compile()
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
					if samplesPerPixel <= 0 {
						// random subsampling
						for i := 0; i < absSamples; i++ {
							fu := rnd.Float64()
							fv := rnd.Float64()
							ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
							sample := sampler.Sample(scene, ray, rnd)
							buf.AddSample(x, y, sample)
						}
					} else {
						// stratified subsampling
						n := int(math.Sqrt(float64(samplesPerPixel)))
						for u := 0; u < n; u++ {
							for v := 0; v < n; v++ {
								fu := (float64(u) + 0.5) / float64(n)
								fv := (float64(v) + 0.5) / float64(n)
								ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
								sample := sampler.Sample(scene, ray, rnd)
								buf.AddSample(x, y, sample)
							}
						}
					}
					// adaptive sampling
					v := Clamp(buf.StandardDeviation(x, y).MaxComponent(), 0, 1)
					v = math.Pow(v, 2)
					extraSamples := int(32 * v)
					for i := 0; i < extraSamples; i++ {
						fu := rnd.Float64()
						fv := rnd.Float64()
						ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
						sample := sampler.Sample(scene, ray, rnd)
						buf.AddSample(x, y, sample)
					}
					// firefly reduction
					if buf.StandardDeviation(x, y).MaxComponent() > 1 {
						for i := 0; i < 256; i++ {
							fu := rnd.Float64()
							fv := rnd.Float64()
							ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
							sample := sampler.Sample(scene, ray, rnd)
							buf.AddSample(x, y, sample)
						}
					}
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
}

func writeImage(path string, buf *Buffer, channel Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	im := buf.Image(channel)
	if err := SavePNG(path, im); err != nil {
		panic(err)
	}
}

func Render(scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) image.Image {
	buf := NewBuffer(w, h)
	render(scene, camera, sampler, samplesPerPixel, buf)
	return buf.Image(ColorChannel)
}

func IterativeRender(pathTemplate string, iterations int, scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) image.Image {
	var wg sync.WaitGroup
	scene.Compile()
	buf := NewBuffer(w, h)
	for i := 1; i <= iterations; i++ {
		fmt.Printf("\n[Iteration %d of %d]\n", i, iterations)
		render(scene, camera, sampler, samplesPerPixel, buf)
		path := pathTemplate
		if strings.Contains(path, "%") {
			path = fmt.Sprintf(pathTemplate, i)
		}
		bufCopy := buf.Copy()
		wg.Add(1)
		go writeImage(path, bufCopy, ColorChannel, &wg)
		wg.Add(1)
		go writeImage("deviation.png", bufCopy, StandardDeviationChannel, &wg)
	}
	wg.Wait()
	return buf.Image(ColorChannel)
}

func ChannelRender(scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int) <-chan image.Image {
	ch := make(chan image.Image)
	go func() {
		scene.Compile()
		buf := NewBuffer(w, h)
		for i := 1; ; i++ {
			render(scene, camera, sampler, samplesPerPixel, buf)
			ch <- buf.Image(ColorChannel)
		}
	}()
	return ch
}

func FrameRender(path string, iterations int, scene *Scene, camera *Camera, sampler Sampler, w, h, samplesPerPixel int, wg *sync.WaitGroup) {
	buf := NewBuffer(w, h)
	for i := 1; i <= iterations; i++ {
		render(scene, camera, sampler, samplesPerPixel, buf)
	}
	wg.Add(1)
	go writeImage(path, buf.Copy(), ColorChannel, wg)
}
