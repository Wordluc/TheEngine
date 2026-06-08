package base

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type QuadTreeElement interface {
	GetPos() Vec[float32]
	GetHitbox() *Hitbox
}

var DEBUG = false

type QuadTree struct {
	Elements       []QuadTreeElement
	Pos            Vec[float32]
	Size           Vec[float32]
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

func NewQuadTree(pos, size Vec[float32], higher *QuadTree) *QuadTree {
	res := QuadTree{}
	res.Pos = pos
	res.Size = size
	res.Elements = make([]QuadTreeElement, 0)
	res.higherQuadTree = higher
	return &res
}

func (q *QuadTree) Insert(e QuadTreeElement) error {
	if len(q.Elements) < 3 && !q.hasSub {
		q.Elements = append(q.Elements, e)
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
				q.Top_left = NewQuadTree(Vec[float32]{X: xQ, Y: yQ}, Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Top_left.higherQuadTree = q
			return q.Top_left, alreadyExist
		case TOP_RIGHT:
			if q.Top_right == nil {
				q.Top_right = NewQuadTree(Vec[float32]{X: xQ + wQ/2, Y: yQ}, Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Top_right.higherQuadTree = q
			return q.Top_right, alreadyExist
		case BOTTOM_LEFT:
			if q.Bottom_left == nil {
				q.Bottom_left = NewQuadTree(Vec[float32]{X: xQ, Y: yQ + hQ/2}, Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Bottom_left.higherQuadTree = q
			return q.Bottom_left, alreadyExist
		case BOTTOM_RIGHT:
			if q.Bottom_right == nil {
				q.Bottom_right = NewQuadTree(Vec[float32]{X: xQ + wQ/2, Y: yQ + hQ/2}, Vec[float32]{X: wQ / 2, Y: hQ / 2}, q)
				alreadyExist = false
			}
			q.Bottom_right.higherQuadTree = q
			return q.Bottom_right, alreadyExist
		}
		return nil, alreadyExist
	}

	elementsToReallocate := append(q.Elements, e)
	q.Elements = make([]QuadTreeElement, 0)
	for i := range elementsToReallocate {
		e := elementsToReallocate[i]
		pos := AddVecs(e.GetPos(), e.GetHitbox().GetPos())
		x, y := pos.Get()
		w, h := e.GetHitbox().GetBox().Get()
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
			continue
		}
		quadtree, _ := createQuadTree(quadtreeToCreate[0])
		if quadtree == nil {
			return fmt.Errorf("Error creating quadtree in %v", quadtreeToCreate[0])
		}
		q.hasSub = true
		err := quadtree.Insert(e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *QuadTree) subQuery(elements []QuadTreeElement, forEach func(o []QuadTreeElement)) {
	elements = append(elements, q.Elements...)
	if q.Top_left != nil {
		q.Top_left.subQuery(elements, forEach)
	}
	if q.Top_right != nil {
		q.Top_right.subQuery(elements, forEach)
	}
	if q.Bottom_left != nil {
		q.Bottom_left.subQuery(elements, forEach)
	}
	if q.Bottom_right != nil {
		q.Bottom_right.subQuery(elements, forEach)
	}
	if DEBUG {
		var direction string
		if q.higherQuadTree != nil {
			if q.higherQuadTree.Bottom_right == q {
				direction = "Bottom_right"
			} else if q.higherQuadTree.Bottom_left == q {
				direction = "Bottom_left"
			} else if q.higherQuadTree.Top_right == q {
				direction = "Top_right"
			} else if q.higherQuadTree.Top_left == q {
				direction = "Top_left"
			}
			fmt.Println(direction, "  ", len(elements))
		}
	}
	forEach(elements)
}

func (q *QuadTree) Query(forEach func([]QuadTreeElement)) {
	q.subQuery(nil, forEach)
}

func (q *QuadTree) Clear() {
	q.Elements = q.Elements[:0] // keep backing array

	if q.Top_left != nil {
		q.Top_left.Clear()
	}
	if q.Top_right != nil {
		q.Top_right.Clear()
	}
	if q.Bottom_left != nil {
		q.Bottom_left.Clear()
	}
	if q.Bottom_right != nil {
		q.Bottom_right.Clear()
	}
}

func (q *QuadTree) DrawBorder() {
	if len(q.Elements) != 0 {
		rl.DrawRectangleLines(int32(q.Pos.X), int32(q.Pos.Y), int32(q.Size.X), int32(q.Size.Y), rl.Red)
		direction := ""
		if q.higherQuadTree != nil {
			if q.higherQuadTree.Bottom_right == q {
				direction = "Bottom_right"
			} else if q.higherQuadTree.Bottom_left == q {
				direction = "Bottom_left"
			} else if q.higherQuadTree.Top_right == q {
				direction = "Top_right"
			} else if q.higherQuadTree.Top_left == q {
				direction = "Top_left"
			}
			rl.DrawText(direction, int32(q.Pos.X), int32(q.Pos.Y), 20, rl.Red)
		}
	}
	if q.Top_left != nil {
		q.Top_left.DrawBorder()
	}
	if q.Top_right != nil {
		q.Top_right.DrawBorder()
	}
	if q.Bottom_left != nil {
		q.Bottom_left.DrawBorder()
	}
	if q.Bottom_right != nil {
		q.Bottom_right.DrawBorder()
	}
}
