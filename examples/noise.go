package main

import (
	. "github.com/fogleman/pt/pt"
	"github.com/ojrac/opensimplex-go"
)

func main() {
	scene := Scene{}
	material := GlossyMaterial(Color{1, 1, 1}, 1.2, Radians(20))
	noise := opensimplex.New()
	n := 80
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n*2; k++ {
				x := float64(i - n/2)
				y := float64(j - n/2)
				z := float64(k)
				m := 0.15
				w := noise.Eval3(x*m, y*m, z*m)
				w = (w + 0.8) / 1.6
				if w <= 0.2 {
					shape := NewSphere(V(x, y, z), 0.333, material)
					scene.Add(shape)
				}
			}
		}
	}
	light := NewSphere(V(100, 0, 50), 5, LightMaterial(Color{1, 1, 1}, 300))
	scene.Add(light)
	camera := LookAt(V(0, 0, -20), V(0, 0, 0), V(0, 1, 0), 30)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
