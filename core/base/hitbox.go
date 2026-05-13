package base

type Hitbox struct {
	Box      *Vec[float32]
	IsActive bool
	o        Object
}

func NewHitbox(o Object, w, h int32) Hitbox {
	return Hitbox{
		Box:      new(NewVec[float32](float32(w), float32(h))),
		IsActive: true,
		o:        o,
	}
}

func (h *Hitbox) GetPos() Vec[float32] {
	return *h.o.GetPos()
}

func (h *Hitbox) GetBox() Vec[float32] {
	return *h.Box
}
