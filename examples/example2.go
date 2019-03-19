package main

import . "github.com/hborntraeger/pt/pt"

func main() {
	scene := Scene{}
	material := GlossyMaterial(HexColor(0xEFC94C), 3, Radians(30))
	whiteMat := GlossyMaterial(White, 3, Radians(30))
	for x := 0; x < 40; x++ {
		for z := 0; z < 40; z++ {
			center := V(float64(x)-19.5, 0, float64(z)-19.5)
			scene.Add(NewSphere(center, 0.4, material))
		}
	}
	scene.Add(NewCube(V(-100, -1, -100), V(100, 0, 100), whiteMat))
	scene.Add(NewSphere(V(-1, 4, -1), 1, LightMaterial(White, 30)))
	camera := LookAt(V(0, 4, -8), V(0, 0, -2), V(0, 1, 0), 45)
	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 960, 540)
	renderer.IterativeRender("out%03d.png", 1000)
}
