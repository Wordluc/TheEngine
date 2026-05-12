package main

import (
	b "game/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetVecForKeyboard(v int) b.Vec[float32] {
	unit := float32(v) * rl.GetFrameTime()
	if rl.IsKeyDown(rl.KeyD) {
		return b.NewVec(unit, 0)
	} else if rl.IsKeyDown(rl.KeyA) {
		return b.NewVec(-unit, 0)
	} else if rl.IsKeyDown(rl.KeyW) {
		return b.NewVec(0, -unit)
	} else if rl.IsKeyDown(rl.KeyS) {
		return b.NewVec(0, unit)
	}
	return b.Vec[float32]{}
}
