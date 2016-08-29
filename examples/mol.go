package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fogleman/mol/mol"
	. "github.com/fogleman/pt/pt"
)

func GetColor(name string) Color {
	switch name {
	case "1", "2", "3":
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
	case "S":
		return HexColor(0xFFD34E)
	case "Co":
		return HexColor(0xD0D0D0)
	default:
		fmt.Println(name)
		return White
	}
}

func GetMaterial(name string) Material {
	switch name {
	case "1", "2", "3":
		return GlossyMaterial(GetColor(name), 1.1, Radians(30))
	default:
		return GlossyMaterial(GetColor(name), 1.3, Radians(30))
	}
}

func main() {
	scene := Scene{}

	molecule, err := mol.ParseFile(os.Args[1])
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

	center.Z -= 4
	fmt.Println(eye)
	fmt.Println(center)
	fmt.Println(up)

	// light coordinate system
	m := LookAtMatrix(eye, center, up).Translate(center.Sub(eye))
	d := 50.0
	a := V(-1, 0.5, -1)
	b := V(1, -0.25, -1)
	c := V(-1, -0.25, 1)
	a = m.MulPosition(a.Normalize().MulScalar(d))
	b = m.MulPosition(b.Normalize().MulScalar(d))
	c = m.MulPosition(c.Normalize().MulScalar(d))
	light := LightMaterial(White, 1000)
	scene.Add(NewSphere(a, 2, light))
	scene.Add(NewSphere(b, 1, light))
	scene.Add(NewSphere(c, 1, light))

	camera := LookAt(eye, center, up, cam.Fovy/1.7)
	sampler := NewSampler(16, 8)
	sampler.SpecularMode = SpecularModeAll
	IterativeRender("out%03d.png", 10000, &scene, &camera, sampler, 2560, 1440, -1)
}
