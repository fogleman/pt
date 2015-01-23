package pt

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.X * a.X + a.Y * a.Y + a.Z * a.Z)
}

func (a Vector) Dot(b Vector) float64 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z;
}

func (a Vector) Cross(b Vector) Vector {
	x := a.Y * b.Z - a.Z * b.Y;
	y := a.Z * b.X - a.X * b.Z;
	z := a.X * b.Y - a.Y * b.X;
	return Vector{x, y, z}
}

func (a Vector) Normalize() Vector {
	d := a.Length()
	return Vector{a.X / d, a.Y / d, a.Z / d}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Mul(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

func (a Vector) Div(b float64) Vector {
	return Vector{a.X / b, a.Y / b, a.Z / b}
}
