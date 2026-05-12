package main

import (
	"game/base"
	b "game/base"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	W_WINDOW = 1200
	H_WINDOW = 700
)

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")

	ball := NewBall(10)
	ball.MoveTo(b.NewVec[float32](40, 40))

	wall := NewSquare(20)

	camera := NewCamera(base.Vec[int32]{
		X: 320,
		Y: 180,
	}, base.Vec[int32]{
		X: 320,
		Y: 180,
	})
	rl.SetTargetFPS(1000)
	for {
		if rl.WindowShouldClose() {
			return
		}

		if rl.IsKeyPressed(rl.KeyF1) {
			camera.SetResolution(320, 180)
		}

		if rl.IsKeyPressed(rl.KeyF2) {
			camera.SetResolution(640, 360)
		}

		if rl.IsKeyPressed(rl.KeyF11) {
			rl.ToggleFullscreen()
		}

		camera.StartRendering(b.CastVec[int32, float32](b.Vec[int32]{}))

		ball.Draw()
		wall.Draw()
		DrawHitbox(&ball)

		camera.StopRendering()
		ball.MoveBy(GetVecForKeyboard())

	}
}
