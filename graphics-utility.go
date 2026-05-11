package main

import (
	"game/object/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawHitbox(o base.Object) {
	size := o.GetHitbox()
	pos := o.GetPos()
	x, y := pos.Get()
	w, h := size.Get()
	rl.DrawRectangleLines(x, y-1, w+1, h+1, rl.Red)
}
