package pt

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseFloats(items []string) []float64 {
	var result []float64
	for _, item := range items {
		f, _ := strconv.ParseFloat(item, 64)
		result = append(result, f)
	}
	return result
}

func parseInts(items []string) []int {
	var result []int
	for _, item := range items {
		f, _ := strconv.ParseInt(item, 0, 0)
		result = append(result, int(f))
	}
	return result
}

func LoadOBJ(path string) (shapes []Shape, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	var vs, vts, vns []Vector
	vs = append(vs, Vector{})   // 1-based indexing
	vts = append(vts, Vector{}) // 1-based indexing
	vns = append(vns, Vector{}) // 1-based indexing
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		keyword := fields[0]
		args := fields[1:]
		if keyword == "v" {
			f := parseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vs = append(vs, v)
		}
		if keyword == "vt" {
			f := parseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vts = append(vts, v)
		}
		if keyword == "vn" {
			f := parseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vns = append(vns, v)
		}
		if keyword == "f" {
			var fvs, fvts, fvns []string
			for _, arg := range args {
				vertex := strings.Split(arg+"//", "/")
				fvs = append(fvs, vertex[0])
				fvts = append(fvts, vertex[1])
				fvns = append(fvns, vertex[2])
			}
			ivs := parseInts(fvs)
			ivts := parseInts(fvts)
			ivns := parseInts(fvns)
			for i := 1; i < len(ivs)-1; i++ {
				i1, i2, i3 := 0, i, i+1
				t := Triangle{}
				t.v1 = vs[ivs[i1]]
				t.v2 = vs[ivs[i2]]
				t.v3 = vs[ivs[i3]]
				t.t1 = vts[ivts[i1]]
				t.t2 = vts[ivts[i2]]
				t.t3 = vts[ivts[i3]]
				t.n1 = vns[ivns[i1]]
				t.n2 = vns[ivns[i2]]
				t.n3 = vns[ivns[i3]]
				shapes = append(shapes, &t)
			}
		}
	}
	err = scanner.Err()
	return
}
