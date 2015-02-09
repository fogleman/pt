package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	material := pt.GlossyMaterial(pt.HexColor(0xEFC94C), 3, pt.Radians(30))
	whiteMat := pt.GlossyMaterial(pt.Color{1, 1, 1}, 3, pt.Radians(30))
	for x := 0; x < 10; x++ {
		for z := 0; z < 10; z++ {
			center := pt.Vector{float64(x) - 4.5, 0, float64(z) - 4.5}
			scene.AddShape(pt.NewSphere(center, 0.4, material))
		}
	}
	// scene.AddShape(pt.NewPlane(pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, whiteMat))
	scene.AddShape(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, whiteMat))
	scene.AddLight(pt.NewSphere(pt.Vector{-1, 3, -1}, 0.5, pt.DiffuseMaterial(pt.Color{1, 1, 1}.MulScalar(60))))
	camera := pt.LookAt(pt.Vector{0, 4, -8}, pt.Vector{0, 0, -2}, pt.Vector{0, 1, 0}, 45)
	im := pt.Render(&scene, &camera, 2560/4, 1440/4, 4, 16, 8)
	if err := pt.SavePNG("out.png", im); err != nil {
		panic(err)
	}
}
