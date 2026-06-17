package base

import "errors"

type Modifier interface {
	setObject(Object)
}

type Drawable interface {
	Draw()
}

type Object interface {
	GetHitbox() *Hitbox
	MoveTo(Vec[float32])
	MoveBy(Vec[float32])
	RotateBy(float32)
	GetPos() Vec[float32]
	SetModifier(Modifier)
	GetModifiers() []Modifier
}

type ObjectBase struct {
	Pos      Vec[float32]
	Modifier []Modifier
	Hitbox   *Hitbox
	angle    float32
}

func (o *ObjectBase) GetPos() Vec[float32] {
	return o.Pos
}

func (o *ObjectBase) RotateBy(angle float32) {
	o.GetHitbox().Rotate(angle)
	o.angle = angle
}

func (o *ObjectBase) Angle() float32 {
	return o.angle
}

func (o *ObjectBase) MoveTo(v Vec[float32]) {
	o.Pos = v
}

func (o *ObjectBase) MoveBy(v Vec[float32]) {
	o.Pos.Add(v)
}

func UseModifierRef[t Modifier](o Object, c func(t)) error {
	for _, m := range o.GetModifiers() {
		if r, ok := m.(t); ok {
			c(r)
		}
	}
	return errors.New("Modifier not found")
}

func GetModifierRef[t Modifier](o Object) (r t) {
	for _, m := range o.GetModifiers() {
		if r, ok := m.(t); ok {
			return r
		}

	}
	return r
}

func (r *ObjectBase) GetHitbox() *Hitbox {
	return r.Hitbox
}

func (o *ObjectBase) GetModifiers() []Modifier {
	return o.Modifier
}

func (o *ObjectBase) SetModifier(m Modifier) {
	m.setObject(o)
	o.Modifier = append(o.Modifier, m)
}
