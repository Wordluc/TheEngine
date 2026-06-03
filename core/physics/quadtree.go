package physics

import (
	"errors"
	"fmt"
	"game/core/base"
)

type QuadTreeElement[t base.Number] interface {
	GetHitbox() *base.Hitbox
	MoveBy(base.Vec[float32])
	SetQuadTree(*QuadTree)
}

type QuadTree struct {
	Elements       []QuadTreeElement[float32]
	Pos            base.Vec[float32]
	Size           base.Vec[float32]
	hasSub         bool
	higherQuadTree *QuadTree
	Top_left       *QuadTree
	Top_right      *QuadTree
	Bottom_left    *QuadTree
	Bottom_right   *QuadTree
}

const (
	TOP_LEFT     = "TOP_LEFT"
	TOP_RIGHT    = "TOP_RIGHT"
	BOTTOM_LEFT  = "BOTTOM_LEFT"
	BOTTOM_RIGHT = "BOTTOM_RIGHT"
)

func NewQuadTree(pos, size base.Vec[float32], higher *QuadTree) *QuadTree {
	res := QuadTree{}
	res.Pos = pos
	res.Size = size
	res.Elements = make([]QuadTreeElement[float32], 0)
	res.higherQuadTree = higher
	return &res
}

func (q *QuadTree) Insert(e QuadTreeElement[float32]) error {
	if len(q.Elements) < 3 && !q.hasSub {
		q.Elements = append(q.Elements, e)
		e.SetQuadTree(q)
		return nil
	}

	xQ, yQ := q.Pos.Get()
	wQ, hQ := q.Size.Get()
	centerX, centerY := xQ+wQ/2, yQ+hQ/2

	createQuadTree := func(where string) (quadtree *QuadTree, alreadyExist bool) {
		alreadyExist = true
		switch where {
		case TOP_LEFT:
			if q.Top_left == nil {
				q.Top_left = NewQuadTree(base.Vec[float32]{X: 0, Y: 0}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Top_left.higherQuadTree = q
			return q.Top_left, alreadyExist
		case TOP_RIGHT:
			if q.Top_right == nil {
				q.Top_right = NewQuadTree(base.Vec[float32]{X: wQ / 2, Y: 0}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Top_right.higherQuadTree = q
			return q.Top_right, alreadyExist
		case BOTTOM_LEFT:
			if q.Bottom_left == nil {
				q.Bottom_left = NewQuadTree(base.Vec[float32]{X: 0, Y: hQ / 2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Bottom_left.higherQuadTree = q
			return q.Bottom_left, alreadyExist
		case BOTTOM_RIGHT:
			if q.Bottom_right == nil {
				q.Bottom_right = NewQuadTree(base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, base.Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Bottom_right.higherQuadTree = q
			return q.Bottom_right, alreadyExist
		}
		return nil, false
	}

	elementsToReallocate := append(q.Elements, e)
	q.Elements = make([]QuadTreeElement[float32], 0)
	var hitbox base.Hitbox
	for i := range elementsToReallocate {
		e := elementsToReallocate[i]
		hitbox = *e.GetHitbox()
		x, y := hitbox.GetPos().Get()
		w, h := hitbox.GetBox().Get()
		quadtreeToCreate := []string{}
		if x < centerX && y < centerY {
			quadtreeToCreate = append(quadtreeToCreate, TOP_LEFT)
		}
		if x < centerX && y+h > centerY {
			quadtreeToCreate = append(quadtreeToCreate, BOTTOM_LEFT)
		}
		if x+w > centerX && y < centerY {
			quadtreeToCreate = append(quadtreeToCreate, TOP_RIGHT)
		}
		if x+w > centerX && y+h > centerY {
			quadtreeToCreate = append(quadtreeToCreate, BOTTOM_RIGHT)
		}
		if len(quadtreeToCreate) == 0 {
			return errors.New("Error inserting elements, could not find inserting quadtree")
		}
		if len(quadtreeToCreate) > 1 {
			q.Elements = append(q.Elements, e)
			e.SetQuadTree(q)
			continue
		}
		quadtree, _ := createQuadTree(quadtreeToCreate[0])
		if quadtree == nil {
			return fmt.Errorf("Error creating quadtree in %v", quadtreeToCreate[0])
		}
		q.hasSub = true
		err := quadtree.Insert(e)
		e.SetQuadTree(quadtree)
		if err != nil {
			return err
		}
	}

	return nil
}
