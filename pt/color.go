package pt

import (
	"math"
)

type Color struct {
	R, G, B float64
}

func HexColor(x int) Color {
	r := float64((x>>16)&0xff) / 255
	g := float64((x>>8)&0xff) / 255
	b := float64((x>>0)&0xff) / 255
	return Color{r, g, b}.Pow(2.2)
}

func (a Color) Add(b Color) Color {
	return Color{a.R + b.R, a.G + b.G, a.B + b.B}
}

func (a Color) Sub(b Color) Color {
	return Color{a.R - b.R, a.G - b.G, a.B - b.B}
}

func (a Color) Mul(b Color) Color {
	return Color{a.R * b.R, a.G * b.G, a.B * b.B}
}

func (a Color) MulScalar(b float64) Color {
	return Color{a.R * b, a.G * b, a.B * b}
}

func (a Color) DivScalar(b float64) Color {
	return Color{a.R / b, a.G / b, a.B / b}
}

func (a Color) Min(b Color) Color {
	return Color{math.Min(a.R, b.R), math.Min(a.G, b.G), math.Min(a.B, b.B)}
}

func (a Color) Max(b Color) Color {
	return Color{math.Max(a.R, b.R), math.Max(a.G, b.G), math.Max(a.B, b.B)}
}

func (a Color) Pow(b float64) Color {
	return Color{math.Pow(a.R, b), math.Pow(a.G, b), math.Pow(a.B, b)}
}

func (a Color) Mix(b Color, pct float64) Color {
	a = a.MulScalar(1 - pct)
	b = b.MulScalar(pct)
	return a.Add(b)
}
