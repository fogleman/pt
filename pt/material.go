package pt

import (
	"bufio"
	"os"
	"strings"
)

type Material struct {
	Color Color
	Index float64 // refractive index
	Gloss float64 // reflection cone angle in radians
	Tint  float64 // specular tinting
}

func DiffuseMaterial(color Color) Material {
	return Material{color, 1, 0, 0}
}

func RefractiveMaterial(color Color, index float64) Material {
	return Material{color, index, 0, 0}
}

func LoadMTL(path string, parent Material, materials map[string]*Material) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parentCopy := parent
	material := &parentCopy
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		keyword := fields[0]
		args := fields[1:]
		switch keyword {
		case "newmtl":
			parentCopy := parent
			material = &parentCopy
			materials[args[0]] = material
		case "Kd":
			c := ParseFloats(args)
			material.Color = Color{c[0], c[1], c[2]}
		}
	}
	return scanner.Err()
}
