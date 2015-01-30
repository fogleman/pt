package pt

import (
	"math"
)

const INF = 1e9
const EPS = 1e-9

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
