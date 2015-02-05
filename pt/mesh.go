package pt

import (
	"math/rand"
)

type Mesh struct {
	shapes    []Shape
	shapeTree *Tree
	color     Color
	material  Material
}

func NewMesh(shapes []Shape) *Mesh {
	mesh := &Mesh{}
	for _, shape := range shapes {
		triangle, ok := shape.(*Triangle)
		if ok {
			triangle.mesh = mesh
		}
	}
	mesh.shapes = shapes
	mesh.shapeTree = NewTree(shapes)
	return mesh
}

func (m *Mesh) SetColor(color Color) {
	m.color = color
}

func (m *Mesh) SetMaterial(material Material) {
	m.material = material
}

func (m *Mesh) Box() Box {
	return BoxForShapes(m.shapes)
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
