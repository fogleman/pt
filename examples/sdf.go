package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}

	light := LightMaterial(Color{1, 1, 1}, 10)
	scene.Add(NewSphere(V(2, 2, 5), 1, light))

	// a := &SphereSDF{1}
	b := &CapsuleSDF{V(0, 0, 0), V(-1, -1, 1), 0.2}
	// c := &IntersectionSDF{a, b}
	c := &TransformSDF{b, Translate(V(2, 0, 0))}
	material := DiffuseMaterial(Color{1, 1, 1})
	scene.Add(NewSDFShape(c, material))

	camera := LookAt(V(3, 3, 3), V(0, 0, 0), V(0, 0, 1), 50)
	sampler := NewSampler(4, 8)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -4)
}
