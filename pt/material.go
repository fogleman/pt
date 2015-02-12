package pt

type Material struct {
	Color       Color
	Texture     Texture
	Emittance   float64
	Attenuation Attenuation
	Index       float64 // refractive index
	Gloss       float64 // reflection cone angle in radians
	Tint        float64 // specular tinting
	Transparent bool
}

func DiffuseMaterial(color Color) Material {
	return Material{color, nil, 0, NoAttenuation, 1, 0, 0, false}
}

func SpecularMaterial(color Color, index float64) Material {
	return Material{color, nil, 0, NoAttenuation, index, 0, 0, false}
}

func GlossyMaterial(color Color, index, gloss float64) Material {
	return Material{color, nil, 0, NoAttenuation, index, gloss, 0, false}
}

func ClearMaterial(index, gloss float64) Material {
	return Material{Color{}, nil, 0, NoAttenuation, index, gloss, 0, true}
}

func TransparentMaterial(color Color, index, gloss, tint float64) Material {
	return Material{color, nil, 0, NoAttenuation, index, gloss, tint, true}
}

func LightMaterial(color Color, emittance float64, attenuation Attenuation) Material {
	return Material{color, nil, emittance, attenuation, 1, 0, 0, false}
}
