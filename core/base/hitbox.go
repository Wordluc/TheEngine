package base

import (
	"slices"
)

type Hitbox struct {
	outerBox *Vec[float32]
	Pos      *Vec[float32]
	vertex   []Vec[float32]
	IsActive bool
}

func NewHitbox(w, h float32) Hitbox {
	return Hitbox{
		outerBox: new(NewVec[float32](0, 0)),
		Pos:      new(NewVec[float32](0, 0)),
		IsActive: true,
		vertex:   make([]Vec[float32], 0),
	}
}

func (h *Hitbox) GetPos() Vec[float32] {
	return *h.Pos
}

func (h *Hitbox) AppendVertex(x, y float32) *Hitbox {
	if h.outerBox.X < x {
		h.outerBox.X = x
	}
	if h.outerBox.Y < y {
		h.outerBox.Y = y
	}
	h.vertex = append(h.vertex, Vec[float32]{x, y})
	return h
}

func (h *Hitbox) ProjectionOn(posObject Vec[float32], v Vec[float32]) (min, max float32) {
	if len(h.vertex) == 0 {
		return 0, 0
	}

	pos := AddVecs(posObject, h.GetPos())
	vert := h.vertex[0].Clone().Add(pos)

	initial := DotProduct(*vert, v)
	min, max = initial, initial
	var proj float32
	for _, vert := range h.vertex[1:] {
		vert = *vert.Add(pos)
		proj = DotProduct(vert, v)
		if proj < min {
			min = proj
		}
		if proj > max {
			max = proj
		}
	}
	return min, max
}

func (h *Hitbox) GetVertex() []Vec[float32] {
	return slices.Clone(h.vertex)
}

func (h *Hitbox) GetOuterBox() Vec[float32] {
	return *h.outerBox
}

func (h *Hitbox) IntersectsPoint(p Vec[float32]) bool {
	xMax, yMax := h.Pos.X+h.outerBox.X, h.Pos.Y+h.outerBox.Y

	return p.X >= h.Pos.X && p.X <= xMax && p.Y >= p.Y && p.Y <= yMax
}
