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

func LoadOBJ(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
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
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
