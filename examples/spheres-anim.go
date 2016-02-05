package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	. "github.com/fogleman/pt/pt"
)

const (
	FPS            = 30
	Duration1      = 14
	Duration2      = 2
	Frames         = (Duration1 + Duration2) * FPS
	BounceDuration = 2
	FadeDuration   = 1
	RPS            = 0.125 / Duration1
)

var materials = []Material{
	GlossyMaterial(HexColor(0x730046), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xBFBB11), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xFFC200), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xE88801), 1.333, Radians(30)),
	GlossyMaterial(HexColor(0xC93C00), 1.333, Radians(30)),
}

var BounceStart = []float64{
	0.25,
	2.25,
	3.00,
	5.75,
	8.00,
	9.50,
	12,
}

var BounceDeviation = []float64{
	0.0,
	0.05,
	0.1,
	0.15,
	0.15,
	0.2,
	0.2,
}

func sphere(scene *Scene, direction, anchor Vector, radius float64, depth, height int, t float64) {
	if height <= 0 {
		return
	}
	tt := t - BounceStart[depth] + rand.NormFloat64()*BounceDeviation[depth]
	tt = math.Max(math.Min(tt/BounceDuration, 1), 0)
	r := radius * easeOutElastic(tt)
	if t > Duration1 {
		r = radius * (1 - easeInQuint((t-Duration1)/FadeDuration))
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

func frame(i int) {
	rand.Seed(123)
	fmt.Println(i)
	t := float64(i) / FPS
	a := t * 2 * math.Pi * RPS
	x := math.Cos(a) * 5
	y := math.Sin(a) * 5
	scene := Scene{}
	scene.SetColor(HexColor(0xFFFFFF))
	sphere(&scene, Vector{}, Vector{}, 1, 0, 7, t)
	scene.Add(NewSphere(Vector{0, 0, 6}, 0.5, LightMaterial(Color{1, 1, 1}, 1, NoAttenuation)))
	camera := LookAt(Vector{x, y, 1}, Vector{0, 0, 0}, Vector{0, 0, 1}, 30)
	template := fmt.Sprintf("out%03d.png", i)
	IterativeRender(template, 1, &scene, &camera, 1920, 1080, 16, 16, 4)
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
	for i := 0; i < Frames; i += 1 {
		frame(i)
	}
	time.Sleep(5 * time.Second)
}
