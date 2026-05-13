package main

import (
	"errors"
	"game/base"
)

type QuadTreeElement[t base.Number] interface {
	GetPos() base.Vec[t]
	GetBox() base.Vec[t]
	SetQuadTree(*QuadTree)
}

type QuadTree struct {
	eles           []QuadTreeElement[float32]
	center         base.Vec[float32]
	size           base.Vec[float32]
	higherQuadTree *QuadTree
	top_left       *QuadTree
	top_right      *QuadTree
	bottom_left    *QuadTree
	bottom_right   *QuadTree
}

const (
	TOP_LEFT = iota
	TOP_RIGHT
	BOTTOM_LEFT
	BOTTOM_RIGHT
)

func NewQuadTree(center, size base.Vec[float32], higher *QuadTree) *QuadTree {
	res := QuadTree{}
	res.center = center
	res.size = size
	res.eles = make([]QuadTreeElement[float32], 0)
	res.higherQuadTree = higher
	return &res
}

func (q *QuadTree) Insert(e QuadTreeElement[float32]) error {
	x, y := e.GetPos().Get()
	w, h := e.GetBox().Get()
	quadTrees := []int{}
	if x < q.center.X && y < q.center.Y {
		quadTrees = append(quadTrees, TOP_LEFT)
	}
	if x < q.center.X && y+h > q.center.Y {
		quadTrees = append(quadTrees, BOTTOM_LEFT)
	}
	if x+w > q.center.X && y < q.center.Y {
		quadTrees = append(quadTrees, TOP_RIGHT)
	}
	if x+w > q.center.X && y+h > q.center.Y {
		quadTrees = append(quadTrees, BOTTOM_RIGHT)
	}

	if len(quadTrees) == 0 {
		return errors.New("Error inserting elements, could not find inserting quadtree")
	}
	if len(quadTrees) > 1 {
		q.eles = append(q.eles, e)
		e.SetQuadTree(q)
		return nil
	}
	xQ, yQ := q.center.Get()
	wQ, hQ := q.size.Get()
	wQ = wQ / 2 //SUB QUADTREE WIDTH
	hQ = hQ / 2 //SUB QUADTREE HEIGHT
	switch quadTrees[0] {
	case TOP_LEFT:
		if q.top_left == nil {
			q.top_left = NewQuadTree(base.Vec[float32]{X: xQ - wQ/2, Y: yQ - hQ/2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
		}
		return q.top_left.Insert(e)
	case TOP_RIGHT:
		if q.top_right == nil {
			q.top_right = NewQuadTree(base.Vec[float32]{X: xQ + wQ/2, Y: yQ - hQ/2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
		}

		return q.top_right.Insert(e)
	case BOTTOM_LEFT:
		if q.bottom_left == nil {
			q.bottom_left = NewQuadTree(base.Vec[float32]{X: xQ - wQ/2, Y: yQ + hQ/2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
		}
		return q.bottom_left.Insert(e)
	case BOTTOM_RIGHT:
		if q.bottom_right == nil {
			q.bottom_right = NewQuadTree(base.Vec[float32]{X: xQ + wQ/2, Y: yQ + hQ/2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
		}
		return q.bottom_right.Insert(e)
	}
	return nil
}
