package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	. "github.com/fogleman/pt/pt"
)

var materials = []Material{
	GlossyMaterial(HexColor(0x730046), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xBFBB11), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xFFC200), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xE88801), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xC93C00), 1.333, Radians(30)),
}

var delay = []float64{
	1 / 100.0,
	8 / 100.0,
	24 / 100.0,
	30 / 100.0,
	44 / 100.0,
	60 / 100.0,
	67 / 100.0,
	80 / 100.0,
}

func sphere(scene *Scene, direction, anchor Vector, radius float64, depth, height int, t float64) {
	if height <= 0 {
		return
	}
	nt := (rand.Float64()*2 - 1) * 0.01
	tt := t - delay[depth] + nt
	tt = math.Max(math.Min(tt*5, 1), 0)
	r := radius * easeOutElastic(tt)
	if t > 1 {
		r = radius * (1 - easeInQuint((t-1)*10))
	}
	center := anchor.Add(direction.MulScalar(r))
	material := materials[(height+6)%len(materials)]
	if tt > 0 && r > 0 {
		scene.Add(NewSphere(center, r, material))
	}
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
				if d == direction.MulScalar(-1) {
					continue
				}
				c2 := center.Add(d.MulScalar(r))
				sphere(scene, d, c2, r2, depth+1, height-1, t)
			}
		}
	}
}

func frame(i, n int) {
	rand.Seed(123)
	fmt.Println(i, n)
	t := float64(i) / float64(n)
	a := t * math.Pi
	x := math.Cos(a) * 5
	y := math.Sin(a) * 5
	scene := Scene{}
	scene.SetColor(HexColor(0xFFFFFF))
	sphere(&scene, Vector{}, Vector{}, 1, 0, 7, t)
	scene.Add(NewSphere(Vector{0, 0, 6}, 0.5, LightMaterial(Color{1, 1, 1}, 1, NoAttenuation)))
	camera := LookAt(Vector{x, y, 1}, Vector{0, 0, 0}, Vector{0, 0, 1}, 30)
	template := fmt.Sprintf("out%03d.png", i)
	IterativeRender(template, 1, &scene, &camera, 2560/4, 1440/4, 1, 1, 4)
}

func easeInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

func easeOutElastic(t float64) float64 {
	p := 0.2
	return math.Pow(2, -10*t)*math.Sin((t-p/4)*(2*math.Pi)/p) + 1
}

func easeInBack(t float64) float64 {
	s := 1.70158
	return 1 - t*t*((s+1)*t-s)
}

func easeInQuint(t float64) float64 {
	return t * t * t * t * t
}

func main() {
	n := 360
	e := 40
	for i := n; i < n+e; i += 2 {
		frame(i, n)
	}
	time.Sleep(5 * time.Second)
}
