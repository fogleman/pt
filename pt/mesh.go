package pt

import (
	"math"

	embree "github.com/fogleman/go-embree"
)

type Mesh struct {
	Triangles []*Triangle
	box       *Box
	mesh      *embree.Mesh
}

func NewMesh(triangles []*Triangle) *Mesh {
	return &Mesh{triangles, nil, nil}
}

func (m *Mesh) dirty() {
	m.box = nil
	m.mesh = nil
}

func (m *Mesh) Copy() *Mesh {
	triangles := make([]*Triangle, len(m.Triangles))
	for i, t := range m.Triangles {
		a := *t
		triangles[i] = &a
	}
	return NewMesh(triangles)
}

func (m *Mesh) Compile() {
	if m.mesh == nil {
		triangles := make([]embree.Triangle, len(m.Triangles))
		for i, t := range m.Triangles {
			triangles[i] = embree.Triangle{
				embree.Vector{t.V1.X, t.V1.Y, t.V1.Z},
				embree.Vector{t.V2.X, t.V2.Y, t.V2.Z},
				embree.Vector{t.V3.X, t.V3.Y, t.V3.Z},
			}
		}
		m.mesh = embree.NewMesh(triangles)
	}
}

func (a *Mesh) Add(b *Mesh) {
	a.Triangles = append(a.Triangles, b.Triangles...)
	a.dirty()
}

func (m *Mesh) BoundingBox() Box {
	if m.box == nil {
		min := m.Triangles[0].V1
		max := m.Triangles[0].V1
		for _, t := range m.Triangles {
			min = min.Min(t.V1).Min(t.V2).Min(t.V3)
			max = max.Max(t.V1).Max(t.V2).Max(t.V3)
		}
		m.box = &Box{min, max}
	}
	return *m.box
}

func (m *Mesh) Intersect(r Ray) Hit {
	ray := embree.Ray{
		embree.Vector{r.Origin.X, r.Origin.Y, r.Origin.Z},
		embree.Vector{r.Direction.X, r.Direction.Y, r.Direction.Z},
	}
	hit := m.mesh.Intersect(ray)
	if hit.Index >= 0 {
		return Hit{m.Triangles[hit.Index], hit.T - 1e-5, nil}
	} else {
		return NoHit
	}
}

func (m *Mesh) UV(p Vector) Vector {
	return Vector{} // not implemented
}

func (m *Mesh) MaterialAt(p Vector) Material {
	return Material{} // not implemented
}

func (m *Mesh) NormalAt(p Vector) Vector {
	return Vector{} // not implemented
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
	for _, t := range m.Triangles {
		lookup[t.V1] = append(lookup[t.V1], t.N1)
		lookup[t.V2] = append(lookup[t.V2], t.N2)
		lookup[t.V3] = append(lookup[t.V3], t.N3)
	}
	for _, t := range m.Triangles {
		t.N1 = smoothNormalsThreshold(t.N1, lookup[t.V1], threshold)
		t.N2 = smoothNormalsThreshold(t.N2, lookup[t.V2], threshold)
		t.N3 = smoothNormalsThreshold(t.N3, lookup[t.V3], threshold)
	}
}

func (m *Mesh) SmoothNormals() {
	lookup := make(map[Vector]Vector)
	for _, t := range m.Triangles {
		lookup[t.V1] = lookup[t.V1].Add(t.N1)
		lookup[t.V2] = lookup[t.V2].Add(t.N2)
		lookup[t.V3] = lookup[t.V3].Add(t.N3)
	}
	for k, v := range lookup {
		lookup[k] = v.Normalize()
	}
	for _, t := range m.Triangles {
		t.N1 = lookup[t.V1]
		t.N2 = lookup[t.V2]
		t.N3 = lookup[t.V3]
	}
}

func (m *Mesh) UnitCube() {
	m.FitInside(Box{Vector{}, Vector{1, 1, 1}}, Vector{})
	m.MoveTo(Vector{}, Vector{0.5, 0.5, 0.5})
}

func (m *Mesh) MoveTo(position, anchor Vector) {
	matrix := Translate(position.Sub(m.BoundingBox().Anchor(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) FitInside(box Box, anchor Vector) {
	scale := box.Size().Div(m.BoundingBox().Size()).MinComponent()
	extra := box.Size().Sub(m.BoundingBox().Size().MulScalar(scale))
	matrix := Identity()
	matrix = matrix.Translate(m.BoundingBox().Min.Negate())
	matrix = matrix.Scale(Vector{scale, scale, scale})
	matrix = matrix.Translate(box.Min.Add(extra.Mul(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) Transform(matrix Matrix) {
	for _, t := range m.Triangles {
		t.V1 = matrix.MulPosition(t.V1)
		t.V2 = matrix.MulPosition(t.V2)
		t.V3 = matrix.MulPosition(t.V3)
		t.N1 = matrix.MulDirection(t.N1)
		t.N2 = matrix.MulDirection(t.N2)
		t.N3 = matrix.MulDirection(t.N3)
	}
	m.dirty()
}

func (m *Mesh) SetMaterial(material Material) {
	for _, t := range m.Triangles {
		t.Material = &material
	}
}

func (m *Mesh) SaveSTL(path string) error {
	return SaveSTL(path, m)
}
