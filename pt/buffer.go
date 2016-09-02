package pt

import "image"

type Pixel struct {
	Samples int
	C, M, V Color
}

func (p *Pixel) AddSample(sample Color) {
	p.Samples++
	p.C = p.C.Add(sample)
	if p.Samples == 1 {
		p.M = sample
		return
	}
	m := p.M
	p.M = p.M.Add(sample.Sub(p.M).DivScalar(float64(p.Samples)))
	p.V = p.V.Add(sample.Sub(m).Mul(sample.Sub(p.M)))
}

func (p *Pixel) Color() Color {
	return p.C.DivScalar(float64(p.Samples))
}

func (p *Pixel) Variance() Color {
	if p.Samples < 2 {
		return Black
	}
	return p.V.DivScalar(float64(p.Samples - 1)).Pow(0.5)
}

type Buffer struct {
	W      int
	H      int
	Pixels []Pixel
}

func NewBuffer(w, h int) *Buffer {
	pixels := make([]Pixel, w*h)
	return &Buffer{w, h, pixels}
}

func (b *Buffer) Copy() *Buffer {
	pixels := make([]Pixel, b.W*b.H)
	copy(pixels, b.Pixels)
	return &Buffer{b.W, b.H, pixels}
}

func (b *Buffer) AddSample(x, y int, sample Color) {
	b.Pixels[y*b.W+x].AddSample(sample)
}

func (b *Buffer) Samples(x, y int) int {
	return b.Pixels[y*b.W+x].Samples
}

func (b *Buffer) Color(x, y int) Color {
	return b.Pixels[y*b.W+x].Color()
}

func (b *Buffer) Variance(x, y int) Color {
	return b.Pixels[y*b.W+x].Variance()
}

func (b *Buffer) Image() image.Image {
	result := image.NewRGBA64(image.Rect(0, 0, b.W, b.H))
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			c := b.Pixels[y*b.W+x].Color().Pow(1 / 2.2)
			result.SetRGBA64(x, y, c.RGBA64())
		}
	}
	return result
}

func (b *Buffer) VarianceImage() image.Image {
	result := image.NewRGBA64(image.Rect(0, 0, b.W, b.H))
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			c := b.Pixels[y*b.W+x].Variance().MaxComponent()
			result.SetRGBA64(x, y, Color{c, c, c}.RGBA64())
		}
	}
	return result
}
