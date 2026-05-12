package base

type Hitbox struct {
	*Vec[int32]
}

func NewHitbox(w, h int32) Hitbox {
	return Hitbox{
		new(NewVec(w, h)),
	}
}
