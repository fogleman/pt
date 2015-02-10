package pt

import (
	"math/rand"
)

type Mesh struct {
	box       Box
	triangles []*Triangle
	tree      *Tree
}

func NewMesh(triangles []*Triangle) *Mesh {
	box := BoxForTriangles(triangles)
	return &Mesh{box, triangles, nil}
}

func (m *Mesh) Compile() {
	if m.tree == nil {
		shapes := make([]Shape, len(m.triangles))
		for i, triangle := range m.triangles {
			shapes[i] = triangle
		}
		m.tree = NewTree(shapes)
	}
}

func (m *Mesh) Box() Box {
	return m.box
}

func (m *Mesh) Intersect(r Ray) Hit {
	return m.tree.Intersect(r)
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

func (m *Mesh) UpdateBox() {
	m.box = BoxForTriangles(m.triangles)
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
	m.UpdateBox()
}

func (m *Mesh) MoveTo(position, anchor Vector) {
	offset := position.Sub(m.box.Anchor(anchor))
	for _, t := range m.triangles {
		t.v1 = t.v1.Add(offset)
		t.v2 = t.v2.Add(offset)
		t.v3 = t.v3.Add(offset)
		t.UpdateBox()
	}
	m.UpdateBox()
	m.tree = nil // dirty
}
