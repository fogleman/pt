package pt

import (
	"bufio"
	"os"
	"strings"
)

func LoadOBJ(path string, parent Material) (*Mesh, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var vs, vts, vns []Vector
	vs = append(vs, Vector{})   // 1-based indexing
	vts = append(vts, Vector{}) // 1-based indexing
	vns = append(vns, Vector{}) // 1-based indexing
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
			v := Vector{f[0], f[1], f[2]}
			vts = append(vts, v)
		case "vn":
			f := ParseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vns = append(vns, v)
		case "f":
			var fvs, fvts, fvns []string
			for _, arg := range args {
				vertex := strings.Split(arg+"//", "/")
				fvs = append(fvs, vertex[0])
				fvts = append(fvts, vertex[1])
				fvns = append(fvns, vertex[2])
			}
			ivs := ParseInts(fvs)
			ivts := ParseInts(fvts)
			ivns := ParseInts(fvns)
			for i := 1; i < len(ivs)-1; i++ {
				i1, i2, i3 := 0, i, i+1
				t := Triangle{}
				t.material = material
				t.v1 = vs[ivs[i1]]
				t.v2 = vs[ivs[i2]]
				t.v3 = vs[ivs[i3]]
				t.t1 = vts[ivts[i1]]
				t.t2 = vts[ivts[i2]]
				t.t3 = vts[ivts[i3]]
				t.n1 = vns[ivns[i1]]
				t.n2 = vns[ivns[i2]]
				t.n3 = vns[ivns[i3]]
				min := t.v1.Min(t.v2).Min(t.v3)
				max := t.v1.Max(t.v2).Max(t.v3)
				t.box = Box{min, max}
				t.FixNormals()
				triangles = append(triangles, &t)
			}
		}
	}
	return NewMesh(triangles), scanner.Err()
}
