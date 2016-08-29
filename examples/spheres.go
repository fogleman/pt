package main

import . "github.com/fogleman/pt/pt"

var materials = []Material{
	GlossyMaterial(HexColor(0x730046), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xBFBB11), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xFFC200), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xE88801), 1.4, Radians(30)),
	GlossyMaterial(HexColor(0xC93C00), 1.4, Radians(30)),
}

func sphere(scene *Scene, previous, center Vector, radius float64, depth int) {
	if depth <= 0 {
		return
	}
	material := materials[(depth+5)%len(materials)]
	scene.Add(NewSphere(center, radius, material))
	r2 := radius / 2.5
	offset := radius + r2
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
				c2 := center.Add(d.MulScalar(offset))
				sphere(scene, d, c2, r2, depth-1)
			}
		}
	}
}

func main() {
	scene := Scene{}
	scene.Color = HexColor(0xFFFFFF)
	sphere(&scene, Vector{}, Vector{}, 1, 8)
	scene.Add(NewSphere(Vector{0, 0, 6}, 0.5, LightMaterial(White, 1)))
	camera := LookAt(Vector{3, 1.75, 1}, Vector{0.75, 0.5, 0}, Vector{0, 0, 1}, 30)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
