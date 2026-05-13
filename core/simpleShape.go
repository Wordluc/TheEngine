package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	hitbox *base.Hitbox
	pos    *base.Vec[float32]
	r      float32
}

func NewBall(r int32) (b Ball) {
	b.hitbox = new(base.NewHitbox(&b, r*2, r*2))
	b.r = float32(r)
	b.pos = new(base.Vec[float32])
	return b
}

func (b *Ball) GetHitbox() *base.Hitbox {
	return b.hitbox
}

func (b *Ball) MoveTo(v base.Vec[float32]) {
	b.pos = new(v)
}

func (b *Ball) MoveBy(v base.Vec[float32]) {
	b.pos.Add(v)
}

func (b *Ball) GetPos() *base.Vec[float32] {
	return b.pos
}

func (b *Ball) Draw() {
	x, y := b.pos.Get()
	rl.DrawCircle(int32(x+b.r), int32(y+b.r), b.r, rl.Black)
}

type Square struct {
	hitbox *base.Hitbox
	pos    *base.Vec[float32]
	l      float32
}

func NewSquare(l int32) (b Square) {
	b.hitbox = new(base.NewHitbox(&b, l, l))
	b.l = float32(l)
	b.pos = new(base.Vec[float32])
	return b
}

func (b *Square) GetHitbox() *base.Hitbox {
	return b.hitbox
}

func (b *Square) MoveTo(v base.Vec[float32]) {
	b.pos = new(v)
}

func (b *Square) MoveBy(v base.Vec[float32]) {
	b.pos.Add(v)
}

func (b *Square) GetPos() *base.Vec[float32] {
	return b.pos
}

func (b *Square) Draw() {
	x, y := b.pos.Get()
	rl.DrawRectangle(int32(x), int32(y), int32(b.l), int32(b.l), rl.Black)
}
