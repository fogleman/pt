package pt

type Triangle struct {
	material   *Material
	box        Box
	v1, v2, v3 Vector
	n1, n2, n3 Vector
	t1, t2, t3 Vector
}

func NewTriangle(v1, v2, v3, t1, t2, t3 Vector, material Material) *Triangle {
	t := Triangle{}
	t.v1 = v1
	t.v2 = v2
	t.v3 = v3
	t.t1 = t1
	t.t2 = t2
	t.t3 = t3
	t.material = &material
	t.UpdateBoundingBox()
	t.FixNormals()
	return &t
}

func (t *Triangle) Vertices() (Vector, Vector, Vector) {
	return t.v1, t.v2, t.v3
}

func (t *Triangle) Compile() {
}

func (t *Triangle) BoundingBox() Box {
	return t.box
}

func (t *Triangle) Intersect(r Ray) Hit {
	e1x := t.v2.X - t.v1.X
	e1y := t.v2.Y - t.v1.Y
	e1z := t.v2.Z - t.v1.Z
	e2x := t.v3.X - t.v1.X
	e2y := t.v3.Y - t.v1.Y
	e2z := t.v3.Z - t.v1.Z
	px := r.Direction.Y*e2z - r.Direction.Z*e2y
	py := r.Direction.Z*e2x - r.Direction.X*e2z
	pz := r.Direction.X*e2y - r.Direction.Y*e2x
	det := e1x*px + e1y*py + e1z*pz
	if det > -EPS && det < EPS {
		return NoHit
	}
	inv := 1 / det
	tx := r.Origin.X - t.v1.X
	ty := r.Origin.Y - t.v1.Y
	tz := r.Origin.Z - t.v1.Z
	u := (tx*px + ty*py + tz*pz) * inv
	if u < 0 || u > 1 {
		return NoHit
	}
	qx := ty*e1z - tz*e1y
	qy := tz*e1x - tx*e1z
	qz := tx*e1y - ty*e1x
	v := (r.Direction.X*qx + r.Direction.Y*qy + r.Direction.Z*qz) * inv
	if v < 0 || u+v > 1 {
		return NoHit
	}
	d := (e2x*qx + e2y*qy + e2z*qz) * inv
	if d < EPS {
		return NoHit
	}
	return Hit{t, d, nil}
}

func (t *Triangle) ColorAt(p Vector) Color {
	if t.material.Texture == nil {
		return t.material.Color
	}
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.t1.MulScalar(u))
	n = n.Add(t.t2.MulScalar(v))
	n = n.Add(t.t3.MulScalar(w))
	return t.material.Texture.Sample(n.X, n.Y)
}

func (t *Triangle) MaterialAt(p Vector) Material {
	return *t.material
}

func (t *Triangle) NormalAt(p Vector) Vector {
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.n1.MulScalar(u))
	n = n.Add(t.n2.MulScalar(v))
	n = n.Add(t.n3.MulScalar(w))
	n = n.Normalize()
	if t.material.NormalTexture != nil {
		b := Vector{}
		b = b.Add(t.t1.MulScalar(u))
		b = b.Add(t.t2.MulScalar(v))
		b = b.Add(t.t3.MulScalar(w))
		ns := t.material.NormalTexture.NormalSample(b.X, b.Y)
		dv1 := t.v2.Sub(t.v1)
		dv2 := t.v3.Sub(t.v1)
		dt1 := t.t2.Sub(t.t1)
		dt2 := t.t3.Sub(t.t1)
		T := dv1.MulScalar(dt2.Y).Sub(dv2.MulScalar(dt1.Y)).Normalize()
		B := dv2.MulScalar(dt1.X).Sub(dv1.MulScalar(dt2.X)).Normalize()
		N := T.Cross(B)
		matrix := Matrix{
			T.X, B.X, N.X, 0,
			T.Y, B.Y, N.Y, 0,
			T.Z, B.Z, N.Z, 0,
			0, 0, 0, 1}
		n = matrix.MulDirection(ns)
	}
	if t.material.BumpTexture != nil {
		b := Vector{}
		b = b.Add(t.t1.MulScalar(u))
		b = b.Add(t.t2.MulScalar(v))
		b = b.Add(t.t3.MulScalar(w))
		bump := t.material.BumpTexture.BumpSample(b.X, b.Y)
		dv1 := t.v2.Sub(t.v1)
		dv2 := t.v3.Sub(t.v1)
		dt1 := t.t2.Sub(t.t1)
		dt2 := t.t3.Sub(t.t1)
		tangent := dv1.MulScalar(dt2.Y).Sub(dv2.MulScalar(dt1.Y)).Normalize()
		bitangent := dv2.MulScalar(dt1.X).Sub(dv1.MulScalar(dt2.X)).Normalize()
		n = n.Add(tangent.MulScalar(bump.X * t.material.BumpMultiplier))
		n = n.Add(bitangent.MulScalar(bump.Y * t.material.BumpMultiplier))
	}
	n = n.Normalize()
	return n
}

func (t *Triangle) Area() float64 {
	e1 := t.v2.Sub(t.v1)
	e2 := t.v3.Sub(t.v1)
	n := e1.Cross(e2)
	return n.Length() / 2
}

func (t *Triangle) Barycentric(p Vector) (u, v, w float64) {
	v0 := t.v2.Sub(t.v1)
	v1 := t.v3.Sub(t.v1)
	v2 := p.Sub(t.v1)
	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)
	d := d00*d11 - d01*d01
	v = (d11*d20 - d01*d21) / d
	w = (d00*d21 - d01*d20) / d
	u = 1 - v - w
	return
}

func (t *Triangle) UpdateBoundingBox() {
	min := t.v1.Min(t.v2).Min(t.v3)
	max := t.v1.Max(t.v2).Max(t.v3)
	t.box = Box{min, max}
}

func (t *Triangle) FixNormals() {
	e1 := t.v2.Sub(t.v1)
	e2 := t.v3.Sub(t.v1)
	n := e1.Cross(e2).Normalize()
	zero := Vector{}
	if t.n1 == zero {
		t.n1 = n
	}
	if t.n2 == zero {
		t.n2 = n
	}
	if t.n3 == zero {
		t.n3 = n
	}
}
