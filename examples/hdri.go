package main

import . "github.com/fogleman/pt/pt"

func main() {
	scene := Scene{}
	texture, err := LoadTexture("/Users/fogleman/Desktop/out.png")
	if err != nil {
		panic(err)
	}
	scene.SetTexture(texture)
	texture, err = LoadTexture("/Users/fogleman/Desktop/checker.png")
	if err != nil {
		panic(err)
	}
	material := GlossyMaterial(Color{1, 1, 1}, 2, Radians(0))
	material.Texture = texture
	scene.Add(NewSphere(Vector{0, 0, 0}, 1, material))
	scene.Add(NewSphere(Vector{-2.5, 0, 0}, 1, material))
	scene.Add(NewSphere(Vector{2.5, 0, 0}, 1, material))
	scene.Add(NewSphere(Vector{0, 0, -2.5}, 1, material))
	scene.Add(NewSphere(Vector{0, 0, 2.5}, 1, material))
	material = GlossyMaterial(HexColor(0xEFECCA), 1.1, Radians(45))
	scene.Add(NewCube(Vector{-100, -100, -100}, Vector{100, -1, 100}, material))
	camera := LookAt(Vector{2, 3, 4}, Vector{0, 0, 0}, Vector{0, 1, 0}, 40)
	IterativeRender("out%03d.png", 10000, &scene, &camera, 1920, 1080, -1, 16, 4)
}
