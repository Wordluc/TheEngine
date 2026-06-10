package base

type Hitbox struct {
	Box      *UVec[float32]
	Pos      *UVec[float32]
	IsActive bool
}

func NewHitbox(w, h float32) Hitbox {
	return Hitbox{
		Box:      new(NewVec(float32(w), float32(h))),
		Pos:      new(NewVec[float32](0, 0)),
		IsActive: true,
	}
}

func (h *Hitbox) GetPos() UVec[float32] {
	return *h.Pos
}

func (h *Hitbox) GetBox() UVec[float32] {
	return *h.Box
}

func (h *Hitbox) IntersectsPoint(p UVec[float32]) bool {
	xMax, yMax := h.Pos.X+h.Box.X, h.Pos.Y+h.Box.Y

	return p.X >= h.Pos.X && p.X <= xMax && p.Y >= p.Y && p.Y <= yMax
}
