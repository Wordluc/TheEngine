package base

type Hitbox struct {
	Box      *Vec[int32]
	IsActive bool
	o        Object
}

func NewHitbox(o Object, w, h int32) Hitbox {
	return Hitbox{
		Box:      new(NewVec(w, h)),
		IsActive: true,
		o:        o,
	}
}
