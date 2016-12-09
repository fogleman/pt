package main

import (
	"fmt"
	"sync"

	. "github.com/fogleman/pt/pt"
)

func render(l, m int, path string) {
	scene := Scene{}

	eye := V(2, 2, 2)
	center := V(0, 0, 0.125)
	up := V(0, 0, 1)

	light := LightMaterial(White, 150)
	scene.Add(NewSphere(V(0, 0, 5), 0.5, light))
	scene.Add(NewSphere(V(5, 0, 2), 0.5, light))
	scene.Add(NewSphere(V(0, 5, 2), 0.5, light))

	pm := GlossyMaterial(HexColor(0x105B63), 1.3, Radians(30))
	nm := GlossyMaterial(HexColor(0xBD4932), 1.3, Radians(30))
	sh := NewSphericalHarmonic(l, m, pm, nm)
	scene.Add(sh)

	axis := DiffuseMaterial(White)
	scene.Add(NewSDFShape(NewCapsuleSDF(Vector{}, V(1, 0, 0), 0.005), axis))
	scene.Add(NewSDFShape(NewCapsuleSDF(Vector{}, V(0, 1, 0), 0.005), axis))
	scene.Add(NewSDFShape(NewCapsuleSDF(Vector{}, V(0, 0, 1), 0.005), axis))

	camera := LookAt(eye, center, up, 35)
	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	renderer := NewRenderer(&scene, &camera, sampler, 1600, 1600)
	renderer.AdaptiveSamples = 32
	// renderer.FireflySamples = 1024
	var wg sync.WaitGroup
	renderer.FrameRender(path, 64, &wg)
	wg.Wait()
}

func main() {
	// render(4, 0, "out.png")
	// return
	const n = 4
	for l := 0; l <= n; l++ {
		for m := -n; m <= n; m++ {
			path := fmt.Sprintf("sh.%d.%d.png", l, m+10)
			if m < -l || m > l {
				// dc := gg.NewContext(800, 800)
				// dc.SetRGB(0, 0, 0)
				// dc.Clear()
				// dc.SavePNG(path)
			} else {
				render(l, m, path)
			}
		}
	}
}
