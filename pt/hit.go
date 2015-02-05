package pt

type Hit struct {
	Shape Shape
	T     float64
}

var NoHit = Hit{T: INF}

func (hit *Hit) Ok() bool {
	return hit.T < INF
}

func (hit *Hit) Info(r Ray) Info {
	shape := hit.Shape
	position := r.Position(hit.T)
	normal := shape.Normal(position)
	ray := Ray{position, normal}
	color := shape.Color(position)
	material := shape.Material(position)
	return Info{shape, position, normal, ray, color, material}
}

type Info struct {
	Shape    Shape
	Position Vector
	Normal   Vector
	Ray      Ray
	Color    Color
	Material Material
}
