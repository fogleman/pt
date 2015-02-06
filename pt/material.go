package pt

type Material struct {
	Color Color
	Index float64 // refractive index
	Gloss float64 // reflection cone angle in radians
	Tint  float64 // specular tinting
}

func DiffuseMaterial(color Color) Material {
	return Material{color, 1, 0, 0}
}

func RefractiveMaterial(color Color, index float64) Material {
	return Material{color, index, 0, 0}
}
