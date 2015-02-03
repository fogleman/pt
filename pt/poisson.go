package pt

import (
	"math"
	"math/rand"
)

type poissonGrid struct {
	r, size float64
	cells map[Vector]Vector
}

func newPoissonGrid(r float64) *poissonGrid {
	size := r / math.Sqrt(2)
	return &poissonGrid{r, size, make(map[Vector]Vector)}
}

func (grid *poissonGrid) normalize(v Vector) Vector {
	i := math.Floor(v.X / grid.size)
	j := math.Floor(v.Y / grid.size)
	return Vector{i, j, 0}
}

func (grid *poissonGrid) insert(v Vector) bool {
	n := grid.normalize(v)
	for i := n.X - 2; i < n.X + 3; i++ {
		for j := n.Y - 2; j < n.Y + 3; j++ {
			if m, ok := grid.cells[Vector{i, j, 0}]; ok {
				if math.Hypot(m.X - v.X, m.Y - v.Y) < grid.r {
					return false
				}
			}
		}
	}
	grid.cells[n] = v
	return true
}

func PoissonDisc(x1, y1, x2, y2, r float64, n int) []Vector {
	var result []Vector
	x := x1 + (x2 - x1) / 2
	y := y1 + (y2 - y1) / 2
	v := Vector{x, y, 0}
	active := []Vector{v}
	grid := newPoissonGrid(r)
	grid.insert(v)
	result = append(result, v)
	for len(active) > 0 {
		index := rand.Intn(len(active))
		point := active[index]
		ok := false
		for i := 0; i < n; i++ {
			a := rand.Float64() * 2 * math.Pi
			d := rand.Float64() * r + r
			x := point.X + math.Cos(a) * d
			y := point.Y + math.Sin(a) * d
			if x < x1 || y < y1 || x > x2 || y > y2 {
				continue
			}
			v := Vector{x, y, 0}
			if !grid.insert(v) {
				continue
			}
			result = append(result, v)
			active = append(active, v)
			ok = true
			break
		}
		if !ok {
			active = append(active[:index], active[index+1:]...)
		}
	}
	return result
}
