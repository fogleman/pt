package main

import (
	"fmt"
	"math"

	. "github.com/fogleman/pt/pt"
)

var materials = []Material{
	GlossyMaterial(HexColor(0x730046), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xBFBB11), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xFFC200), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xE88801), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xC93C00), 1.4, Radians(30)),
}

func sphere(scene *Scene, previous, anchor Vector, radius float64, depth int, t float64) {
	if depth <= 0 {
		return
	}
	if t <= 0 {
		return
	}
	r := radius * ease(math.Min(t, 1))
	center := anchor.Add(previous.MulScalar(r))
	material := materials[(depth+6)%len(materials)]
	scene.Add(NewSphere(center, r, material))
	r2 := radius / 2.5
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				n := 0
				if dx != 0 {
					n++
				}
				if dy != 0 {
					n++
				}
				if dz != 0 {
					n++
				}
				if n != 1 {
					continue
				}
				d := Vector{float64(dx), float64(dy), float64(dz)}
				if d == previous.MulScalar(-1) {
					continue
				}
				c2 := center.Add(d.MulScalar(r))
				sphere(scene, d, c2, r2, depth-1, t-0.2)
			}
		}
	}
}

func frame(i, n int) {
	fmt.Println(i, n)
	t := float64(i) / float64(n)
	a := t * math.Pi
	x := math.Cos(a) * 5
	y := math.Sin(a) * 5
	scene := Scene{}
	scene.SetColor(HexColor(0xFFFFFF))
	sphere(&scene, Vector{}, Vector{}, 1, 7, 2*t)
	scene.Add(NewSphere(Vector{0, 0, 6}, 0.5, LightMaterial(Color{1, 1, 1}, 1, NoAttenuation)))
	camera := LookAt(Vector{x, y, 1}, Vector{0, 0, 0}, Vector{0, 0, 1}, 30)
	template := fmt.Sprintf("out%%03d-%03d.png", i)
	IterativeRender(template, 1, &scene, &camera, 2560/2, 1440/2, 9, 9, 4)
}

func ease(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

func main() {
	n := 360
	for i := 0; i < n; i++ {
		frame(i, n)
	}
}
