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
	offset     base.Vec[float32]
	spriteSize base.Vec[float32]
	from       int8
	to         int8
}

// Use this to create a spriteSheet with a texture already loaded
func NewSpriteSheetWithTexture(texture rl.Texture2D, spriteSize base.Vec[float32], offset base.Vec[float32], from, to int8) utils.Result[SpriteSheet] {
	if texture.Width == 0 {
		return utils.ResultErr[SpriteSheet](fmt.Errorf("Error using texture"))
	}
	cols := int8(texture.Width / int32(spriteSize.X))
	rows := int8(texture.Height / int32(spriteSize.Y))
	if to == 0 {
		to = cols * rows
	}
	return utils.ResultOk(SpriteSheet{
		Texture2D:  texture,
		totalCols:  cols,
		totalRows:  rows,
		spriteSize: spriteSize,
		offset:     offset,
		from:       from,
		to:         to,
	})
}

func NewSpriteSheet(path string, spriteSize base.Vec[float32], offset base.Vec[float32], from, to int8) utils.Result[SpriteSheet] {
	texture := rl.LoadTexture(path)
	if texture.Width == 0 {
		return utils.ResultErr[SpriteSheet](fmt.Errorf("Error loading texture '%v' ", path))
	}
	cols := int8(texture.Width / int32(spriteSize.X))
	rows := int8(texture.Height / int32(spriteSize.Y))
	if to == 0 {
		to = cols * rows
	}
	return utils.ResultOk(SpriteSheet{
		Texture2D:  texture,
		totalCols:  cols,
		totalRows:  rows,
		spriteSize: spriteSize,
		offset:     offset,
		from:       from,
		to:         to,
	})
}

func (s SpriteSheet) GetTexture() rl.Texture2D {
	return s.Texture2D
}

func (s SpriteSheet) GetRectangle(currentSprite int8) rl.Rectangle {
	col := float32(currentSprite%s.totalCols) * (s.spriteSize.X + float32(s.offset.X)*2)
	row := float32(currentSprite/s.totalRows) * (s.spriteSize.Y + float32(s.offset.Y)*2)
	return rl.Rectangle{
		X:      col,
		Y:      row,
		Width:  s.spriteSize.X,
		Height: s.spriteSize.Y,
	}
}
