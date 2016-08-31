package pt

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseIndex(value string, length int) int {
	parsed, _ := strconv.ParseInt(value, 0, 0)
	n := int(parsed)
	if n < 0 {
		n += length
	}
	return n
}

func LoadOBJ(path string, parent Material) (*Mesh, error) {
	fmt.Printf("Loading OBJ: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	vs := make([]Vector, 1, 1024)  // 1-based indexing
	vts := make([]Vector, 1, 1024) // 1-based indexing
	vns := make([]Vector, 1, 1024) // 1-based indexing
	var triangles []*Triangle
	materials := make(map[string]*Material)
	material := &parent
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		keyword := fields[0]
		args := fields[1:]
		switch keyword {
		case "mtllib":
			p := RelativePath(path, args[0])
			if err := LoadMTL(p, parent, materials); err != nil {
				return nil, err
			}
		case "usemtl":
			if m, ok := materials[args[0]]; ok {
				material = m
			}
		case "v":
			f := ParseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vs = append(vs, v)
		case "vt":
			f := ParseFloats(args)
			v := Vector{f[0], f[1], 0}
			vts = append(vts, v)
		case "vn":
			f := ParseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vns = append(vns, v)
		case "f":
			fvs := make([]int, len(args))
			fvts := make([]int, len(args))
			fvns := make([]int, len(args))
			for i, arg := range args {
				vertex := strings.Split(arg+"//", "/")
				fvs[i] = parseIndex(vertex[0], len(vs))
				fvts[i] = parseIndex(vertex[1], len(vts))
				fvns[i] = parseIndex(vertex[2], len(vns))
			}
			for i := 1; i < len(fvs)-1; i++ {
				i1, i2, i3 := 0, i, i+1
				t := Triangle{}
				t.Material = material
				t.V1 = vs[fvs[i1]]
				t.V2 = vs[fvs[i2]]
				t.V3 = vs[fvs[i3]]
				t.T1 = vts[fvts[i1]]
				t.T2 = vts[fvts[i2]]
				t.T3 = vts[fvts[i3]]
				t.N1 = vns[fvns[i1]]
				t.N2 = vns[fvns[i2]]
				t.N3 = vns[fvns[i3]]
				t.FixNormals()
				triangles = append(triangles, &t)
			}
		}
	}
	return NewMesh(triangles), scanner.Err()
}

func LoadMTL(path string, parent Material, materials map[string]*Material) error {
	fmt.Printf("Loading MTL: %s\n", path)
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
		case "Ke":
			c := ParseFloats(args)
			max := math.Max(math.Max(c[0], c[1]), c[2])
			if max > 0 {
				material.Color = Color{c[0] / max, c[1] / max, c[2] / max}
				material.Emittance = max
			}
		case "Kd":
			c := ParseFloats(args)
			material.Color = Color{c[0], c[1], c[2]}
		case "map_Kd":
			p := RelativePath(path, args[0])
			material.Texture = GetTexture(p)
		case "map_bump":
			p := RelativePath(path, args[0])
			material.NormalTexture = GetTexture(p).Pow(1 / 2.2)
		}
	}
	return scanner.Err()
}
