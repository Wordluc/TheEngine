package core

import (
	"fmt"
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteSheet struct {
	rl.Texture2D
	cols int32
	rows int32
}
type Sprite struct {
	base.ObjectBase
	spriteSheets       map[string]SpriteSheet
	currentSpriteSheet string
	CurrentSprite      int
	_currentSprite     float32
	spriteSize         base.Vec[float32]
	SpeedSpriteLoop    float32
}

func NewSprite(spriteSize base.Vec[float32], spriteSheetPaths map[string]string) (s Sprite, err error) {
	s.spriteSheets = map[string]SpriteSheet{}
	for name, path := range spriteSheetPaths {
		texture := rl.LoadTexture(path)
		if texture.Width == 0 {
			return s, fmt.Errorf("Error loading sprite '%v'  in '%v' ", name, path)
		}
		s.spriteSheets[name] = SpriteSheet{
			Texture2D: texture,
			cols:      texture.Width / int32(spriteSize.X),
			rows:      texture.Height / int32(spriteSize.Y),
		}
		s.currentSpriteSheet = name

	}
	s.spriteSize = spriteSize
	s.SpeedSpriteLoop = 10
	return s, nil
}

func (s *Sprite) GetHitbox() *base.Hitbox {
	return nil
}

func (s *Sprite) SpriteLoop() {
	s._currentSprite = s._currentSprite + rl.GetFrameTime()*s.SpeedSpriteLoop
	if int32(s._currentSprite)%2 == 0 {
		s.CurrentSprite++
		s._currentSprite++
	}
}

func (s *Sprite) Draw() {
	sheet := s.spriteSheets[s.currentSpriteSheet]
	x, y := s.Pos.Get()
	col, row := float32(s.CurrentSprite%int(sheet.cols))*s.spriteSize.X, float32(s.CurrentSprite/int(sheet.rows))*s.spriteSize.Y
	rl.DrawTextureRec(
		sheet.Texture2D,
		rl.Rectangle{X: col, Y: row, Width: s.spriteSize.X, Height: s.spriteSize.Y},
		rl.Vector2{X: x, Y: y},
		rl.White,
	)
}
