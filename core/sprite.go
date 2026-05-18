package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	base.ObjectBase
	sprite        map[string]rl.Texture2D
	currentType   string
	currentSprite int
	sizeSprite    base.Vec[float32]
}

func NewSprite() Sprite {
	return Sprite{}
}
func (s *Sprite) GetHitbox() *base.Hitbox {
	return nil
}

func (s *Sprite) Draw() {
	x, y := s.Pos.Get()
	rl.DrawTextureRec(
		s.sprite[s.currentType],
		rl.Rectangle{X: 0, Y: 0, Width: s.sizeSprite.X, Height: s.sizeSprite.Y},
		rl.Vector2{X: x, Y: y},
		rl.White,
	)
}
