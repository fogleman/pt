package main

import . "github.com/fogleman/pt/pt"

func main() {
	white := DiffuseMaterial(Color{0.740, 0.742, 0.734})
	red := DiffuseMaterial(Color{0.366, 0.037, 0.042})
	green := DiffuseMaterial(Color{0.163, 0.409, 0.083})
	light := LightMaterial(Color{0.780, 0.780, 0.776}, 20, NoAttenuation)
	scene := Scene{}
	n := 10.0
	scene.Add(NewCube(V(-n, -11, -n), V(n, -10, n), white))
	scene.Add(NewCube(V(-n, 10, -n), V(n, 11, n), white))
	scene.Add(NewCube(V(-n, -n, 10), V(n, n, 11), white))
	scene.Add(NewCube(V(-11, -n, -n), V(-10, n, n), red))
	scene.Add(NewCube(V(10, -n, -n), V(11, n, n), green))
	scene.Add(NewSphere(V(3, -7, -3), 3, white))
	cube := NewCube(V(-3, -4, -3), V(3, 4, 3), white)
	transform := Rotate(V(0, 1, 0), Radians(30)).Translate(V(-3, -6, 4))
	scene.Add(NewTransformedShape(cube, transform))
	scene.Add(NewSphere(V(0, 7.75, 0), 2, light))
	camera := LookAt(V(0, 0, -20), V(0, 0, 1), V(0, 1, 0), 65)
	sampler := DefaultSampler{16, 16}
	IterativeRender("out%03d.png", 1000, &scene, &camera, &sampler, 384, 384, -1)
}
