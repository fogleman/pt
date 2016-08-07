package pt

import (
	"image"
	"math"
	"math/rand"
)

type Volume struct {
	W, H, D     int
	ZScale      float64
	Data        []float64
	Windows     []VolumeWindow
	BoundingBox Box
}

type VolumeWindow struct {
	Lo, Hi   float64
	Material Material
}

func NewVolume(box Box, images []image.Image, sliceSpacing float64, windows []VolumeWindow) *Volume {
	w := images[0].Bounds().Size().X
	h := images[0].Bounds().Size().Y
	d := len(images)
	// TODO: w/h aspect ratio
	zs := (sliceSpacing * float64(d)) / float64(w)
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
	return &Volume{w, h, d, zs, data, windows, box}
}

func (v *Volume) Get(x, y, z int) float64 {
	if x < 0 || y < 0 || z < 0 || x >= v.W || y >= v.H || z >= v.D {
		return 0
	}
	return v.Data[x+y*v.W+z*v.W*v.H]
}

func (v *Volume) Sample(x, y, z float64) float64 {
	z /= v.ZScale
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
	return v.BoundingBox
}

func (v *Volume) Sign(a Vector) int {
	s := v.Sample(a.X, a.Y, a.Z)
	for i, window := range v.Windows {
		if s < window.Lo {
			return i + 1
		}
		if s > window.Hi {
			continue
		}
		return 0
	}
	return len(v.Windows) + 1
}

func (v *Volume) Intersect(ray Ray) Hit {
	tmin, tmax := v.BoundingBox.Intersect(ray)
	step := 1.0 / 512
	start := math.Max(step, tmin)
	sign := -1
	for t := start; t <= tmax; t += step {
		p := ray.Position(t)
		s := v.Sign(p)
		if s == 0 || (sign >= 0 && s != sign) {
			t -= step
			step /= 64
			t += step
			for i := 0; i < 64; i++ {
				if v.Sign(ray.Position(t)) == 0 {
					return Hit{v, t - step, nil}
				}
				t += step
			}
		}
		sign = s
	}
	return NoHit
}

func (v *Volume) Color(p Vector) Color {
	return v.Material(p).Color
}

func (v *Volume) Material(p Vector) Material {
	be := 1e9
	bm := Material{}
	s := v.Sample(p.X, p.Y, p.Z)
	for _, window := range v.Windows {
		if s >= window.Lo && s <= window.Hi {
			return window.Material
		}
		e := math.Min(math.Abs(s-window.Lo), math.Abs(s-window.Hi))
		if e < be {
			be = e
			bm = window.Material
		}
	}
	return bm

}

func (v *Volume) Normal(p Vector) Vector {
	eps := 0.001
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
