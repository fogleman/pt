package pt

import (
	"fmt"
	"math"
	"sort"
)

type Tree struct {
	box  Box
	root *Node
}

func NewTree(shapes []Shape) *Tree {
	fmt.Printf("Building k-d tree (%d shapes)... ", len(shapes))
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
		if h1.T <= tsplit {
			return h1
		}
		h2 := second.Intersect(r, tsplit, math.Min(tmax, h1.T))
		if h1.T <= h2.T {
			return h1
		} else {
			return h2
		}
	}
}

func (node *Node) PartitionScore(axis Axis, point float64) int {
	left, right := 0, 0
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
	if left >= right {
		return left
	} else {
		return right
	}
}

func (node *Node) Partition(size int, axis Axis, point float64) (left, right []Shape) {
	left = make([]Shape, 0, size)
	right = make([]Shape, 0, size)
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
	if len(node.shapes) < 8 {
		return
	}
	xs := make([]float64, 0, len(node.shapes)*2)
	ys := make([]float64, 0, len(node.shapes)*2)
	zs := make([]float64, 0, len(node.shapes)*2)
	for _, shape := range node.shapes {
		box := shape.Box()
		xs = append(xs, box.Min.X)
		xs = append(xs, box.Max.X)
		ys = append(ys, box.Min.Y)
		ys = append(ys, box.Max.Y)
		zs = append(zs, box.Min.Z)
		zs = append(zs, box.Max.Z)
	}
	sort.Float64s(xs)
	sort.Float64s(ys)
	sort.Float64s(zs)
	mx, my, mz := Median(xs), Median(ys), Median(zs)
	best := int(float64(len(node.shapes)) * 0.85)
	bestAxis := AxisNone
	bestPoint := 0.0
	sx := node.PartitionScore(AxisX, mx)
	if sx < best {
		best = sx
		bestAxis = AxisX
		bestPoint = mx
	}
	sy := node.PartitionScore(AxisY, my)
	if sy < best {
		best = sy
		bestAxis = AxisY
		bestPoint = my
	}
	sz := node.PartitionScore(AxisZ, mz)
	if sz < best {
		best = sz
		bestAxis = AxisZ
		bestPoint = mz
	}
	if bestAxis == AxisNone {
		return
	}
	l, r := node.Partition(best, bestAxis, bestPoint)
	node.axis = bestAxis
	node.point = bestPoint
	node.left = NewNode(l)
	node.right = NewNode(r)
	node.left.Split(depth + 1)
	node.right.Split(depth + 1)
	node.shapes = nil // only needed at leaf nodes
}
