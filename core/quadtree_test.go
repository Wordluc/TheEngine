package core

import (
	"game/core/base"

	"testing"
)

type MockElement struct {
	Pos base.Vec[float32]
	Box base.Vec[float32]
	qt  *QuadTree
}

func (m *MockElement) GetPos() base.Vec[float32] { return m.Pos }
func (m *MockElement) GetBox() base.Vec[float32] { return m.Box }
func (m *MockElement) SetQuadTree(q *QuadTree)   { m.qt = q }

func elem(x, y, w, h float32) *MockElement {
	return &MockElement{
		Pos: base.Vec[float32]{X: x, Y: y},
		Box: base.Vec[float32]{X: w, Y: h},
	}
}

func rootTree() *QuadTree {
	center := base.Vec[float32]{X: 0, Y: 0}
	size := base.Vec[float32]{X: 50, Y: 50}
	return NewQuadTree(center, size, nil)
}

func TestNewQuadTree_InitialisedCorrectly(t *testing.T) {
	q := rootTree()
	if q == nil {
		t.Fatal("NewQuadTree returned nil")
	}
}

func TestInsert_1Element(t *testing.T) {
	q := rootTree()
	e := elem(10, 10, 5, 5)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(q.Elements) != 1 {
		t.Errorf("Element should  be stored in the root quadtree")
	}
}

func TestInsert_3Element(t *testing.T) {
	q := rootTree()
	es := []*MockElement{
		elem(10, 10, 5, 5),
		elem(20, 10, 5, 5),
		elem(20, 20, 5, 5),
	}
	for _, e := range es {
		if err := q.Insert(e); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(q.Elements) != 3 {
		t.Errorf("Element should  be stored in the root quadtree")
	}
}

func TestInsert_4Element(t *testing.T) {
	q := rootTree()
	es := []*MockElement{
		elem(10, 10, 5, 5),
		elem(20, 10, 5, 5),
		elem(20, 20, 5, 5),
		elem(20, 30, 5, 5),
	}
	for _, e := range es {
		if err := q.Insert(e); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(q.Elements) != 0 {
		t.Errorf("Element should't  be stored in the root quadtree")
	}

	if len(q.Top_left.Elements) != 3 {
		t.Errorf("3 Elements should  be stored in the top_left quadtree")
	}

	if len(q.Bottom_left.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the bottom_left quadtree")
	}
}

// Root [0,0 → 50×50]
// │
// ├── Top_left [0,0 → 25×25]
// │   │
// │   ├── Top_left (TL·TL)  [0,0 → 12.5×12.5]
// │   │       ├── A  pos(5,5)  box(2×2)
// │   │       └── B  pos(6,5)  box(2×2)
// │   │
// │   ├── Top_right (TL·TR) [12.5,0 → 12.5×12.5]
// │   │       └── C  pos(14,5) box(2×2)
// │   │
// │   ├── Bottom_left (TL·BL) [0,12.5 → 12.5×12.5]
// │   │       └── D  pos(4,15) box(2×2)
// │   │
// │   └── Bottom_right (TL·BR) [12.5,12.5 → 12.5×12.5]
// │           └── E  pos(17,17) box(2×2)
// │
// ├── Top_right    [empty / null]
// ├── Bottom_left  [empty / null]
// └── Bottom_right [empty / null]

func TestInsert_5Element_SplitTopLeft(t *testing.T) {
	q := rootTree()
	es := []*MockElement{
		elem(5, 5, 2, 2),
		elem(6, 5, 2, 2),
		elem(14, 5, 2, 2),
		elem(4, 15, 2, 2),
		elem(17, 17, 2, 2),
	}
	for _, e := range es {
		if err := q.Insert(e); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(q.Elements) != 0 {
		t.Errorf("Element should't  be stored in the root quadtree")
	}
	if len(q.Top_left.Top_left.Elements) != 2 {
		t.Errorf("2 Elements should  be stored in the top_left->top_left quadtree")
	}
	if len(q.Top_left.Top_right.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->top_right quadtree")
	}
	if len(q.Top_left.Bottom_left.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->bottom_left quadtree")
	}

	if len(q.Top_left.Bottom_right.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->bottom_right quadtree")
	}
}

// Root [0,0 → 50×50]
// │
// ├── Top_left [0,0 → 25×25]
// │   │   Elements:  pos(12,5)  box(5×2)  ← spans TL·TL and TL·TR
// │   │
// │   ├── TL · Top_left     [0,0 → 12.5×12.5]
// │   │       ├── A  pos(5,5)   box(2×2)
// │   │       └── B  pos(6,5)   box(2×2)
// │   │
// │   ├── TL · Top_right    [12.5,0 → 12.5×12.5]
// │   │       └── C  pos(14,5)  box(2×2)
// │   │
// │   ├── TL · Bottom_left  [0,12.5 → 12.5×12.5]
// │   │       └── D  pos(4,15)  box(2×2)
// │   │
// │   └── TL · Bottom_right [12.5,12.5 → 12.5×12.5]
// │           └── E  pos(17,17) box(2×2)
// │
// ├── Top_right     [null]
// ├── Bottom_left   [null]
// └── Bottom_right  [null]

func TestInsert_5Element_ElementIsInMoreQuadtree(t *testing.T) {
	q := rootTree()
	es := []*MockElement{
		elem(5, 5, 2, 2),
		elem(6, 5, 2, 2),
		elem(12, 5, 5, 2),
		elem(14, 5, 2, 2),
		elem(4, 15, 2, 2),
		elem(17, 17, 2, 2),
	}
	for _, e := range es {
		if err := q.Insert(e); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if len(q.Elements) != 0 {
		t.Errorf("Element should's be stored in root")
	}
	if len(q.Top_left.Elements) != 1 {
		t.Errorf("1 Element should be stored in root top_left")
	}
	if len(q.Top_left.Top_left.Elements) != 2 {
		t.Errorf("3 Elements should  be stored in the top_left->top_left quadtree")
	}
	if len(q.Top_left.Top_right.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->top_right quadtree")
	}
	if len(q.Top_left.Bottom_left.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->bottom_left quadtree")
	}
	if len(q.Top_left.Bottom_right.Elements) != 1 {
		t.Errorf("1 Elements should  be stored in the top_left->bottom_right quadtree")
	}
}
