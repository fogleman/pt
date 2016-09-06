package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	material := DiffuseMaterial(HexColor(0xFCFAE1))
	scene.Add(NewCube(V(-1000, -1, -1000), V(1000, 0, 1000), material))
	for x := -20; x <= 20; x++ {
		for z := -20; z <= 20; z++ {
			if (x+z)%2 == 0 {
				continue
			}
			s := 0.1
			min := V(float64(x)-s, 0, float64(z)-s)
			max := V(float64(x)+s, 2, float64(z)+s)
			scene.Add(NewCube(min, max, material))
		}
	}
	scene.Add(NewCube(V(-5, 10, -5), V(5, 11, 5), LightMaterial(White, 5)))
	camera := LookAt(V(20, 10, 0), V(8, 0, 0), V(0, 1, 0), 45)
	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}
