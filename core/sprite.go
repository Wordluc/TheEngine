package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	base.ObjectBase
	spriteSheets       map[string]SpriteSheet
	currentSpriteSheet string
	CurrentSprite      int8
	_currentSprite     float32
	SpeedSpriteLoop    float32
	FlippedY           bool
	FlippedX           bool
}

func NewSprite(spriteSheetPaths map[string]SpriteSheet) (s Sprite) {
	s.spriteSheets = map[string]SpriteSheet{}
	for name, spriteSheet := range spriteSheetPaths {
		s.currentSpriteSheet = name
		s.spriteSheets[name] = spriteSheet

	}
	s.SpeedSpriteLoop = 10
	return s
}

func (s *Sprite) GetHitbox() *base.Hitbox {
	return s.Hitbox
}

func (s *Sprite) SpriteLoop() {
	s._currentSprite = s._currentSprite + rl.GetFrameTime()*s.SpeedSpriteLoop
	if int32(s._currentSprite)%2 == 0 {
		spriteSheet := s.spriteSheets[s.currentSpriteSheet]
		s.CurrentSprite = max((s.CurrentSprite+1)%spriteSheet.to, spriteSheet.from)
		s._currentSprite++
	}
}

func (s *Sprite) Draw() {
	spriteSheet := s.spriteSheets[s.currentSpriteSheet]
	x, y := s.Pos.Get()
	size := spriteSheet.spriteSize
	dest := rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  size.X,
		Height: size.Y,
	}
	source := spriteSheet.GetRectangle(s.CurrentSprite)
	if s.FlippedX {
		source.Width = -source.Width
	}
	if s.FlippedY {
		source.Height = -source.Height
	}
	rl.DrawTexturePro(
		spriteSheet.Texture2D,
		source,
		dest,
		rl.Vector2{},
		0,
		rl.White,
	)
}
