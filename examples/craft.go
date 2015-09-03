package main

import (
	"fmt"

	"github.com/fogleman/pt/pt"
)

const N = 16

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
	{0, 1 - 0, 0},
	{1.0 / N, 1 - 0, 0},
	{0, 1 - 1.0/N, 0},
	{1.0 / N, 1 - 1.0/N, 0},
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

func Cube(tiles []int, material pt.Material) []*pt.Triangle {
	var result []*pt.Triangle
	for i, t := range Triangles {
		tile := tiles[i/2]
		m := pt.Vector{float64(tile%N) / N, float64(tile/N) / N, 0}
		v1, v2, v3 := t[0][0], t[1][0], t[2][0]
		t1, t2, t3 := t[0][1], t[1][1], t[2][1]
		t1, t2, t3 = t1.Add(m), t2.Add(m), t3.Add(m)
		fmt.Println(t1, t2, t3)
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
	triangles = append(triangles, Cube([]int{0, 1, 2, 3, 4, 5}, material)...)
	mesh := pt.NewMesh(triangles)
	scene.Add(mesh)
	camera := pt.LookAt(pt.Vector{3, 1, 3}, pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 512*2, 512*2, -1, 16, 4)
}
