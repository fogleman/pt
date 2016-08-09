package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"

	. "github.com/fogleman/pt/pt"
	colorful "github.com/lucasb-eyer/go-colorful"
)

func NewMaterial() Material {
	p := rand.Float64()
	p = p*0.7 + 0.2
	h := 164 - p*52
	s := 1 - p*0.58
	v := 0.15 + p*0.78
	c := colorful.Hsv(h, s, v)
	color := Color{c.R, c.G, c.B}.Pow(2.2)
	color = HexColor(0x468966)
	return GlossyMaterial(color, 1.4, Radians(20))
}

func LoadTriangles(path string) []*Triangle {
	materials := make(map[string]Material)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	var result []*Triangle
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
		v1 := V(x1, y1, z1)
		v2 := V(x2, y2, z2)
		v3 := V(x3, y3, z3)
		v := Vector{}
		t := NewTriangle(v1, v2, v3, v, v, v, materials[name])
		result = append(result, t)
	}
	return result
}

func main() {
	rand.Seed(6)
	floor := GlossyMaterial(HexColor(0xFCFFF5), 1.5, Radians(20))
	light := LightMaterial(HexColor(0xFFFFFF), 1, QuadraticAttenuation(0.002))
	triangles := LoadTriangles("examples/counties.csv")
	mesh := NewMesh(triangles)
	mesh.FitInside(Box{V(-1, -1, 0), V(1, 1, 1)}, V(0.5, 0.5, 0))
	scene := Scene{}
	scene.Add(mesh)
	scene.Add(NewCube(V(-100, -100, -1), V(100, 100, 0.03), floor))
	scene.Add(NewSphere(V(0, 4, 10), 4, light))
	camera := LookAt(V(0, -0.25, 2), V(0, 0, 0), V(0, 0, 1), 35)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}
