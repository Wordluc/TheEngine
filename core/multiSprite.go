package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MultiSprite struct {
	base.ObjectBase
	spriteSheet       SpriteSheet
	SpriteSheetOffset int
	SelectedSprite    int8
	BlockSize         base.Vec[int32]
	SpriteSize        base.Vec[float32]
}

func NewMultiSprite(spriteSheet SpriteSheet) (s MultiSprite) {
	s.spriteSheet = spriteSheet
	return s
}

func (s *MultiSprite) GetHitbox() *base.Hitbox {
	return nil
}

func (s *MultiSprite) Draw() {
	x, y := s.Pos.Get()
	source := s.spriteSheet.GetRectangle(s.SelectedSprite)
	var i_y, i_x int32
	size := s.SpriteSize
	if size.IsNull() {
		size = s.spriteSheet.spriteSize
	}
	for range s.BlockSize.Y {
		for range s.BlockSize.X {
			rl.DrawTexturePro(
				s.spriteSheet.Texture2D,
				source,
				rl.Rectangle{
					X:      x + float32(i_x*int32(size.X)),
					Y:      y + float32(i_y*int32(size.Y)),
					Width:  size.X,
					Height: size.Y,
				},
				rl.Vector2{},
				0,
				rl.White,
			)
			i_x++
		}
		i_y++
	}
}
