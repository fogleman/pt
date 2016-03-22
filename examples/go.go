package main

import (
	"math"
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

func offset(stdev float64) Vector {
	a := rand.Float64() * 2 * math.Pi
	r := rand.NormFloat64() * stdev
	x := math.Cos(a) * r
	y := math.Sin(a) * r
	return Vector{x, 0, y}
}

func intersects(scene *Scene, shape Shape) bool {
	box := shape.Box()
	for _, other := range scene.Shapes() {
		if box.Intersects(other.Box()) {
			return true
		}
	}
	return false
}

func main() {
	scene := Scene{}
	black := GlossyMaterial(HexColor(0x111111), 1.5, Radians(45))
	white := GlossyMaterial(HexColor(0xFFFFFF), 1.6, Radians(20))
	for _, p := range blackPositions {
		for {
			m := Scale(Vector{0.48, 0.2, 0.48}).Translate(Vector{p[0] - 9.5, 0, p[1] - 9.5})
			m = m.Translate(offset(0.02))
			shape := NewTransformedShape(NewSphere(Vector{}, 1, black), m)
			if intersects(&scene, shape) {
				continue
			}
			scene.Add(shape)
			break
		}
	}
	for _, p := range whitePositions {
		for {
			m := Scale(Vector{0.48, 0.2, 0.48}).Translate(Vector{p[0] - 9.5, 0, p[1] - 9.5})
			m = m.Translate(offset(0.02))
			shape := NewTransformedShape(NewSphere(Vector{}, 1, white), m)
			if intersects(&scene, shape) {
				continue
			}
			scene.Add(shape)
			break
		}
	}
	for i := 0; i < 19; i++ {
		x := float64(i) - 9.5
		m := 0.015
		scene.Add(NewCube(Vector{x - m, -1, -9.5}, Vector{x + m, -0.195, 8.5}, black))
		scene.Add(NewCube(Vector{-9.5, -1, x - m}, Vector{8.5, -0.195, x + m}, black))
	}
	material := GlossyMaterial(HexColor(0xEFECCA), 1.2, Radians(30))
	material.Texture = GetTexture("examples/wood.jpg", 2.2)
	scene.Add(NewCube(Vector{-12, -12, -12}, Vector{12, -0.2, 12}, material))
	// texture, err := LoadTexture("examples/river_rocks_ccyby/river_rocks_8k.png")
	scene.SetTexture(GetTexture("examples/courtyard_ccby/courtyard_8k.png", 1))
	camera := LookAt(Vector{-0.5, 5, 5}, Vector{-0.5, 0, 0.5}, Vector{0, 1, 0}, 50)
	IterativeRender("out%03d.png", 10000, &scene, &camera, 2560, 1440, -1, 16, 4)
}

var blackPositions = [][]float64{
	{7, 3}, {14, 17}, {14, 4}, {18, 4}, {0, 7}, {5, 8}, {11, 5}, {10, 7}, {7, 6}, {6, 10}, {12, 6}, {3, 2}, {5, 11}, {7, 5}, {14, 15}, {12, 11}, {8, 12}, {4, 15}, {2, 11}, {9, 9}, {10, 3}, {6, 17}, {7, 2}, {14, 5}, {13, 3}, {13, 16}, {3, 6}, {1, 10}, {4, 1}, {10, 9}, {5, 17}, {12, 7}, {3, 5}, {2, 7}, {5, 10}, {10, 10}, {5, 7}, {7, 4}, {12, 4}, {8, 13}, {9, 8}, {15, 17}, {3, 10}, {4, 13}, {2, 13}, {8, 16}, {12, 3}, {17, 5}, {13, 2}, {15, 3}, {2, 3}, {6, 5}, {11, 7}, {16, 5}, {11, 8}, {14, 7}, {15, 6}, {1, 7}, {5, 9}, {10, 11}, {6, 6}, {4, 18}, {7, 14}, {17, 3}, {4, 9}, {10, 12}, {6, 3}, {16, 7}, {14, 14}, {16, 18}, {3, 13}, {1, 13}, {2, 10}, {7, 9}, {13, 1}, {12, 15}, {4, 3}, {5, 2}, {10, 2},
}

var whitePositions = [][]float64{
	{16, 6}, {16, 9}, {13, 4}, {1, 6}, {0, 10}, {3, 7}, {1, 11}, {8, 5}, {6, 7}, {5, 5}, {15, 11}, {13, 7}, {18, 9}, {2, 6}, {7, 10}, {15, 14}, {13, 10}, {17, 18}, {7, 15}, {5, 14}, {3, 18}, {15, 16}, {14, 8}, {12, 8}, {7, 13}, {1, 15}, {8, 9}, {6, 14}, {12, 2}, {17, 6}, {18, 5}, {17, 11}, {9, 7}, {6, 4}, {5, 4}, {6, 11}, {11, 9}, {13, 6}, {18, 6}, {0, 8}, {8, 3}, {4, 6}, {9, 2}, {4, 17}, {14, 12}, {13, 9}, {18, 11}, {3, 15}, {4, 8}, {2, 8}, {12, 9}, {16, 17}, {8, 10}, {9, 11}, {17, 7}, {16, 11}, {14, 10}, {3, 9}, {1, 9}, {8, 7}, {2, 14}, {9, 6}, {5, 3}, {14, 16}, {5, 16}, {16, 8}, {13, 5}, {8, 4}, {4, 7}, {5, 6}, {11, 2}, {12, 5}, {15, 8}, {2, 9}, {9, 15}, {8, 1}, {4, 4}, {16, 15}, {12, 10}, {13, 11}, {2, 16}, {4, 14}, {5, 15}, {10, 1}, {6, 8}, {6, 12}, {17, 9}, {8, 8},
}
