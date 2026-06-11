package utils

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawHitbox(o base.Object) {
	hitBox := o.GetHitbox()
	if !hitBox.IsActive {
		return
	}
	pos := o.GetPos()
	pos.Add(*hitBox.Pos)
	var currentV base.Vec[float32] = *pos.Clone()
	var nextV base.Vec[float32]
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(int32(pos.X)), Y: float32(int32(pos.Y)), Width: float32(hitBox.GetOuterBox().X), Height: float32(hitBox.GetOuterBox().Y)}, 1, rl.Blue)
	for _, v := range hitBox.GetVertex() {
		nextV = *v.Clone().Add(pos)
		rl.DrawLine(int32(currentV.X), int32(currentV.Y), int32(nextV.X), int32(nextV.Y), rl.Blue)
		currentV = nextV
	}
}
