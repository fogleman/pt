package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"

	"github.com/fogleman/pt/pt"
	"github.com/lucasb-eyer/go-colorful"
)

func NewMaterial() pt.Material {
	p := rand.Float64()
	p = p*0.7 + 0.2
	h := 164 - p*52
	s := 1 - p*0.58
	v := 0.15 + p*0.78
	c := colorful.Hsv(h, s, v)
	color := pt.Color{c.R, c.G, c.B}.Pow(2.2)
	color = pt.HexColor(0x468966)
	return pt.GlossyMaterial(color, 1.4, pt.Radians(20))
}

func LoadTriangles(path string) []*pt.Triangle {
	materials := make(map[string]pt.Material)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	var result []*pt.Triangle
	for _, row := range rows {
		name := row[0]
		if _, ok := materials[name]; !ok {
			materials[name] = NewMaterial()
		}
		x1, _ := strconv.ParseFloat(row[1], 64)
		y1, _ := strconv.ParseFloat(row[2], 64)
		z1, _ := strconv.ParseFloat(row[3], 64)
		x2, _ := strconv.ParseFloat(row[4], 64)
		y2, _ := strconv.ParseFloat(row[5], 64)
		z2, _ := strconv.ParseFloat(row[6], 64)
		x3, _ := strconv.ParseFloat(row[7], 64)
		y3, _ := strconv.ParseFloat(row[8], 64)
		z3, _ := strconv.ParseFloat(row[9], 64)
		v1 := pt.Vector{x1, y1, z1}
		v2 := pt.Vector{x2, y2, z2}
		v3 := pt.Vector{x3, y3, z3}
		v := pt.Vector{}
		t := pt.NewTriangle(v1, v2, v3, v, v, v, materials[name])
		result = append(result, t)
	}
	return result
}

func main() {
	rand.Seed(6)
	floor := pt.GlossyMaterial(pt.HexColor(0xFCFFF5), 1.5, pt.Radians(20))
	light := pt.LightMaterial(pt.HexColor(0xFFFFFF), 1, pt.QuadraticAttenuation(0.002))
	triangles := LoadTriangles("examples/counties.csv")
	mesh := pt.NewMesh(triangles)
	mesh.FitInside(pt.Box{pt.Vector{-1, -1, 0}, pt.Vector{1, 1, 1}}, pt.Vector{0.5, 0.5, 0})
	scene := pt.Scene{}
	scene.Add(mesh)
	scene.Add(pt.NewCube(pt.Vector{-100, -100, -1}, pt.Vector{100, 100, 0.03}, floor))
	scene.Add(pt.NewSphere(pt.Vector{0, 4, 10}, 4, light))
	camera := pt.LookAt(pt.Vector{0, -0.25, 2}, pt.Vector{0, 0, 0}, pt.Vector{0, 0, 1}, 35)
	pt.IterativeRender("out%03d.png", 1000, &scene, &camera, 2560/2, 1440/2, -1, 4, 4)
}
