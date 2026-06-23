package core

import (
	"github.com/Wordluc/TheEngine/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MultiSprite struct {
	base.ObjectBase
	spriteSheet       SpriteSheet
	SpriteSheetOffset int
	SelectedSprite    int8
	blockSize         base.Vec[int32]
	spriteSize        base.Vec[float32]
}

func NewMultiSprite(spriteSheet SpriteSheet, blockSize base.Vec[int32]) (s MultiSprite) {
	s.spriteSheet = spriteSheet
	s.blockSize = blockSize
	if s.spriteSize.IsNull() {
		s.spriteSize = s.spriteSheet.spriteSize
	}
	s.Hitbox = GetRectangleHitbox(0, 0, float32(s.blockSize.X)*s.spriteSize.X, float32(s.blockSize.Y)*s.spriteSize.Y)

	return s
}

func (s *MultiSprite) Draw() {
	x, y := s.GetPos().Get()
	source := s.spriteSheet.GetRectangle(s.SelectedSprite)
	var i_y, i_x int32
	size := s.spriteSize
	if size.IsNull() {
		size = s.spriteSheet.spriteSize
	}
	var dest rl.Rectangle
	for range s.blockSize.Y {
		for range s.blockSize.X {

			dest.X = x + float32(i_x*int32(size.X))
			dest.Y = y + float32(i_y*int32(size.Y))
			dest.Width = size.X
			dest.Height = size.Y
			rl.DrawTexturePro(
				s.spriteSheet.Texture2D,
				source,
				dest,
				rl.Vector2{},
				0,
				rl.White,
			)
			i_x++
		}
		i_y++
	}
}
