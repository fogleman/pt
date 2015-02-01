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

func NodeForShapes(shapes []Shape) *Node {
	box := BoxForShapes(shapes)
	return &Node{AxisNone, 0, box, shapes, nil, nil}
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
	for i := 0; i < len(xs) - 1; i++ {
		x := (xs[i] + xs[i + 1]) / 2
		l, r := node.PartitionCount(AxisX, x)
		n := int(math.Max(float64(l), float64(r)))
		if n < best {
			best = n
			bestAxis = AxisX
			bestPoint = x
		}
	}
	for i := 0; i < len(ys) - 1; i++ {
		y := (ys[i] + ys[i + 1]) / 2
		l, r := node.PartitionCount(AxisY, y)
		n := int(math.Max(float64(l), float64(r)))
		if n < best {
			best = n
			bestAxis = AxisY
			bestPoint = y
		}
	}
	for i := 0; i < len(zs) - 1; i++ {
		z := (zs[i] + zs[i + 1]) / 2
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
}
