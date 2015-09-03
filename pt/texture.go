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

var textures map[string]Texture

func init() {
	textures = make(map[string]Texture)
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

func PNGTexture(path string) (Texture, error) {
	fmt.Printf("Loading PNG: %s\n", path)
	im, err := LoadPNG(path)
	if err != nil {
		return nil, err
	}
	return NewTexture(im), nil
}

func JPGTexture(path string) (Texture, error) {
	fmt.Printf("Loading JPG: %s\n", path)
	im, err := LoadJPG(path)
	if err != nil {
		return nil, err
	}
	return NewTexture(im), nil
}

type ColorTexture struct {
	width, height int
	data          []Color
}

func NewTexture(im image.Image) Texture {
	size := im.Bounds().Max
	data := make([]Color, size.X*size.Y)
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			index := y*size.X + x
			r, g, b, _ := im.At(x, y).RGBA()
			fr := float64(r) / 65535
			fg := float64(g) / 65535
			fb := float64(b) / 65535
			data[index] = Color{fr, fg, fb}.Pow(2.2)
		}
	}
	return &ColorTexture{size.X, size.Y, data}
}

func (t *ColorTexture) Sample(u, v float64) Color {
	u = Fract(Fract(u) + 1)
	v = Fract(Fract(v) + 1)
	v = 1 - v
	x := int(u * float64(t.width))
	y := int(v * float64(t.height))
	return t.data[y*t.width+x]
}
