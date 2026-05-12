package main

import (
	"game/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	hitbox *base.Hitbox
	pos    *base.Vec[int32]
	r      float32
}

func NewBall(r int32) (b Ball) {
	b.hitbox = new(base.NewHitbox(r*2, r*2))
	b.r = float32(r)
	b.pos = new(base.Vec[int32])
	return b
}

func (b *Ball) GetHitbox() *base.Hitbox {
	return b.hitbox
}

func (b *Ball) MoveTo(v base.Vec[int32]) {
	b.pos = new(v)
}

func (b *Ball) MoveBy(v base.Vec[int32]) {
	b.pos.Add(v)
}

func (b *Ball) GetPos() *base.Vec[int32] {
	return b.pos
}

func (b *Ball) Draw() {
	x, y := b.pos.Get()
	rl.DrawCircle(x+int32(b.r), y+int32(b.r), b.r, rl.Black)
}

type Square struct {
	hitbox *base.Hitbox
	pos    *base.Vec[int32]
	l      float32
}

func NewSquare(l int32) (b Square) {
	b.hitbox = new(base.NewHitbox(l, l))
	b.l = float32(l)
	b.pos = new(base.Vec[int32])
	return b
}

func (b *Square) GetHitbox() *base.Hitbox {
	return b.hitbox
}

func (b *Square) MoveTo(v base.Vec[int32]) {
	b.pos = new(v)
}

func (b *Square) MoveBy(v base.Vec[int32]) {
	b.pos.Add(v)
}

func (b *Square) GetPos() *base.Vec[int32] {
	return b.pos
}

func (b *Square) Draw() {
	x, y := b.pos.Get()
	rl.DrawRectangle(x, y, int32(b.l), int32(b.l), rl.Black)
}
