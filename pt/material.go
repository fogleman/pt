package pt

type Material struct {
	Color          Color
	Texture        Texture
	NormalTexture  Texture
	BumpTexture    Texture
	BumpMultiplier float64
	Emittance      float64
	Index          float64 // refractive index
	Gloss          float64 // reflection cone angle in radians
	Tint           float64 // specular and refractive tinting
	Transparent    bool
}

func DiffuseMaterial(color Color) Material {
	return Material{color, nil, nil, nil, 1, 0, 1, 0, 0, false}
}

func SpecularMaterial(color Color, index float64) Material {
	return Material{color, nil, nil, nil, 1, 0, index, 0, 0, false}
}

func GlossyMaterial(color Color, index, gloss float64) Material {
	return Material{color, nil, nil, nil, 1, 0, index, gloss, 0, false}
}

func ClearMaterial(index, gloss float64) Material {
	return Material{Color{}, nil, nil, nil, 1, 0, index, gloss, 0, true}
}

func TransparentMaterial(color Color, index, gloss, tint float64) Material {
	return Material{color, nil, nil, nil, 1, 0, index, gloss, tint, true}
}

func LightMaterial(color Color, emittance float64) Material {
	return Material{color, nil, nil, nil, 1, emittance, 1, 0, 0, false}
}
