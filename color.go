package pt

type Color struct {
	R, G, B float64
}

func (a Color) Add(b Color) Color {
	return Color{a.R + b.R, a.G + b.G, a.B + b.B}
}

func (a Color) Sub(b Color) Color {
	return Color{a.R - b.R, a.G - b.G, a.B - b.B}
}

func (a Color) Mul(b float64) Color {
	return Color{a.R * b, a.G * b, a.B * b}
}

func (a Color) MulColor(b Color) Color {
	return Color{a.R * b.R, a.G * b.G, a.B * b.B}
}

func (a Color) Div(b float64) Color {
	return Color{a.R / b, a.G / b, a.B / b}
}
