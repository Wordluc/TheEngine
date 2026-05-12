package main

import (
	b "game/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetVecForKeyboard() b.Vec[int32] {
	if rl.IsKeyDown(rl.KeyD) {
		return b.NewVec[int32](1, 0)
	} else if rl.IsKeyDown(rl.KeyA) {
		return b.NewVec[int32](-1, 0)
	} else if rl.IsKeyDown(rl.KeyW) {
		return b.NewVec[int32](0, -1)
	} else if rl.IsKeyDown(rl.KeyS) {
		return b.NewVec[int32](0, 1)
	}
	return b.Vec[int32]{}
}
