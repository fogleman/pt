package main

import (
	"math/rand"

	"github.com/fogleman/pt/pt"
)

func main() {
	scene := pt.Scene{}
	floor := pt.GlossyMaterial(pt.HexColor(0xFCFFF5), 1.2, pt.Radians(30))
	cube := pt.GlossyMaterial(pt.HexColor(0xFF8C00), 1.3, pt.Radians(20))
	ball := pt.GlossyMaterial(pt.HexColor(0xD90000), 1.4, pt.Radians(10))
	n := 7
	fn := float64(n)
	for z := 0; z < n; z++ {
		for x := 0; x < n-z; x++ {
			for y := 0; y < n-z-x; y++ {
				fx, fy, fz := float64(x), float64(y), float64(z)
				scene.Add(pt.NewCube(pt.Vector{fx, fy, fz}, pt.Vector{fx + 1, fy + 1, fz + 1}, cube))
				if x+y == n-z-1 {
					if rand.Float64() > 0.75 {
						scene.Add(pt.NewSphere(pt.Vector{fx + 0.5, fy + 0.5, fz + 1.5}, 0.35, ball))
					}
				}
			}
		}
	}
	scene.Add(pt.NewCube(pt.Vector{-1000, -1000, -1}, pt.Vector{1000, 1000, 0}, floor))
	scene.Add(pt.NewSphere(pt.Vector{fn, fn / 3, fn * 2}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 1, pt.NoAttenuation)))
	camera := pt.LookAt(pt.Vector{fn * 2, fn * 2, fn * 2}, pt.Vector{0, 0, fn / 4}, pt.Vector{0, 0, 1}, 35)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2880, 1800, -1, 16, 4)
}
