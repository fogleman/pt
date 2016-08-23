package pt

import "math"

type Cylinder struct {
	Radius   float64
	Z0, Z1   float64
	Material Material
}

func NewCylinder(radius, z0, z1 float64, material Material) *Cylinder {
	return &Cylinder{radius, z0, z1, material}
}

func NewTransformedCylinder(v0, v1 Vector, radius float64, material Material) Shape {
	up := Vector{0, 0, 1}
	d := v1.Sub(v0)
	z := d.Length()
	a := math.Acos(d.Normalize().Dot(up))
	m := Translate(v0)
	if a != 0 {
		u := d.Cross(up).Normalize()
		m = Rotate(u, a).Translate(v0)
	}
	c := NewCylinder(radius, 0, z, material)
	return NewTransformedShape(c, m)
}

func (c *Cylinder) Compile() {
}

func (c *Cylinder) BoundingBox() Box {
	r := c.Radius
	return Box{Vector{-r, -r, c.Z0}, Vector{r, r, c.Z1}}
}

func (shape *Cylinder) Intersect(ray Ray) Hit {
	r := shape.Radius
	o := ray.Origin
	d := ray.Direction
	a := d.X*d.X + d.Y*d.Y
	b := 2*o.X*d.X + 2*o.Y*d.Y
	c := o.X*o.X + o.Y*o.Y - r*r
	q := b*b - 4*a*c
	if q < EPS {
		return NoHit
	}
	s := math.Sqrt(q)
	t0 := (-b + s) / (2 * a)
	t1 := (-b - s) / (2 * a)
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	z0 := o.Z + t0*d.Z
	z1 := o.Z + t1*d.Z
	if t0 > EPS && shape.Z0 < z0 && z0 < shape.Z1 {
		return Hit{shape, t0, nil}
	}
	if t1 > EPS && shape.Z0 < z1 && z1 < shape.Z1 {
		return Hit{shape, t1, nil}
	}
	return NoHit

}

func (c *Cylinder) UV(p Vector) Vector {
	return Vector{}
}

func (c *Cylinder) MaterialAt(p Vector) Material {
	return c.Material
}

func (c *Cylinder) NormalAt(p Vector) Vector {
	p.Z = 0
	return p.Normalize()
}
