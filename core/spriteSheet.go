package core

import (
	"fmt"
	"game/core/base"
	"game/core/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteSheet struct {
	rl.Texture2D
	totalCols  int8
	totalRows  int8
	offset     int8
	spriteSize base.Vec[float32]
}

func NewSpriteSheet(path string, spriteSize base.Vec[float32], offset int8) utils.Result[SpriteSheet] {
	texture := rl.LoadTexture(path)
	if texture.Width == 0 {
		return utils.ResultErr[SpriteSheet](fmt.Errorf("Error loading sprite '%v' ", path))
	}

	return utils.ResultOk(SpriteSheet{
		Texture2D:  texture,
		totalCols:  int8(texture.Width / int32(spriteSize.X)),
		totalRows:  int8(texture.Height / int32(spriteSize.Y)),
		spriteSize: spriteSize,
		offset:     offset,
	})
}

func (s SpriteSheet) GetRectangle(currentSprite int8) rl.Rectangle {
	col := float32(currentSprite%s.totalCols) * (s.spriteSize.X + float32(s.offset)*2)
	row := float32(currentSprite/s.totalRows) * (s.spriteSize.Y + float32(s.offset)*2)
	return rl.Rectangle{
		X:      col,
		Y:      row,
		Width:  s.spriteSize.X,
		Height: s.spriteSize.Y,
	}
}
