package pt

import (
	"fmt"
	"math"
)

type Tree struct {
	box  Box
	root *Node
}

func NewTree(shapes []Shape) *Tree {
	fmt.Print("Building k-d tree... ")
	defer fmt.Println("OK")
	box := BoxForShapes(shapes)
	node := NewNode(shapes)
	node.Split(0)
	return &Tree{box, node}
}

func (tree *Tree) Intersect(r Ray) Hit {
	tmin, tmax := tree.box.Intersect(r)
	if tmax < tmin || tmax <= 0 {
		return NoHit
	}
	return tree.root.Intersect(r, tmin, tmax)
}

type Node struct {
	axis   Axis
	point  float64
	shapes []Shape
	left   *Node
	right  *Node
}

func NewNode(shapes []Shape) *Node {
	return &Node{AxisNone, 0, shapes, nil, nil}
}

func (node *Node) Intersect(r Ray, tmin, tmax float64) Hit {
	if node.axis == AxisNone {
		hit := NoHit
		for _, shape := range node.shapes {
			h := shape.Intersect(r)
			if h.T < hit.T {
				hit = h
			}
		}
		return hit
	}
	var tsplit float64
	var leftFirst bool
	switch node.axis {
	case AxisX:
		tsplit = (node.point - r.Origin.X) / r.Direction.X
		leftFirst = (r.Origin.X < node.point) || (r.Origin.X == node.point && r.Direction.X <= 0)
	case AxisY:
		tsplit = (node.point - r.Origin.Y) / r.Direction.Y
		leftFirst = (r.Origin.Y < node.point) || (r.Origin.Y == node.point && r.Direction.Y <= 0)
	case AxisZ:
		tsplit = (node.point - r.Origin.Z) / r.Direction.Z
		leftFirst = (r.Origin.Z < node.point) || (r.Origin.Z == node.point && r.Direction.Z <= 0)
	}
	var first, second *Node
	if leftFirst {
		first = node.left
		second = node.right
	} else {
		first = node.right
		second = node.left
	}
	if tsplit > tmax || tsplit <= 0 {
		return first.Intersect(r, tmin, tmax)
	} else if tsplit < tmin {
		return second.Intersect(r, tmin, tmax)
	} else {
		h1 := first.Intersect(r, tmin, tsplit)
		h2 := second.Intersect(r, tsplit, math.Min(tmax, h1.T))
		if h1.T <= h2.T {
			return h1
		} else {
			return h2
		}
	}
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
	mx, my, mz := Median(xs), Median(ys), Median(zs)
	xs = []float64{mx, mx}
	ys = []float64{my, my}
	zs = []float64{mz, mz}
	best := len(node.shapes)
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
	if bestAxis == AxisNone {
		return
	}
	node.axis = bestAxis
	node.point = bestPoint
	l, r := node.Partition(bestAxis, bestPoint)
	node.left = NewNode(l)
	node.right = NewNode(r)
	node.left.Split(depth + 1)
	node.right.Split(depth + 1)
	node.shapes = nil
}
