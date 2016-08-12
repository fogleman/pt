package pt

type Func func(x, y float64) float64

type Function struct {
	Function Func
	Box      Box
	Material Material
}

func NewFunction(function Func, box Box, material Material) Shape {
	return &Function{function, box, material}
}

func (f *Function) Compile() {
}

func (f *Function) BoundingBox() Box {
	return f.Box
}

func (f *Function) Contains(v Vector) bool {
	return v.Z < f.Function(v.X, v.Y)
}

func (f *Function) Intersect(ray Ray) Hit {
	step := 1.0 / 32
	sign := f.Contains(ray.Position(step))
	for t := step; t < 12; t += step {
		v := ray.Position(t)
		if f.Contains(v) != sign && f.Box.Contains(v) {
			return Hit{f, t - step, nil}
		}
	}
	return NoHit
}

func (f *Function) UV(p Vector) Vector {
	x1, x2 := f.Box.Min.X, f.Box.Max.X
	y1, y2 := f.Box.Min.Y, f.Box.Max.Y
	u := (p.X - x1) / (x2 - x1)
	v := (p.Y - y1) / (y2 - y1)
	return Vector{u, v, 0}
}

func (f *Function) MaterialAt(p Vector) Material {
	return f.Material
}

func (f *Function) NormalAt(p Vector) Vector {
	eps := 1e-3
	v := Vector{
		f.Function(p.X-eps, p.Y) - f.Function(p.X+eps, p.Y),
		f.Function(p.X, p.Y-eps) - f.Function(p.X, p.Y+eps),
		2 * eps,
	}
	return v.Normalize()
}
