package base

type Hitbox struct {
	Box      *Vec[float32]
	Pos      *Vec[float32]
	IsActive bool
}

func NewHitbox(w, h int32) Hitbox {
	return Hitbox{
		Box:      new(NewVec(float32(w), float32(h))),
		Pos:      new(NewVec[float32](0, 0)),
		IsActive: true,
	}
}

func (h *Hitbox) GetPos() Vec[float32] {
	return *h.Pos
}

func (h *Hitbox) GetBox() Vec[float32] {
	return *h.Box
}
