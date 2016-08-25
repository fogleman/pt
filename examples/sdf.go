package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}

	light := LightMaterial(Color{1, 1, 1}, 100)

	d := 3.0
	scene.Add(NewSphere(V(-1, -1, 0.5).Normalize().MulScalar(d), 0.33, light))
	scene.Add(NewSphere(V(0, -1, -0.25).Normalize().MulScalar(d), 0.33, light))
	scene.Add(NewSphere(V(-1, 1, -0.25).Normalize().MulScalar(d), 0.33, light))

	material := GlossyMaterial(HexColor(0x468966), 1.2, Radians(20))
	sphere := NewSphereSDF(0.65)
	cube := NewCubeSDF(V(1, 1, 1))
	roundedCube := NewIntersectionSDF(sphere, cube)
	a := NewCylinderSDF(0.25, 1.1)
	b := NewTransformSDF(a, Rotate(V(1, 0, 0), Radians(90)))
	c := NewTransformSDF(a, Rotate(V(0, 0, 1), Radians(90)))
	difference := NewDifferenceSDF(roundedCube, a, b, c)
	sdf := NewTransformSDF(difference, Rotate(V(0, 0, 1), Radians(30)))
	scene.Add(NewSDFShape(sdf, material))

	floor := GlossyMaterial(HexColor(0xFFF0A5), 1.2, Radians(20))
	scene.Add(NewPlane(V(0, 0, -0.5), V(0, 0, 1), floor))

	camera := LookAt(V(-3, 0, 1), V(0, 0, 0), V(0, 0, 1), 35)
	sampler := NewSampler(4, 4)
	sampler.SpecularMode = SpecularModeFirst
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1024, 1024, -1)
}
