package pt

import (
	"image"
)

type Texture interface {
	Sample(u, v float64) Color
}

type ImageTexture struct {
	image.Image
}

func NewTexture(im image.Image) Texture {
	return &ImageTexture{im}
}

func PNGTexture(path string) (Texture, error) {
	im, err := LoadPNG(path)
	if err != nil {
		return nil, err
	}
	return NewTexture(im), nil
}

func JPGTexture(path string) (Texture, error) {
	im, err := LoadJPG(path)
	if err != nil {
		return nil, err
	}
	return NewTexture(im), nil
}

func (t *ImageTexture) Sample(u, v float64) Color {
	u = Fract(Fract(u) + 1)
	v = Fract(Fract(v) + 1)
	size := t.Image.Bounds().Max
	x := int(u * float64(size.X-1))
	y := int(v * float64(size.Y-1))
	r, g, b, _ := t.Image.At(x, y).RGBA()
	return Color{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535}
}
