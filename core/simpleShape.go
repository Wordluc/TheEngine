package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Circle struct {
	r float32
	m base.Modifier
	base.ObjectBase
}

func NewCircle(r float32) (b Circle) {
	b.Hitbox = NewRectangle(r*2, r*2).Hitbox
	b.r = float32(r)
	b.Pos = base.Vec[float32]{}
	return b
}

func (b *Circle) Draw() {
	x, y := b.Pos.Get()
	rl.DrawCircle(int32(x+b.r), int32(y+b.r), b.r, rl.Black)
}

type Square struct {
	base.ObjectBase
	w float32
	h float32
}

func GetRectangleHitbox(w, h float32) *base.Hitbox {
	return new(base.NewHitbox()).AppendVertex(0, 0).AppendVertex(w, 0).AppendVertex(w, h).AppendVertex(0, h).AppendVertex(0, 0)
}

func NewRectangle(w, h float32) (b Square) {
	b.Hitbox = GetRectangleHitbox(w, h)
	b.w = w
	b.h = h
	return b
}

func (b *Square) Draw() {
	x, y := b.Pos.Get()
	rl.DrawRectangle(int32(x), int32(y), int32(b.w), int32(b.h), rl.Black)
}

type Triangle struct {
	base.ObjectBase
	h float32
	l float32
}

func NewTriangle(h, l float32) (b Triangle) {
	b.Hitbox = new(base.NewHitbox()).AppendVertex(0, h).AppendVertex(l, h).AppendVertex(l, 0).AppendVertex(0, h)
	b.h = h
	b.l = l
	return b
}

func (b *Triangle) Draw() {
	//x, y := b.Pos.Get()
	//rl.DrawRectangle(int32(x), int32(y), int32(b.w), int32(b.h), rl.Black)
}
