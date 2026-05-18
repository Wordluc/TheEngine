package base

type Object interface {
	GetHitbox() *Hitbox
	Draw()
	MoveTo(Vec[float32])
	MoveBy(Vec[float32])
	GetPos() *Vec[float32]
}

type ObjectBase struct {
	Pos *Vec[float32]
}

func (o *ObjectBase) GetPos() *Vec[float32] {
	return o.Pos
}

func (o *ObjectBase) MoveTo(v Vec[float32]) {
	o.Pos = new(v)
}

func (o *ObjectBase) MoveBy(v Vec[float32]) {
	o.Pos.Add(v)
}
