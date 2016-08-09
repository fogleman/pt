package pt

type Func func(x, y float64) float64

type Function struct {
	function Func
	box      Box
	material Material
}

func NewFunction(function Func, box Box, material Material) Shape {
	return &Function{function, box, material}
}

func (f *Function) Compile() {
}

func (f *Function) BoundingBox() Box {
	return f.box
}

func (f *Function) Contains(v Vector) bool {
	return v.Z < f.function(v.X, v.Y)
}

func (f *Function) Intersect(ray Ray) Hit {
	step := 1.0 / 32
	sign := f.Contains(ray.Position(step))
	for t := step; t < 12; t += step {
		v := ray.Position(t)
		if f.Contains(v) != sign && f.box.Contains(v) {
			return Hit{f, t - step, nil}
		}
	}
	return NoHit
}

func (f *Function) ColorAt(p Vector) Color {
	if f.material.Texture == nil {
		return f.material.Color
	}
	x1, x2 := f.box.Min.X, f.box.Max.X
	y1, y2 := f.box.Min.Y, f.box.Max.Y
	u := (p.X - x1) / (x2 - x1)
	v := (p.Y - y1) / (y2 - y1)
	return f.material.Texture.Sample(u, v)
}

func (f *Function) MaterialAt(p Vector) Material {
	return f.material
}

func (f *Function) NormalAt(p Vector) Vector {
	eps := 1e-3
	v := Vector{
		f.function(p.X-eps, p.Y) - f.function(p.X+eps, p.Y),
		f.function(p.X, p.Y-eps) - f.function(p.X, p.Y+eps),
		2 * eps,
	}
	return v.Normalize()
}
