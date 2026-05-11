package main

import (
	b "game/object/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetVecForKeyboard() b.Vec {
	if rl.IsKeyDown(rl.KeyD) {
		return b.NewVec(1, 0)
	} else if rl.IsKeyDown(rl.KeyA) {
		return b.NewVec(-1, 0)
	} else if rl.IsKeyDown(rl.KeyW) {
		return b.NewVec(0, -1)
	} else if rl.IsKeyDown(rl.KeyS) {
		return b.NewVec(0, 1)
	}
	return b.Vec{}
}
