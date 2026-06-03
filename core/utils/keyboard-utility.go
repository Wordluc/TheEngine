package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetVecForKeyboard(v int) (res base.Vec[float32]) {
	unit := float32(v) * rl.GetFrameTime()
	if rl.IsKeyDown(rl.KeyD) {
		res.AddScalars(unit, 0)
	} else if rl.IsKeyDown(rl.KeyA) {
		res.AddScalars(-unit, 0)
	}
	if rl.IsKeyDown(rl.KeyW) {
		res.AddScalars(0, -unit)
	} else if rl.IsKeyDown(rl.KeyS) {
		res.AddScalars(0, unit)
	}
	return res
}
