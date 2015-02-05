package pt

import (
	"math/rand"
)

type Mesh struct {
	color     Color
	material  Material
	shapeTree *Tree
}

func NewMesh(color Color, material Material) *Mesh {
	return &Mesh{color, material, nil}
}

func (m *Mesh) LoadOBJ(path string) error {
	shapes, err := LoadOBJ(path)
	if err != nil {
		return err
	}
	m.SetShapes(shapes)
	return nil
}

func (m *Mesh) SetShapes(shapes []Shape) {
	for _, shape := range shapes {
		if triangle, ok := shape.(*Triangle); ok {
			triangle.mesh = m
			triangle.FixNormals()
		}
	}
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
