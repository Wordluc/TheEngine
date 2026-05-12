package base

type Object interface {
	GetHitbox() *Hitbox
	MoveTo(Vec[int32])
	MoveBy(Vec[int32])
	GetPos() *Vec[int32]
	Draw()
}
