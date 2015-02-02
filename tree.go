package pt

import (
	"math"
)

type Node struct {
	axis   Axis
	point  float64
	box    Box
	shapes []Shape
	left   *Node
	right  *Node
}

func NewTree(shapes []Shape) *Node {
	node := NodeForShapes(shapes)
	node.Split(0)
	return node
}

func NodeForShapes(shapes []Shape) *Node {
	box := BoxForShapes(shapes)
	return &Node{AxisNone, 0, box, shapes, nil, nil}
}

func (node *Node) Intersect(r Ray) (hit Hit, ok bool) {
	tmin, tmax := node.box.Intersect(r)
	if tmax < tmin || tmax < 0 {
		return
	}
	hit = node.RecursiveIntersect(r, tmin, tmax)
	ok = hit.T < INF
	return
}

func (node *Node) RecursiveIntersect(r Ray, tmin, tmax float64) (hit Hit) {
	if node.axis == AxisNone {
		return node.IntersectShapes(r)
	}
	var tsplit float64
	var swap bool
	switch node.axis {
	case AxisX:
		tsplit = (node.point - r.Origin.X) / r.Direction.X
		swap = r.Origin.X > node.point
	case AxisY:
		tsplit = (node.point - r.Origin.Y) / r.Direction.Y
		swap = r.Origin.Y > node.point
	case AxisZ:
		tsplit = (node.point - r.Origin.Z) / r.Direction.Z
		swap = r.Origin.Z > node.point
	}
	var left, right bool
	if tsplit < tmin {
		right = true
	} else if tsplit > tmax {
		left = true
	} else {
		right = true
		left = true
	}
	if swap {
		left, right = right, left
	}
	// tsplit tmin tmax - right
	// tmin tsplit tmax - both
	// tmin tmax split - left
	// TODO: only check children if needed and in appropriate order
	h1, h2 := Hit{}, Hit{}
	h1.T = INF
	h2.T = INF
	if left {
		h1 = node.left.RecursiveIntersect(r, tmin, tmax)
	}
	if right {
		h2 = node.right.RecursiveIntersect(r, tmin, tmax)
	}
	if h1.T < h2.T {
		return h1
	} else {
		return h2
	}
}

func (node *Node) IntersectShapes(r Ray) (hit Hit) {
	hit.T = INF
	for _, shape := range node.shapes {
		t := shape.Intersect(r)
		if t < hit.T {
			p := r.Position(t)
			n := shape.Normal(p)
			hit = Hit{shape, Ray{p, n}, t}
		}
	}
	return
}

func (node *Node) PartitionCount(axis Axis, point float64) (left, right int) {
	for _, shape := range node.shapes {
		box := shape.Box()
		l, r := box.Partition(axis, point)
		if l {
			left++
		}
		if r {
			right++
		}
	}
	return
}

func (node *Node) Partition(axis Axis, point float64) (left, right []Shape) {
	for _, shape := range node.shapes {
		box := shape.Box()
		l, r := box.Partition(axis, point)
		if l {
			left = append(left, shape)
		}
		if r {
			right = append(right, shape)
		}
	}
	return
}

func (node *Node) Split(depth int) {
	// TODO: max depth?
	if len(node.shapes) < 10 {
		return
	}
	var xs, ys, zs []float64
	for _, shape := range node.shapes {
		box := shape.Box()
		xs = append(xs, box.Min.X)
		xs = append(xs, box.Max.X)
		ys = append(ys, box.Min.Y)
		ys = append(ys, box.Max.Y)
		zs = append(zs, box.Min.Z)
		zs = append(zs, box.Max.Z)
	}
	xs = Distinct(xs)
	ys = Distinct(ys)
	zs = Distinct(zs)
	best := len(node.shapes) + 1
	bestAxis := AxisNone
	bestPoint := 0.0
	for i := 0; i < len(xs)-1; i++ {
		x := (xs[i] + xs[i+1]) / 2
		l, r := node.PartitionCount(AxisX, x)
		n := int(math.Max(float64(l), float64(r)))
		if n < best {
			best = n
			bestAxis = AxisX
			bestPoint = x
		}
	}
	for i := 0; i < len(ys)-1; i++ {
		y := (ys[i] + ys[i+1]) / 2
		l, r := node.PartitionCount(AxisY, y)
		n := int(math.Max(float64(l), float64(r)))
		if n < best {
			best = n
			bestAxis = AxisY
			bestPoint = y
		}
	}
	for i := 0; i < len(zs)-1; i++ {
		z := (zs[i] + zs[i+1]) / 2
		l, r := node.PartitionCount(AxisZ, z)
		n := int(math.Max(float64(l), float64(r)))
		if n < best {
			best = n
			bestAxis = AxisZ
			bestPoint = z
		}
	}
	// TODO: check pct?
	// pct := float64(best) / float64(len(node.shapes))
	l, r := node.Partition(bestAxis, bestPoint)
	node.left = NodeForShapes(l)
	node.right = NodeForShapes(r)
	node.left.Split(depth + 1)
	node.right.Split(depth + 1)
	// node.shapes = nil
}
