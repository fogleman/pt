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
	NormalSample(u, v float64) Vector
	BumpSample(u, v float64) Vector
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
			data[index] = NewColor(im.At(x, y)).Pow(2.2)
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

func (t *ColorTexture) NormalSample(u, v float64) Vector {
	c := t.Sample(u, v).Pow(1 / 2.2)
	return Vector{c.R*2 - 1, c.G*2 - 1, c.B*2 - 1}.Normalize()
}

func (t *ColorTexture) BumpSample(u, v float64) Vector {
	u = Fract(Fract(u) + 1)
	v = Fract(Fract(v) + 1)
	v = 1 - v
	x := int(u * float64(t.width))
	y := int(v * float64(t.height))
	x1, x2 := ClampInt(x-1, 0, t.width-1), ClampInt(x+1, 0, t.width-1)
	y1, y2 := ClampInt(y-1, 0, t.height-1), ClampInt(y+1, 0, t.height-1)
	cx := t.data[y*t.width+x1].Sub(t.data[y*t.width+x2])
	cy := t.data[y1*t.width+x].Sub(t.data[y2*t.width+x])
	return Vector{cx.R, cy.R, 0}
}
