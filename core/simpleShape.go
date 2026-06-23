package core

import (
	"github.com/Wordluc/TheEngine/core/base"

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
	return b
}

func (b *Circle) Draw() {
	x, y := b.GetPos().Get()
	rl.DrawCircle(int32(x+b.r), int32(y+b.r), b.r, rl.Black)
}

type Square struct {
	base.ObjectBase
	w float32
	h float32
}

func GetRectangleHitbox(x, y, w, h float32) *base.Hitbox {
	hitbox := new(base.NewHitbox()).AppendVertex(0, 0).AppendVertex(w, 0).AppendVertex(w, h).AppendVertex(0, h).AppendVertex(0, 0)
	hitbox.Pos = new(base.NewVec[float32](x, y))
	return hitbox
}

func NewRectangle(w, h float32) (b Square) {
	b.Hitbox = GetRectangleHitbox(0, 0, w, h)
	b.w = w
	b.h = h
	return b
}

func (b *Square) Draw() {
	x, y := b.GetPos().Get()
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
