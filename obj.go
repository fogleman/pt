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
	color := HexColor(0xEFC94C)
	material := DiffuseMaterial()
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	var vs []Vector
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
		if keyword == "f" {
			indexes := parseInts(args)
			for i := 1; i < len(indexes)-1; i++ {
				a, b, c := indexes[0], indexes[i], indexes[i+1]
				shape := NewTriangle(vs[a-1], vs[b-1], vs[c-1], color, material)
				shapes = append(shapes, shape)
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}
