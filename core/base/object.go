package base

type Object interface {
	GetHitbox() *Hitbox
	MoveTo(Vec[float32])
	MoveBy(Vec[float32])
	GetPos() *Vec[float32]
	Draw()
}
