package pt

type Material struct {
	Color   Color
	Texture Texture
	Index   float64 // refractive index
	Gloss   float64 // reflection cone angle in radians
	Tint    float64 // specular tinting
}

func DiffuseMaterial(color Color) Material {
	return Material{color, nil, 1, 0, 0}
}

func SpecularMaterial(color Color, index float64) Material {
	return Material{color, nil, index, 0, 0}
}

func GlossyMaterial(color Color, index, gloss float64) Material {
	return Material{color, nil, index, gloss, 0}
}
