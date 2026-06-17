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
	Hitbox   *Hitbox
	pos      Vec[float32]
	modifier []Modifier
	angle    float32
	Node
}

func (o *ObjectBase) GetPos() Vec[float32] {
	return o.pos
}

func (o *ObjectBase) RotateBy(angle float32) {
	if h := o.GetHitbox(); h != nil {
		h.Rotate(angle)
	}
	o.angle = angle
}

func (o *ObjectBase) Angle() float32 {
	return o.angle
}

func (o *ObjectBase) MoveTo(v Vec[float32]) {
	o.pos = v
	var diff Vec[float32]
	o.ForEachObjects(func(o Object) error {
		diff = SubVecs(v, o.GetPos())
		o.MoveBy(diff)
		return nil
	})
}

func (o *ObjectBase) MoveBy(v Vec[float32]) {
	o.pos.Add(v)
	o.ForEachObjects(func(o Object) error {
		o.MoveBy(v)
		return nil
	})
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
	return o.modifier
}

func (o *ObjectBase) SetModifier(m Modifier) {
	m.setObject(o)
	o.modifier = append(o.modifier, m)
}
