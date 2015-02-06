package pt

import (
	"math/rand"
)

type Mesh struct {
	triangles []*Triangle
	shapeTree *Tree
}

func NewMesh(triangles []*Triangle) *Mesh {
	shapes := make([]Shape, len(triangles))
	for i := range triangles {
		shapes[i] = triangles[i]
	}
	return &Mesh{triangles, NewTree(shapes)}
}

func (m *Mesh) SmoothNormals() {
	lookup := make(map[Vector]Vector)
	for _, t := range m.triangles {
		lookup[t.v1] = lookup[t.v1].Add(t.n1)
		lookup[t.v2] = lookup[t.v2].Add(t.n2)
		lookup[t.v3] = lookup[t.v3].Add(t.n3)
	}
	for k, v := range lookup {
		lookup[k] = v.Normalize()
	}
	for _, t := range m.triangles {
		t.n1 = lookup[t.v1]
		t.n2 = lookup[t.v2]
		t.n3 = lookup[t.v3]
	}
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
