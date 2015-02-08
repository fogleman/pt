package pt

import (
	"encoding/binary"
	"fmt"
	"os"
)

type STLHeader struct {
	_     [80]uint8
	Count uint32
}

type STLTriangle struct {
	N, V1, V2, V3 [3]float32
	_             uint16
}

func LoadBinarySTL(path string, material Material) (*Mesh, error) {
	fmt.Printf("Loading %s... ", path)
	defer fmt.Println("OK")
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
	triangles := make([]*Triangle, 0, count)
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
		normal := Vector{float64(d.N[0]), float64(d.N[1]), float64(d.N[2])}
		t.n1 = normal
		t.n2 = normal
		t.n3 = normal
		t.UpdateBox()
		t.FixNormals()
		triangles = append(triangles, &t)
	}
	return NewMesh(triangles), nil
}
