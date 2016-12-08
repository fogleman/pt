package pt

type Hit struct {
	Shape   Shape
	T       float64
	HitInfo *HitInfo
}

type HitInfo struct {
	Shape    Shape
	Position Vector
	Normal   Vector
	Ray      Ray
	Material Material
	Inside   bool
}

var NoHit = Hit{nil, INF, nil}

func (hit *Hit) Ok() bool {
	return hit.T < INF
}

func (hit *Hit) Info(r Ray) HitInfo {
	if hit.HitInfo != nil {
		return *hit.HitInfo
	}
	shape := hit.Shape
	position := r.Position(hit.T)
	normal := shape.NormalAt(position)
	material := MaterialAt(shape, position)
	inside := false
	if normal.Dot(r.Direction) > 0 {
		normal = normal.Negate()
		inside = true
		switch shape.(type) {
		case *Volume, *SDFShape, *SphericalHarmonic:
			inside = false
		}
	}
	ray := Ray{position, normal}
	return HitInfo{shape, position, normal, ray, material, inside}
}
