package pt

import (
	"image"
	"math"
	"math/rand"
)

type Volume struct {
	W, H, D  int
	Data     []float64
	Lo, Hi   float64
	material Material
	box      Box
}

func NewVolume(images []image.Image, lo, hi float64, material Material) *Volume {
	w := images[0].Bounds().Size().X
	h := images[0].Bounds().Size().Y
	d := len(images)
	data := make([]float64, w*h*d)
	for z, im := range images {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, _, _, _ := im.At(x, y).RGBA()
				f := float64(r) / 65535
				data[x+y*w+z*w*h] = f
			}
		}
	}
	box := Box{Vector{-1, -1, -1}, Vector{0, 1, 1}}
	return &Volume{w, h, d, data, lo, hi, material, box}
}

func (v *Volume) Get(x, y, z int) float64 {
	if x < 0 || y < 0 || z < 0 || x >= v.W || y >= v.H || z >= v.D {
		return 0
	}
	return v.Data[x+y*v.W+z*v.W*v.H]
}

func (v *Volume) Sample(x, y, z float64) float64 {
	z /= 0.625
	x = ((x + 1) / 2) * float64(v.W)
	y = ((y + 1) / 2) * float64(v.H)
	z = ((z + 1) / 2) * float64(v.D)
	x0 := int(math.Floor(x))
	y0 := int(math.Floor(y))
	z0 := int(math.Floor(z))
	x1 := x0 + 1
	y1 := y0 + 1
	z1 := z0 + 1
	v000 := v.Get(x0, y0, z0)
	v001 := v.Get(x0, y0, z1)
	v010 := v.Get(x0, y1, z0)
	v011 := v.Get(x0, y1, z1)
	v100 := v.Get(x1, y0, z0)
	v101 := v.Get(x1, y0, z1)
	v110 := v.Get(x1, y1, z0)
	v111 := v.Get(x1, y1, z1)
	x -= float64(x0)
	y -= float64(y0)
	z -= float64(z0)
	c00 := v000*(1-x) + v100*x
	c01 := v001*(1-x) + v101*x
	c10 := v010*(1-x) + v110*x
	c11 := v011*(1-x) + v111*x
	c0 := c00*(1-y) + c10*y
	c1 := c01*(1-y) + c11*y
	c := c0*(1-z) + c1*z
	return c
}

func (v *Volume) Compile() {
}

func (v *Volume) Box() Box {
	return v.box
}

func (v *Volume) Contains(a Vector) bool {
	if !v.box.Contains(a) {
		return false
	}
	s := v.Sample(a.X, a.Y, a.Z)
	return s >= v.Lo && s <= v.Hi
}

func (v *Volume) Intersect(ray Ray) Hit {
	start := math.Max(0, ray.Origin.Length()-1)
	step := 1.0 / 1024
	for t := start; t < start+2; t += step {
		if v.Contains(ray.Position(t)) {
			t -= step
			step /= 64
			t += step
			for i := 0; i < 64; i++ {
				if v.Contains(ray.Position(t)) {
					return Hit{v, t - step, nil}
				}
				t += step
			}
		}
	}
	return NoHit
}

func (v *Volume) Color(p Vector) Color {
	return v.material.Color
}

func (v *Volume) Material(p Vector) Material {
	return v.material
}

func (v *Volume) Normal(p Vector) Vector {
	eps := 0.01
	n := Vector{
		v.Sample(p.X-eps, p.Y, p.Z) - v.Sample(p.X+eps, p.Y, p.Z),
		v.Sample(p.X, p.Y-eps, p.Z) - v.Sample(p.X, p.Y+eps, p.Z),
		v.Sample(p.X, p.Y, p.Z-eps) - v.Sample(p.X, p.Y, p.Z+eps),
	}
	return n.Normalize()
}

func (v *Volume) RandomPoint(rnd *rand.Rand) Vector {
	return Vector{}
}
