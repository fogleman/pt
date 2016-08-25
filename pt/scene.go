package pt

import "sync/atomic"

type Scene struct {
	Color             Color
	Texture           Texture
	TextureAngle      float64
	Shapes            []Shape
	Lights            []Shape
	DirectionalLights []*DirectionalLight
	tree              *Tree
	rays              uint64
}

func (s *Scene) Compile() {
	for _, shape := range s.Shapes {
		shape.Compile()
	}
	if s.tree == nil {
		s.tree = NewTree(s.Shapes)
	}
}

func (s *Scene) Add(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
	if shape.MaterialAt(Vector{}).Emittance > 0 {
		s.Lights = append(s.Lights, shape)
	}
}

func (s *Scene) AddDirectionalLight(light *DirectionalLight) {
	s.DirectionalLights = append(s.DirectionalLights, light)
}

func (s *Scene) RayCount() uint64 {
	return atomic.LoadUint64(&s.rays)
}

func (s *Scene) Intersect(r Ray) Hit {
	atomic.AddUint64(&s.rays, 1)
	return s.tree.Intersect(r)
}
