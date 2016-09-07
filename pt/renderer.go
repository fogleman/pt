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

type Renderer struct {
	Scene              *Scene
	Camera             *Camera
	Sampler            Sampler
	Buffer             *Buffer
	SamplesPerPixel    int
	StratifiedSampling bool
	AdaptiveSamples    int
	AdaptiveThreshold  float64
	AdaptiveExponent   float64
	FireflySamples     int
	FireflyThreshold   float64
	NumCPU             int
	Verbose            bool
}

func NewRenderer(scene *Scene, camera *Camera, sampler Sampler, w, h int) *Renderer {
	r := Renderer{}
	r.Scene = scene
	r.Camera = camera
	r.Sampler = sampler
	r.Buffer = NewBuffer(w, h)
	r.SamplesPerPixel = 1
	r.StratifiedSampling = false
	r.AdaptiveSamples = 0
	r.AdaptiveThreshold = 1
	r.AdaptiveExponent = 1
	r.FireflySamples = 0
	r.FireflyThreshold = 1
	r.NumCPU = runtime.NumCPU()
	r.Verbose = true
	return &r
}

func (r *Renderer) run() {
	scene := r.Scene
	camera := r.Camera
	sampler := r.Sampler
	buf := r.Buffer
	w, h := buf.W, buf.H
	spp := r.SamplesPerPixel
	sppRoot := int(math.Sqrt(float64(r.SamplesPerPixel)))
	ncpu := r.NumCPU

	runtime.GOMAXPROCS(ncpu)
	scene.Compile()
	ch := make(chan int, h)
	r.printf("%d x %d pixels, %d spp, %d cores\n", w, h, spp, ncpu)
	start := time.Now()
	scene.rays = 0
	for i := 0; i < ncpu; i++ {
		go func(i int) {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for y := i; y < h; y += ncpu {
				for x := 0; x < w; x++ {
					if r.StratifiedSampling {
						// stratified subsampling
						for u := 0; u < sppRoot; u++ {
							for v := 0; v < sppRoot; v++ {
								fu := (float64(u) + 0.5) / float64(sppRoot)
								fv := (float64(v) + 0.5) / float64(sppRoot)
								ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
								sample := sampler.Sample(scene, ray, rnd)
								buf.AddSample(x, y, sample)
							}
						}
					} else {
						// random subsampling
						for i := 0; i < spp; i++ {
							fu := rnd.Float64()
							fv := rnd.Float64()
							ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
							sample := sampler.Sample(scene, ray, rnd)
							buf.AddSample(x, y, sample)
						}
					}
					// adaptive sampling
					if r.AdaptiveSamples > 0 {
						v := buf.StandardDeviation(x, y).MaxComponent()
						v = Clamp(v/r.AdaptiveThreshold, 0, 1)
						v = math.Pow(v, r.AdaptiveExponent)
						samples := int(v * float64(r.AdaptiveSamples))
						for i := 0; i < samples; i++ {
							fu := rnd.Float64()
							fv := rnd.Float64()
							ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
							sample := sampler.Sample(scene, ray, rnd)
							buf.AddSample(x, y, sample)
						}
					}
					// firefly reduction
					if r.FireflySamples > 0 {
						if buf.StandardDeviation(x, y).MaxComponent() > r.FireflyThreshold {
							for i := 0; i < r.FireflySamples; i++ {
								fu := rnd.Float64()
								fv := rnd.Float64()
								ray := camera.CastRay(x, y, w, h, fu, fv, rnd)
								sample := sampler.Sample(scene, ray, rnd)
								buf.AddSample(x, y, sample)
							}
						}
					}
				}
				ch <- 1
			}
		}(i)
	}
	r.showProgress(start, scene.RayCount(), 0, h)
	for i := 0; i < h; i++ {
		<-ch
		r.showProgress(start, scene.RayCount(), i+1, h)
	}
	r.printf("\n")
}

func (r *Renderer) printf(format string, a ...interface{}) {
	if !r.Verbose {
		return
	}
	fmt.Printf(format, a...)
}

func (r *Renderer) showProgress(start time.Time, rays uint64, i, h int) {
	if !r.Verbose {
		return
	}
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

func (r *Renderer) writeImage(path string, buf *Buffer, channel Channel, wg *sync.WaitGroup) {
	defer wg.Done()
	im := buf.Image(channel)
	if err := SavePNG(path, im); err != nil {
		panic(err)
	}
}

func (r *Renderer) Render() image.Image {
	r.run()
	return r.Buffer.Image(ColorChannel)
}

func (r *Renderer) IterativeRender(pathTemplate string, iterations int) image.Image {
	var wg sync.WaitGroup
	for i := 1; i <= iterations; i++ {
		r.printf("\n[Iteration %d of %d]\n", i, iterations)
		r.run()
		path := pathTemplate
		if strings.Contains(path, "%") {
			path = fmt.Sprintf(pathTemplate, i)
		}
		buf := r.Buffer.Copy()
		wg.Add(1)
		go r.writeImage(path, buf, ColorChannel, &wg)
		// wg.Add(1)
		// go r.writeImage("deviation.png", buf, StandardDeviationChannel, &wg)
	}
	wg.Wait()
	return r.Buffer.Image(ColorChannel)
}

func (r *Renderer) ChannelRender() <-chan image.Image {
	ch := make(chan image.Image)
	go func() {
		for i := 1; ; i++ {
			r.run()
			ch <- r.Buffer.Image(ColorChannel)
		}
	}()
	return ch
}

func (r *Renderer) FrameRender(path string, iterations int, wg *sync.WaitGroup) {
	for i := 1; i <= iterations; i++ {
		r.run()
	}
	buf := r.Buffer.Copy()
	wg.Add(1)
	go r.writeImage(path, buf, ColorChannel, wg)
}

func (r *Renderer) TimedRender(duration time.Duration) image.Image {
	start := time.Now()
	for {
		r.run()
		if time.Since(start) > duration {
			break
		}
	}
	return r.Buffer.Image(ColorChannel)
}
