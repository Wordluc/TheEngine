package main

import (
	"game/base"
	"testing"
)

// ---------------------------------------------------------------------------
// Test double: a simple element that satisfies QuadTreeElement[float32]
// ---------------------------------------------------------------------------

type MockElement struct {
	pos base.Vec[float32] // top-left corner
	box base.Vec[float32] // width, height
	qt  *QuadTree
}

func (m *MockElement) GetPos() base.Vec[float32] { return m.pos }
func (m *MockElement) GetBox() base.Vec[float32] { return m.box }
func (m *MockElement) SetQuadTree(q *QuadTree)   { m.qt = q }

// helper to build a MockElement at (x,y) with size (w,h)
func elem(x, y, w, h float32) *MockElement {
	return &MockElement{
		pos: base.Vec[float32]{X: x, Y: y},
		box: base.Vec[float32]{X: w, Y: h},
	}
}

// ---------------------------------------------------------------------------
// Root tree used in most tests
// World: center=(50,50), half-size=(50,50) → covers (0,0)→(100,100)
// ---------------------------------------------------------------------------

func rootTree() *QuadTree {
	center := base.Vec[float32]{X: 25, Y: 25}
	size := base.Vec[float32]{X: 50, Y: 50}
	return NewQuadTree(center, size, nil)
}

// ---------------------------------------------------------------------------
// 1. NewQuadTree – basic construction
// ---------------------------------------------------------------------------

func TestNewQuadTree_InitialisedCorrectly(t *testing.T) {
	q := rootTree()
	if q == nil {
		t.Fatal("NewQuadTree returned nil")
	}
}

// ---------------------------------------------------------------------------
// 2. Insert – element lands entirely in TOP_LEFT quadrant
//    Quadrant boundary: x < 50 AND y < 50
//    Element (10,10) 5×5 → entirely in TOP_LEFT → must recurse, NOT stay at root
// ---------------------------------------------------------------------------

func TestInsert_TopLeftQuadrant_RecursesDown(t *testing.T) {
	q := rootTree()
	e := elem(10, 10, 5, 5) // fully inside top-left quadrant

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// The element must NOT be stored at the root
	if e.qt == q {
		t.Error("element should have been pushed to a child node, not stored at root")
	}
	// The element must be stored somewhere (qt not nil)
	if e.qt == nil {
		t.Error("SetQuadTree was never called; element has no owning node")
	}
}

func TestInsert_TopRightQuadrant_RecursesDown(t *testing.T) {
	q := rootTree()
	e := elem(40, 10, 5, 5)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt == q {
		t.Error("element should have been pushed to a child node")
	}
}

func TestInsert_BottomLeftQuadrant_RecursesDown(t *testing.T) {
	q := rootTree()
	e := elem(10, 40, 5, 5)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt == q {
		t.Error("element should have been pushed to a child node")
	}
}

func TestInsert_BottomRightQuadrant_RecursesDown(t *testing.T) {
	q := rootTree()
	e := elem(40, 40, 5, 5)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt == q {
		t.Error("element should have been pushed to a child node")
	}
}

func TestInsert_StradlesHorizontalCenter_StoredAtCurrentNode(t *testing.T) {
	q := rootTree()
	e := elem(40, 10, 20, 5)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt != q {
		t.Error("element straddling the vertical divider should stay at the root node")
	}
}

func TestInsert_StradlesVerticalCenter_StoredAtCurrentNode(t *testing.T) {
	q := rootTree()
	e := elem(10, 40, 5, 20)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt != q {
		t.Error("element straddling the horizontal divider should stay at the root node")
	}
}

// ---------------------------------------------------------------------------
// 8. Insert – element straddles BOTH axes (all 4 quadrants) → stored at root
//    Element (40,40) 20×20
// ---------------------------------------------------------------------------

func TestInsert_StradlesBothAxes_StoredAtCurrentNode(t *testing.T) {
	q := rootTree()
	e := elem(40, 40, 20, 20)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt != q {
		t.Error("element spanning all 4 quadrants should stay at the root node")
	}
}

// ---------------------------------------------------------------------------
// 9. Insert – multiple elements in different quadrants don't interfere
// ---------------------------------------------------------------------------

func TestInsert_MultipleElements_IndependentPlacement(t *testing.T) {
	q := rootTree()

	tl := elem(10, 10, 5, 5) // top-left
	tr := elem(40, 10, 5, 5) // top-right
	bl := elem(10, 40, 5, 5) // bottom-left
	br := elem(40, 40, 5, 5) // bottom-right

	for _, e := range []*MockElement{tl, tr, bl, br} {
		if err := q.Insert(e); err != nil {
			t.Fatalf("unexpected error inserting element: %v", err)
		}
	}

	// All four must be owned by child nodes (recursed), not by root
	for _, e := range []*MockElement{tl, tr, bl, br} {
		if e.qt == q {
			t.Errorf("element at (%v,%v) should have been placed in a child",
				e.pos.X, e.pos.Y)
		}
		if e.qt == nil {
			t.Errorf("element at (%v,%v) has no owning node", e.pos.X, e.pos.Y)
		}
	}
}

// ---------------------------------------------------------------------------
// 10. Insert – deep recursion: inserting into an already-split child
//     Two elements both in top-left but different sub-quadrants
// ---------------------------------------------------------------------------

func TestInsert_DeepRecursion_SameQuadrantTwice(t *testing.T) {
	q := rootTree()

	// Both land in TOP_LEFT of root (x<50, y<50)
	// Within TOP_LEFT child (center≈25,25): a is TL-of-TL, b is TR-of-TL
	a := elem(5, 5, 2, 2)  // top-left of top-left
	b := elem(35, 5, 2, 2) // top-right of top-left

	if err := q.Insert(a); err != nil {
		t.Fatalf("unexpected error inserting a: %v", err)
	}
	if err := q.Insert(b); err != nil {
		t.Fatalf("unexpected error inserting b: %v", err)
	}

	// Both must be placed somewhere, and they shouldn't share the same node
	if a.qt == nil || b.qt == nil {
		t.Fatal("one or both elements were never assigned a node")
	}
	if a.qt == b.qt {
		t.Error("elements in different sub-quadrants should not share the same node")
	}
}

// ---------------------------------------------------------------------------
// 11. Insert – element exactly on the centre point
//     pos=(50,50) box=(1,1): x+w=51 > 50 and y+h=51 > 50 → BOTTOM_RIGHT only
//     (x < 50 is false, y < 50 is false) → single quadrant → recurses
// ---------------------------------------------------------------------------

func TestInsert_ElementOnCenterPoint_RecursesBottomRight(t *testing.T) {
	q := rootTree()
	e := elem(40, 40, 1, 1)

	if err := q.Insert(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.qt == q {
		t.Error("element just past the centre should recurse into bottom-right child")
	}
}

// ---------------------------------------------------------------------------
// 12. SetQuadTree back-reference is always the owning node
// ---------------------------------------------------------------------------

func TestInsert_SetQuadTreePointsToOwningNode(t *testing.T) {
	q := rootTree()

	// This element straddles the boundary so it must live at root
	e := elem(40, 40, 20, 20)
	_ = q.Insert(e)

	if e.qt != q {
		t.Errorf("expected SetQuadTree to be called with root node, got %v", e.qt)
	}
}
