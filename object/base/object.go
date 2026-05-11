package base

type Object interface {
	GetHitbox() *Hitbox
	MoveTo(Vec)
	MoveBy(Vec)
	GetPos() *Vec
	Draw()
}
