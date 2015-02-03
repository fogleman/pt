package pt

type Material struct {
	Index float64 // refractive index
	Gloss float64 // reflection cone angle in radians
	Tint  float64 // specular tinting
}

func DiffuseMaterial() Material {
	return Material{1, 0, 0}
}

func RefractiveMaterial(index float64) Material {
	return Material{index, 0, 0}
}
