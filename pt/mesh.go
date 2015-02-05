package pt

import (
	"math/rand"
)

type Mesh struct {
	color     Color
	material  Material
	triangles []*Triangle
	shapeTree *Tree
}

func NewMesh(color Color, material Material) *Mesh {
	return &Mesh{color, material, nil, nil}
}

func (m *Mesh) LoadOBJ(path string) error {
	triangles, err := LoadOBJ(path)
	if err != nil {
		return err
	}
	m.SetTriangles(triangles)
	return nil
}

func (m *Mesh) SetTriangles(triangles []*Triangle) {
	for _, triangle := range triangles {
		triangle.mesh = m
		triangle.FixNormals()
	}
	shapes := make([]Shape, len(triangles))
	for i := range triangles {
		shapes[i] = triangles[i]
	}
	m.triangles = triangles
	m.shapeTree = NewTree(shapes)
}

func (m *Mesh) Box() Box {
	return m.shapeTree.box
}

func (m *Mesh) Intersect(r Ray) Hit {
	return m.shapeTree.Intersect(r)
}

func (m *Mesh) Color(p Vector) Color {
	return Color{} // not implemented
}

func (m *Mesh) Material(p Vector) Material {
	return Material{} // not implemented
}

func (m *Mesh) Normal(p Vector) Vector {
	return Vector{} // not implemented
}

func (m *Mesh) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{} // not implemented
}
