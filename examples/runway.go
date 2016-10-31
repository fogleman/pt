package main

import . "github.com/fogleman/pt/pt"

const radius = 2
const height = 3
const emission = 3

func main() {
	scene := Scene{}

	white := DiffuseMaterial(White)
	floor := NewCube(V(-250, -1500, -1), V(250, 6200, 0), white)
	scene.Add(floor)

	light := LightMaterial(Kelvin(2700), emission)
	for y := 0; y <= 6000; y += 40 {
		scene.Add(NewSphere(V(-100, float64(y), height), radius, light))
		scene.Add(NewSphere(V(0, float64(y), height), radius, light))
		scene.Add(NewSphere(V(100, float64(y), height), radius, light))
	}

	for y := -40; y >= -750; y -= 20 {
		scene.Add(NewSphere(V(-10, float64(y), height), radius, light))
		scene.Add(NewSphere(V(0, float64(y), height), radius, light))
		scene.Add(NewSphere(V(10, float64(y), height), radius, light))
	}

	green := LightMaterial(HexColor(0x0BDB46), emission)
	red := LightMaterial(HexColor(0xDC4522), emission)
	for x := -160; x <= 160; x += 10 {
		scene.Add(NewSphere(V(float64(x), -20, height), radius, green))
		scene.Add(NewSphere(V(float64(x), 6100, height), radius, red))
	}

	scene.Add(NewSphere(V(-160, 250, height), radius, red))
	scene.Add(NewSphere(V(-180, 250, height), radius, red))
	scene.Add(NewSphere(V(-200, 250, height), radius, light))
	scene.Add(NewSphere(V(-220, 250, height), radius, light))

	for i := 0; i < 5; i++ {
		y := float64((i + 1) * -120)
		for j := 1; j <= 4; j++ {
			x := float64((j + 4)) * 7.5
			scene.Add(NewSphere(V(x, y, height), radius, red))
			scene.Add(NewSphere(V(-x, y, height), radius, red))
			scene.Add(NewSphere(V(x, -y, height), radius, light))
			scene.Add(NewSphere(V(-x, -y, height), radius, light))
		}
	}

	camera := LookAt(V(0, -1500, 200), V(0, -100, 0), V(0, 0, 1), 20)
	camera.SetFocus(V(0, 20000, 0), 1)

	sampler := NewSampler(4, 4)
	renderer := NewRenderer(&scene, &camera, sampler, 1024, 1024)
	renderer.IterativeRender("out%03d.png", 1000)
}
