package pt

import (
	"math"
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

func smoothNormalsThreshold(normal Vector, normals []Vector, threshold float64) Vector {
	result := Vector{}
	for _, x := range normals {
		if x.Dot(normal) >= threshold {
			result = result.Add(x)
		}
	}
	return result.Normalize()
}

func (m *Mesh) SmoothNormalsThreshold(radians float64) {
	threshold := math.Cos(radians)
	lookup := make(map[Vector][]Vector)
	for _, t := range m.triangles {
		lookup[t.v1] = append(lookup[t.v1], t.n1)
		lookup[t.v2] = append(lookup[t.v2], t.n2)
		lookup[t.v3] = append(lookup[t.v3], t.n3)
	}
	for _, t := range m.triangles {
		t.n1 = smoothNormalsThreshold(t.n1, lookup[t.v1], threshold)
		t.n2 = smoothNormalsThreshold(t.n2, lookup[t.v2], threshold)
		t.n3 = smoothNormalsThreshold(t.n3, lookup[t.v3], threshold)
	}
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

func (m *Mesh) MoveTo(position, anchor Vector) {
	matrix := Translate(position.Sub(m.box.Anchor(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) FitInside(box Box, anchor Vector) {
	scale := box.Size().Div(m.Box().Size()).MinComponent()
	extra := box.Size().Sub(m.Box().Size().MulScalar(scale))
	matrix := Identity()
	matrix = matrix.Translate(m.Box().Min.MulScalar(-1))
	matrix = matrix.Scale(Vector{scale, scale, scale})
	matrix = matrix.Translate(box.Min.Add(extra.Mul(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) Transform(matrix Matrix) {
	for _, t := range m.triangles {
		t.v1 = matrix.MulPosition(t.v1)
		t.v2 = matrix.MulPosition(t.v2)
		t.v3 = matrix.MulPosition(t.v3)
		t.n1 = matrix.MulDirection(t.n1)
		t.n2 = matrix.MulDirection(t.n2)
		t.n3 = matrix.MulDirection(t.n3)
		t.UpdateBox()
	}
	m.UpdateBox()
	m.tree = nil // dirty
}
