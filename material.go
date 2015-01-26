package pt

type Material struct {
	Gloss, Cone float64
}

func DiffuseMaterial() Material {
	return Material{0, 0}
}

func ReflectiveMaterial() Material {
	return Material{1, 0}
}
