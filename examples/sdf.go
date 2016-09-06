package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}

	light := LightMaterial(White, 180)

	d := 4.0
	scene.Add(NewSphere(V(-1, -1, 0.5).Normalize().MulScalar(d), 0.25, light))
	scene.Add(NewSphere(V(0, -1, 0.25).Normalize().MulScalar(d), 0.25, light))
	scene.Add(NewSphere(V(-1, 1, 0).Normalize().MulScalar(d), 0.25, light))

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
	sampler.LightMode = LightModeAll
	sampler.SpecularMode = SpecularModeAll
	renderer := NewRenderer(&scene, &camera, sampler, 1600, 1600)
	renderer.IterativeRender("out%03d.png", 1000)
}
