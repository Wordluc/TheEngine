package utils

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawHitbox(o base.Object) {
	hitBox := o.GetHitbox()
	if hitBox.IsActive {
		return
	}
	pos := o.GetPos()
	x, y := pos.Get()
	w, h := hitBox.Box.Get()
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(int32(x)), Y: float32(int32(y)), Width: float32(w), Height: float32(h)}, 1, rl.Red)
}
