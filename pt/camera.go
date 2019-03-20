package pt

import (
	"math"
	"math/rand"
)

type Camera struct {
	p, u, v, w     Vector
	m              float64
	focalDistance  float64
	apertureRadius float64
}

func LookAt(eye, center, up Vector, fovy float64) Camera {
	c := Camera{}
	c.p = eye
	c.w = center.Sub(eye).Normalize()
	c.u = up.Cross(c.w).Normalize()
	c.v = c.w.Cross(c.u).Normalize()
	c.m = 1 / math.Tan(fovy*math.Pi/360)
	return c
}

func (c *Camera) SetFocus(focalPoint Vector, apertureRadius float64) {
	c.focalDistance = focalPoint.Sub(c.p).Length()
	c.apertureRadius = apertureRadius
}

func (c *Camera) CastRay(x, y, w, h int, u, v float64, rnd *rand.Rand) Ray {
	aspect := float64(w) / float64(h)
	px := ((float64(x)+u-0.5)/(float64(w)-1))*2 - 1
	py := ((float64(y)+v-0.5)/(float64(h)-1))*2 - 1
	d := Vector{}
	d = d.Add(c.u.MulScalar(-px * aspect))
	d = d.Add(c.v.MulScalar(-py))
	d = d.Add(c.w.MulScalar(c.m))
	d = d.Normalize()
	p := c.p
	if c.apertureRadius > 0 {
		focalPoint := c.p.Add(d.MulScalar(c.focalDistance))
		angle := rnd.Float64() * 2 * math.Pi
		radius := rnd.Float64() * c.apertureRadius
		p = p.Add(c.u.MulScalar(math.Cos(angle) * radius))
		p = p.Add(c.v.MulScalar(math.Sin(angle) * radius))
		d = focalPoint.Sub(p).Normalize()
	}
	return Ray{p, d}
}

// OrthogonalCamera implements a simple orthogonal camera
type OrthogonalCamera struct {
	up, right, pos, dir Vector
	width               float64
}

// OrthoLookAt sets up a new camera with the center of the camera at location
// its center points at the given target (i.e. the target point is centered in the image)
// up defines the up direction for the camera, width defines the width of the image
func OrthoLookAt(location, up, target Vector, width float64) OrthogonalCamera {
	oc := OrthogonalCamera{
		up:    up.Normalize(),
		pos:   location,
		dir:   target.Sub(location).Normalize(),
		width: width,
	}
	oc.right = oc.dir.Cross(oc.up).Normalize()
	oc.up = oc.dir.Cross(oc.right).Normalize()
	return oc
}

// CalculateOrthoSize computes width and height that would be needed by the camera to record an image that covers
// all given points in front of the camera. points behind the camera are ignored
// width can either be adjusted to fit the desired aspect ratio or the aspect ratio can be choosen to fit the size
func CalculateOrthoSize(location, up, target Vector, points ...Vector) (w, h float64) {
	// create camera to have accurate precomputed values
	oc := OrthoLookAt(location, up, target, 1.0)
	// max values
	dir := oc.dir.MulScalar(-1)
	for _, p := range points {
		t := oc.pos.Sub(p).Dot(dir)
		if t < 0 {
			// we can not work with points behind the plane
			continue
		}
		// compute point in plane
		pp := p.Add(dir.MulScalar(t))

		// we are only interested in the absolute since we keep the image symetric
		u := math.Abs(oc.up.Dot(pp))
		v := math.Abs(oc.right.Dot(pp))
		w = math.Max(w, v)
		h = math.Max(h, u)
	}

	// since we calculated distance from the center, this is both
	return w * 2, h * 2

}

// CastRay implements the RenderCamera interface and creates the ray used for rendering
func (oc *OrthogonalCamera) CastRay(x, y, w, h int, u, v float64, rnd *rand.Rand) Ray {

	wf, xf := float64(w), float64(x)
	hf, yf := float64(h), float64(y)
	size := oc.width / wf
	right := oc.right.MulScalar(size * (xf - wf/2 + u))
	up := oc.up.MulScalar(size * (yf - hf/2 + v))

	return Ray{
		Origin:    oc.pos.Add(up).Add(right),
		Direction: oc.dir,
	}
}
