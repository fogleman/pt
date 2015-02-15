package pt

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type STLHeader struct {
	_     [80]uint8
	Count uint32
}

type STLTriangle struct {
	_, V1, V2, V3 [3]float32
	_             uint16
}

func LoadBinarySTL(path string, material Material) (*Mesh, error) {
	fmt.Printf("Loading STL (Binary): %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	header := STLHeader{}
	if err := binary.Read(file, binary.LittleEndian, &header); err != nil {
		return nil, err
	}
	count := int(header.Count)
	triangles := make([]*Triangle, count)
	for i := 0; i < count; i++ {
		d := STLTriangle{}
		if err := binary.Read(file, binary.LittleEndian, &d); err != nil {
			return nil, err
		}
		t := Triangle{}
		t.material = &material
		t.v1 = Vector{float64(d.V1[0]), float64(d.V1[1]), float64(d.V1[2])}
		t.v2 = Vector{float64(d.V2[0]), float64(d.V2[1]), float64(d.V2[2])}
		t.v3 = Vector{float64(d.V3[0]), float64(d.V3[1]), float64(d.V3[2])}
		t.UpdateBox()
		t.FixNormals()
		triangles[i] = &t
	}
	return NewMesh(triangles), nil
}

func LoadSTL(path string, material Material) (*Mesh, error) {
	fmt.Printf("Loading STL (ASCII): %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var vertexes []Vector
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 4 && fields[0] == "vertex" {
			f := ParseFloats(fields[1:])
			v := Vector{f[0], f[1], f[2]}
			vertexes = append(vertexes, v)
		}
	}
	var triangles []*Triangle
	for i := 0; i < len(vertexes); i += 3 {
		t := Triangle{}
		t.material = &material
		t.v1 = vertexes[i+0]
		t.v2 = vertexes[i+1]
		t.v3 = vertexes[i+2]
		t.UpdateBox()
		t.FixNormals()
		triangles = append(triangles, &t)
	}
	return NewMesh(triangles), scanner.Err()
}
