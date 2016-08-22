package main

import (
	"fmt"
	"strconv"

	"github.com/fogleman/mol/mol"
	. "github.com/fogleman/pt/pt"
)

func GetColor(name string) Color {
	switch name {
	case "1":
		return Color{0.1, 0.1, 0.1}
	case "2":
		return Color{0.1, 0.1, 0.1}
	case "H":
		return HexColor(0xECF0F1)
	case "C":
		return HexColor(0x222222)
	case "N":
		return HexColor(0x3498DB)
	case "O":
		return HexColor(0xE74C3C)
	case "P":
		return HexColor(0xFF9800)
	case "Co":
		return HexColor(0xD0D0D0)
	default:
		fmt.Println(name)
		return Color{1, 1, 1}
	}
}

func GetMaterial(name string) Material {
	switch name {
	case "1", "2":
		return GlossyMaterial(GetColor(name), 1.5, Radians(10))
	default:
		return GlossyMaterial(GetColor(name), 1.3, Radians(30))
	}
}

func main() {
	scene := Scene{}

	molecule, err := mol.ParseFile("../mol/examples/caffeine.txt")
	if err != nil {
		panic(err)
	}

	spheres, cylinders := molecule.Solids()

	for _, s := range spheres {
		center := V(s.Center.X, s.Center.Y, s.Center.Z)
		radius := s.Radius * 0.66
		material := GetMaterial(s.Symbol)
		scene.Add(NewSphere(center, radius, material))
	}

	for _, c := range cylinders {
		v0 := V(c.A.X, c.A.Y, c.A.Z)
		v1 := V(c.B.X, c.B.Y, c.B.Z)
		radius := c.Radius
		material := GetMaterial(strconv.Itoa(c.Type))
		scene.Add(NewTransformedCylinder(v0, v1, radius, material))
	}

	// camera
	cam := molecule.Camera()
	eye := V(cam.Eye.X, cam.Eye.Y, cam.Eye.Z)
	center := V(cam.Center.X, cam.Center.Y, cam.Center.Z)
	up := V(cam.Up.X, cam.Up.Y, cam.Up.Z)
	camera := LookAt(eye, center, up, cam.Fovy)

	// light coordinate system
	m := LookAtMatrix(eye, center, up)
	d := center.Sub(eye).Length()
	light := LightMaterial(Color{1, 1, 1}, 2000)
	scene.Add(NewSphere(m.MulPosition(V(-d/2, d/2, -d)), 2, light))
	scene.Add(NewSphere(m.MulPosition(V(d/2, -d/2, -d)), 1, light))
	scene.Add(NewSphere(m.MulPosition(V(d, 0, d)), 1, light))

	sampler := NewSampler(64, 8)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1024/1, 1024/1, -1)
}
