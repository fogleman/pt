package main

import (
	"math/rand"

	"github.com/fogleman/pt/pt"
)

const N = 16

const A = 1.0 / 2048
const B = 1.0/N - A

var Dirt = []int{0, 0, 0, 0, 0, 0}
var Grass = []int{16, 32, 16, 0, 16, 16}

var V []pt.Vector = []pt.Vector{
	{-0.5, -0.5, 0.5},
	{0.5, -0.5, 0.5},
	{-0.5, 0.5, 0.5},
	{0.5, 0.5, 0.5},
	{-0.5, 0.5, -0.5},
	{0.5, 0.5, -0.5},
	{-0.5, -0.5, -0.5},
	{0.5, -0.5, -0.5},
}

var VT []pt.Vector = []pt.Vector{
	{A, A, 0},
	{B, A, 0},
	{A, B, 0},
	{B, B, 0},
}

var Triangles = [][3][2]pt.Vector{
	{{V[0], VT[0]}, {V[1], VT[1]}, {V[2], VT[2]}},
	{{V[2], VT[2]}, {V[1], VT[1]}, {V[3], VT[3]}},
	{{V[2], VT[0]}, {V[3], VT[1]}, {V[4], VT[2]}},
	{{V[4], VT[2]}, {V[3], VT[1]}, {V[5], VT[3]}},
	{{V[4], VT[3]}, {V[5], VT[2]}, {V[6], VT[1]}},
	{{V[6], VT[1]}, {V[5], VT[2]}, {V[7], VT[0]}},
	{{V[6], VT[0]}, {V[7], VT[1]}, {V[0], VT[2]}},
	{{V[0], VT[2]}, {V[7], VT[1]}, {V[1], VT[3]}},
	{{V[1], VT[0]}, {V[7], VT[1]}, {V[3], VT[2]}},
	{{V[3], VT[2]}, {V[7], VT[1]}, {V[5], VT[3]}},
	{{V[6], VT[0]}, {V[0], VT[1]}, {V[4], VT[2]}},
	{{V[4], VT[2]}, {V[0], VT[1]}, {V[2], VT[3]}},
}

func Cube(p pt.Vector, material pt.Material, tiles []int) []*pt.Triangle {
	var result []*pt.Triangle
	for i, t := range Triangles {
		tile := tiles[i/2]
		m := pt.Vector{float64(tile%N) / N, float64(tile/N) / N, 0}
		v1, v2, v3 := t[0][0], t[1][0], t[2][0]
		v1, v2, v3 = v1.Add(p), v2.Add(p), v3.Add(p)
		t1, t2, t3 := t[0][1], t[1][1], t[2][1]
		t1, t2, t3 = t1.Add(m), t2.Add(m), t3.Add(m)
		result = append(result, pt.NewTriangle(v1, v2, v3, t1, t2, t3, material))
	}
	return result
}

func main() {
	scene := pt.Scene{}
	texture, err := pt.LoadTexture("examples/texture.png")
	if err != nil {
		panic(err)
	}
	material := pt.GlossyMaterial(pt.HexColor(0xFCFAE1), 1.1, pt.Radians(20))
	material.Texture = texture
	var triangles []*pt.Triangle
	for x := -10; x <= 10; x++ {
		for z := -10; z <= 10; z++ {
			h := rand.Intn(4)
			for y := 0; y <= h; y++ {
				p := pt.Vector{float64(x), float64(y), float64(z)}
				tiles := Dirt
				if y == h {
					tiles = Grass
				}
				cube := Cube(p, material, tiles)
				triangles = append(triangles, cube...)
			}
		}
	}
	mesh := pt.NewMesh(triangles)
	scene.Add(mesh)
	camera := pt.LookAt(pt.Vector{-13, 11, -7}, pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, -1, 4, 4)
}
