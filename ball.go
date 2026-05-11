package main

import (
	"game/object/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	hitbox *base.Hitbox
	pos    *base.Vec
	r      float32
}

func NewBall(r int32) (b Ball) {
	b.hitbox = new(base.NewHitbox(r*2, r*2))
	b.r = float32(r)
	b.pos = new(base.Vec)
	return b
}

func (b *Ball) GetHitbox() *base.Hitbox {
	return b.hitbox
}

func (b *Ball) MoveTo(v base.Vec) {
	b.pos = new(v)
}

func (b *Ball) MoveBy(v base.Vec) {
	b.pos.Add(v)
}

func (b *Ball) GetPos() *base.Vec {
	return b.pos
}

func (b *Ball) Draw() {
	x, y := b.pos.Get()
	rl.DrawCircle(x+int32(b.r), y+int32(b.r), b.r, rl.Black)
}
