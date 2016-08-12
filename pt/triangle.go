package pt

type Triangle struct {
	Material   *Material
	Box        Box
	V1, V2, V3 Vector
	N1, N2, N3 Vector
	T1, T2, T3 Vector
}

func NewTriangle(v1, v2, v3, t1, t2, t3 Vector, material Material) *Triangle {
	t := Triangle{}
	t.V1 = v1
	t.V2 = v2
	t.V3 = v3
	t.T1 = t1
	t.T2 = t2
	t.T3 = t3
	t.Material = &material
	t.UpdateBoundingBox()
	t.FixNormals()
	return &t
}

func (t *Triangle) Vertices() (Vector, Vector, Vector) {
	return t.V1, t.V2, t.V3
}

func (t *Triangle) Compile() {
}

func (t *Triangle) BoundingBox() Box {
	return t.Box
}

func (t *Triangle) Intersect(r Ray) Hit {
	e1x := t.V2.X - t.V1.X
	e1y := t.V2.Y - t.V1.Y
	e1z := t.V2.Z - t.V1.Z
	e2x := t.V3.X - t.V1.X
	e2y := t.V3.Y - t.V1.Y
	e2z := t.V3.Z - t.V1.Z
	px := r.Direction.Y*e2z - r.Direction.Z*e2y
	py := r.Direction.Z*e2x - r.Direction.X*e2z
	pz := r.Direction.X*e2y - r.Direction.Y*e2x
	det := e1x*px + e1y*py + e1z*pz
	if det > -EPS && det < EPS {
		return NoHit
	}
	inv := 1 / det
	tx := r.Origin.X - t.V1.X
	ty := r.Origin.Y - t.V1.Y
	tz := r.Origin.Z - t.V1.Z
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

func (t *Triangle) UV(p Vector) Vector {
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.T1.MulScalar(u))
	n = n.Add(t.T2.MulScalar(v))
	n = n.Add(t.T3.MulScalar(w))
	return Vector{n.X, n.Y, 0}
}

func (t *Triangle) MaterialAt(p Vector) Material {
	return *t.Material
}

func (t *Triangle) NormalAt(p Vector) Vector {
	u, v, w := t.Barycentric(p)
	n := Vector{}
	n = n.Add(t.N1.MulScalar(u))
	n = n.Add(t.N2.MulScalar(v))
	n = n.Add(t.N3.MulScalar(w))
	n = n.Normalize()
	if t.Material.NormalTexture != nil {
		b := Vector{}
		b = b.Add(t.T1.MulScalar(u))
		b = b.Add(t.T2.MulScalar(v))
		b = b.Add(t.T3.MulScalar(w))
		ns := t.Material.NormalTexture.NormalSample(b.X, b.Y)
		dv1 := t.V2.Sub(t.V1)
		dv2 := t.V3.Sub(t.V1)
		dt1 := t.T2.Sub(t.T1)
		dt2 := t.T3.Sub(t.T1)
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
	if t.Material.BumpTexture != nil {
		b := Vector{}
		b = b.Add(t.T1.MulScalar(u))
		b = b.Add(t.T2.MulScalar(v))
		b = b.Add(t.T3.MulScalar(w))
		bump := t.Material.BumpTexture.BumpSample(b.X, b.Y)
		dv1 := t.V2.Sub(t.V1)
		dv2 := t.V3.Sub(t.V1)
		dt1 := t.T2.Sub(t.T1)
		dt2 := t.T3.Sub(t.T1)
		tangent := dv1.MulScalar(dt2.Y).Sub(dv2.MulScalar(dt1.Y)).Normalize()
		bitangent := dv2.MulScalar(dt1.X).Sub(dv1.MulScalar(dt2.X)).Normalize()
		n = n.Add(tangent.MulScalar(bump.X * t.Material.BumpMultiplier))
		n = n.Add(bitangent.MulScalar(bump.Y * t.Material.BumpMultiplier))
	}
	n = n.Normalize()
	return n
}

func (t *Triangle) Area() float64 {
	e1 := t.V2.Sub(t.V1)
	e2 := t.V3.Sub(t.V1)
	n := e1.Cross(e2)
	return n.Length() / 2
}

func (t *Triangle) Barycentric(p Vector) (u, v, w float64) {
	v0 := t.V2.Sub(t.V1)
	v1 := t.V3.Sub(t.V1)
	v2 := p.Sub(t.V1)
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
	min := t.V1.Min(t.V2).Min(t.V3)
	max := t.V1.Max(t.V2).Max(t.V3)
	t.Box = Box{min, max}
}

func (t *Triangle) FixNormals() {
	e1 := t.V2.Sub(t.V1)
	e2 := t.V3.Sub(t.V1)
	n := e1.Cross(e2).Normalize()
	zero := Vector{}
	if t.N1 == zero {
		t.N1 = n
	}
	if t.N2 == zero {
		t.N2 = n
	}
	if t.N3 == zero {
		t.N3 = n
	}
}
