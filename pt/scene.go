package pt

import "sync/atomic"

type Scene struct {
	color      Color
	texture    Texture
	visibility float64
	shapes     []Shape
	lights     []Shape
	dlights    []DirectionalLight
	tree       *Tree
	rays       uint64
}

func (s *Scene) Shapes() []Shape {
	return s.shapes
}

func (s *Scene) SetColor(color Color) {
	s.color = color
}

func (s *Scene) SetTexture(texture Texture) {
	s.texture = texture
}

func (s *Scene) SetVisibility(visibility float64) {
	s.visibility = visibility
}

func (s *Scene) AddDirectionalLight(c Color, d Vector, t float64) {
	s.dlights = append(s.dlights, DirectionalLight{c, d.Normalize(), t})
}

func (s *Scene) Compile() {
	for _, shape := range s.shapes {
		shape.Compile()
	}
	if s.tree == nil {
		s.tree = NewTree(s.shapes)
	}
}

func (s *Scene) Add(shape Shape) {
	s.shapes = append(s.shapes, shape)
	if shape.Material(Vector{}).Emittance > 0 {
		s.lights = append(s.lights, shape)
	}
}

func (s *Scene) RayCount() uint64 {
	return atomic.LoadUint64(&s.rays)
}

func (s *Scene) Intersect(r Ray) Hit {
	atomic.AddUint64(&s.rays, 1)
	return s.tree.Intersect(r)
}

func (s *Scene) Shadow(r Ray, light Shape, max float64) bool {
	hit := s.Intersect(r)
	return hit.Shape != light && hit.T < max
}
