package pt

type Scene struct {
	Shapes []Shape
}

func (s *Scene) Add(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) Intersect(r Ray) Hit {
	hit := Hit{}
	hit.T = INF
	for _, shape := range s.Shapes {
		t := shape.Intersect(r)
		if t < hit.T {
			hit.T = t
			hit.Shape = shape
			hit.Position = r.Origin.Add(r.Direction.Mul(t))
		}
	}
	return hit
}
