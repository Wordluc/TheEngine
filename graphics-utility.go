package main

import (
	"game/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawHitbox(o base.Object) {
	size := o.GetHitbox()
	pos := o.GetPos()
	x, y := pos.Get()
	w, h := size.Get()
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(w), Height: float32(h)}, 1, rl.Red)
}
