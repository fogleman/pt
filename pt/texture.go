package pt

import (
	"errors"
	"fmt"
	"image"
	"path"
	"strings"
)

type Texture interface {
	Sample(u, v float64) Color
}

type ImageTexture struct {
	image.Image
}

var textures map[string]Texture

func init() {
	textures = make(map[string]Texture)
}

func PNGTexture(path string) (Texture, error) {
	fmt.Printf("Loading PNG: %s\n", path)
	im, err := LoadPNG(path)
	if err != nil {
		return nil, err
	}
	return &ImageTexture{im}, nil
}

func JPGTexture(path string) (Texture, error) {
	fmt.Printf("Loading JPG: %s\n", path)
	im, err := LoadJPG(path)
	if err != nil {
		return nil, err
	}
	return &ImageTexture{im}, nil
}

func LoadTexture(p string) (Texture, error) {
	ext := strings.ToLower(path.Ext(p))
	switch ext {
	case ".png":
		return PNGTexture(p)
	case ".jpg":
		return JPGTexture(p)
	case ".jpeg":
		return JPGTexture(p)
	}
	err := errors.New(fmt.Sprintf("Unrecognized texture extension: %s", p))
	return nil, err
}

func GetTexture(path string) Texture {
	if texture, ok := textures[path]; ok {
		return texture
	}
	if texture, err := LoadTexture(path); err == nil {
		textures[path] = texture
		return texture
	}
	return nil
}

func (t *ImageTexture) Sample(u, v float64) Color {
	u = Fract(Fract(u) + 1)
	v = Fract(Fract(v) + 1)
	size := t.Image.Bounds().Max
	x := int(u * float64(size.X-1))
	y := int(v * float64(size.Y-1))
	r, g, b, _ := t.Image.At(x, y).RGBA()
	return Color{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535}.Pow(2.2)
}
