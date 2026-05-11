package base

type Hitbox struct {
	*Vec
}

func NewHitbox(w, h int32) Hitbox {
	return Hitbox{
		new(NewVec(w, h)),
	}
}
