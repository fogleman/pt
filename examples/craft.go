package main

import (
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

const N = 16

const A = 1.0 / 2048
const B = 1.0/N - A

var Dirt = []int{0, 0, 0, 0, 0, 0}
var Grass = []int{16, 32, 16, 0, 16, 16}

var VP []Vector = []Vector{
	{-0.5, -0.5, 0.5},
	{0.5, -0.5, 0.5},
	{-0.5, 0.5, 0.5},
	{0.5, 0.5, 0.5},
	{-0.5, 0.5, -0.5},
	{0.5, 0.5, -0.5},
	{-0.5, -0.5, -0.5},
	{0.5, -0.5, -0.5},
}

var VT []Vector = []Vector{
	{A, A, 0},
	{B, A, 0},
	{A, B, 0},
	{B, B, 0},
}

var Triangles = [][3][2]Vector{
	{{VP[0], VT[0]}, {VP[1], VT[1]}, {VP[2], VT[2]}},
	{{VP[2], VT[2]}, {VP[1], VT[1]}, {VP[3], VT[3]}},
	{{VP[2], VT[0]}, {VP[3], VT[1]}, {VP[4], VT[2]}},
	{{VP[4], VT[2]}, {VP[3], VT[1]}, {VP[5], VT[3]}},
	{{VP[4], VT[3]}, {VP[5], VT[2]}, {VP[6], VT[1]}},
	{{VP[6], VT[1]}, {VP[5], VT[2]}, {VP[7], VT[0]}},
	{{VP[6], VT[0]}, {VP[7], VT[1]}, {VP[0], VT[2]}},
	{{VP[0], VT[2]}, {VP[7], VT[1]}, {VP[1], VT[3]}},
	{{VP[1], VT[0]}, {VP[7], VT[1]}, {VP[3], VT[2]}},
	{{VP[3], VT[2]}, {VP[7], VT[1]}, {VP[5], VT[3]}},
	{{VP[6], VT[0]}, {VP[0], VT[1]}, {VP[4], VT[2]}},
	{{VP[4], VT[2]}, {VP[0], VT[1]}, {VP[2], VT[3]}},
}

func Block(p Vector, material Material, tiles []int) []*Triangle {
	var result []*Triangle
	for i, t := range Triangles {
		tile := tiles[i/2]
		m := Vector{float64(tile%N) / N, float64(tile/N) / N, 0}
		v1, v2, v3 := t[0][0], t[1][0], t[2][0]
		v1, v2, v3 = v1.Add(p), v2.Add(p), v3.Add(p)
		t1, t2, t3 := t[0][1], t[1][1], t[2][1]
		t1, t2, t3 = t1.Add(m), t2.Add(m), t3.Add(m)
		result = append(result, NewTriangle(v1, v2, v3, t1, t2, t3, material))
	}
	return result
}

func main() {
	scene := Scene{}
	scene.Color = White
	texture, err := LoadTexture("examples/texture.png")
	if err != nil {
		panic(err)
	}
	material := GlossyMaterial(HexColor(0xFCFAE1), 1.1, Radians(20))
	material.Texture = texture
	var triangles []*Triangle
	for x := -10; x <= 10; x++ {
		for z := -10; z <= 10; z++ {
			h := rand.Intn(4)
			for y := 0; y <= h; y++ {
				p := Vector{float64(x), float64(y), float64(z)}
				tiles := Dirt
				if y == h {
					tiles = Grass
				}
				cube := Block(p, material, tiles)
				triangles = append(triangles, cube...)
			}
		}
	}
	mesh := NewMesh(triangles)
	scene.Add(mesh)
	camera := LookAt(Vector{-13, 11, -7}, Vector{0, 0, 0}, Vector{0, 1, 0}, 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
