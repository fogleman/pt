package main

import (
	"fmt"
	"sync"

	. "github.com/hborntraeger/pt/pt"
)

func render(l, m int) {
	scene := Scene{}

	eye := V(1, 1, 1)
	center := V(0, 0, 0)
	up := V(0, 0, 1)

	light := LightMaterial(White, 150)
	scene.Add(NewSphere(V(0, 0, 5), 0.5, light))
	scene.Add(NewSphere(V(5, 0, 2), 0.5, light))
	scene.Add(NewSphere(V(0, 5, 2), 0.5, light))

	pm := GlossyMaterial(HexColor(0x105B63), 1.3, Radians(30))
	nm := GlossyMaterial(HexColor(0xBD4932), 1.3, Radians(30))
	sh := NewSphericalHarmonic(l, m, pm, nm)
	scene.Add(sh)

	camera := LookAt(eye, center, up, 50)
	sampler := NewSampler(4, 4)
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeFirst
	renderer := NewRenderer(&scene, &camera, sampler, 1600/2, 1600/2)
	renderer.AdaptiveSamples = 32
	// renderer.IterativeRender("out%03d.png", 1000)
	var wg sync.WaitGroup
	renderer.FrameRender(fmt.Sprintf("sh.%d.%d.png", l, m), 10, &wg)
	wg.Wait()
}

func main() {
	for l := 0; l <= 4; l++ {
		for m := -l; m <= l; m++ {
			render(l, m)
		}
	}
}
