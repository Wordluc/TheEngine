package base

type Modifier interface {
	SetObject(Modificable)
}

type Modificable interface {
	SetModifier(Modifier)
}

type Object interface {
	GetHitbox() *Hitbox
	Draw()
	MoveTo(Vec[float32])
	MoveBy(Vec[float32])
	GetPos() *Vec[float32]
}

type ObjectBase struct {
	Pos      Vec[float32]
	Modifier []Modifier
}

func (o *ObjectBase) GetPos() *Vec[float32] {
	return &o.Pos
}

func (o *ObjectBase) MoveTo(v Vec[float32]) {
	o.Pos = v
}

func (o *ObjectBase) MoveBy(v Vec[float32]) {
	o.Pos.Add(v)
}

func (o *ObjectBase) SetModifier(m Modifier) {
	m.SetObject(o)
	o.Modifier = append(o.Modifier, m)
}
