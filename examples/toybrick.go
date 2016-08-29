package main

import (
	"math/rand"

	. "github.com/fogleman/pt/pt"
)

const H = 1.46875

func CreateBrick(color int) *Mesh {
	material := GlossyMaterial(HexColor(Colors[color]), 1.3, Radians(20))
	mesh, err := LoadSTL("examples/toybrick.stl", material)
	if err != nil {
		panic(err)
	}
	mesh.SmoothNormalsThreshold(Radians(20))
	mesh.FitInside(Box{Vector{}, Vector{2, 4, 10}}, Vector{0, 0, 0})
	return mesh
}

func main() {
	scene := Scene{}
	scene.Color = White
	meshes := []*Mesh{
		CreateBrick(1),  // white
		CreateBrick(21), // bright red
		CreateBrick(23), // bright blue
		CreateBrick(24), // bright yellow
		CreateBrick(26), // black
		CreateBrick(28), // dark green
	}
	for x := -30; x <= 50; x += 2 {
		for y := -50; y <= 20; y += 4 {
			h := rand.Intn(5) + 1
			for i := 0; i < h; i++ {
				dy := 0
				if (x/2+i)%2 == 0 {
					dy = 2
				}
				z := float64(i) * H
				mesh := meshes[rand.Intn(len(meshes))]
				m := Translate(Vector{float64(x), float64(y + dy), z})
				scene.Add(NewTransformedShape(mesh, m))
			}
		}
	}
	camera := LookAt(Vector{-23, 13, 20}, Vector{0, 0, 0}, Vector{0, 0, 1}, 45)
	sampler := NewSampler(4, 4)
	IterativeRender("out%03d.png", 1000, &scene, &camera, sampler, 1920/2, 1080/2, -1)
}

var Colors = map[int]int{
	1:   0xF2F3F2,
	2:   0xA1A5A2,
	3:   0xF9E999,
	5:   0xD7C599,
	6:   0xC2DAB8,
	9:   0xE8BAC7,
	12:  0xCB8442,
	18:  0xCC8E68,
	21:  0xC4281B,
	22:  0xC470A0,
	23:  0x0D69AB,
	24:  0xF5CD2F,
	25:  0x624732,
	26:  0x1B2A34,
	27:  0x6D6E6C,
	28:  0x287F46,
	29:  0xA1C48B,
	36:  0xF3CF9B,
	37:  0x4B974A,
	38:  0xA05F34,
	39:  0xC1CADE,
	45:  0xB4D2E3,
	100: 0xEEC4B6,
	101: 0xDA8679,
	102: 0x6E99C9,
	103: 0xC7C1B7,
	104: 0x6B327B,
	105: 0xE29B3F,
	106: 0xDA8540,
	107: 0x008F9B,
	108: 0x685C43,
	110: 0x435493,
	112: 0x6874AC,
	115: 0xC7D23C,
	116: 0x55A5AF,
	118: 0xB7D7D5,
	119: 0xA4BD46,
	120: 0xD9E4A7,
	121: 0xE7AC58,
	123: 0xD36F4C,
	124: 0x923978,
	125: 0xEAB891,
	127: 0xDCBC81,
	128: 0xAE7A59,
	131: 0x9CA3A8,
	135: 0x74869C,
	136: 0x877C90,
	137: 0xE09864,
	138: 0x958A73,
	140: 0x203A56,
	141: 0x27462C,
	145: 0x7988A1,
	146: 0x958EA3,
	147: 0x938767,
	148: 0x575857,
	149: 0x161D32,
	150: 0xABADAC,
	151: 0x789081,
	153: 0x957976,
	154: 0x7B2E2F,
	168: 0x756C62,
	180: 0xD7A94B,
	200: 0x828A5D,
	190: 0xF9D62E,
	191: 0xE8AB2D,
	192: 0x694027,
	193: 0xCF6024,
	194: 0xA3A2A4,
	195: 0x4667A4,
	196: 0x23478B,
	198: 0x8E4285,
	199: 0x635F61,
	208: 0xE5E4DE,
	209: 0xB08E44,
	210: 0x709578,
	211: 0x79B5B5,
	212: 0x9FC3E9,
	213: 0x6C81B7,
	216: 0x8F4C2A,
	217: 0x7C5C45,
	218: 0x96709F,
	219: 0x6B629B,
	220: 0xA7A9CE,
	221: 0xCD6298,
	222: 0xE4ADC8,
	223: 0xDC9095,
	224: 0xF0D5A0,
	225: 0xEBB87F,
	226: 0xFDEA8C,
	232: 0x7DBBDD,
	268: 0x342B75,
}
