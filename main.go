package main

import (
	"game/base"
	b "game/base"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	W_RESOLUTION = 640
	H_RESOLUTION = 360
	W_WINDOW     = 1280
	H_WINDOW     = 720
)

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")

	ball := NewBall(10)
	ball.MoveTo(b.NewVec[float32](1260, 40))

	wall := NewSquare(20)

	camera := NewCamera(base.Vec[int32]{
		X: W_RESOLUTION,
		Y: H_RESOLUTION,
	}, base.Vec[int32]{
		X: W_WINDOW,
		Y: H_WINDOW,
	})

	rl.SetTargetFPS(30)
	for {
		if rl.WindowShouldClose() {
			return
		}

		if rl.IsKeyPressed(rl.KeyF1) {
			camera.SetResolution(W_RESOLUTION, H_RESOLUTION)
		}

		if rl.IsKeyPressed(rl.KeyF2) {
			camera.SetResolution(W_RESOLUTION*2, H_RESOLUTION*2)
		}

		if rl.IsKeyPressed(rl.KeyR) {
			rl.ToggleFullscreen()
			camera.SetScreenSize(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
		}

		camera.StartRendering(b.CastVec[int32, float32](b.Vec[int32]{}))

		ball.Draw()
		wall.Draw()
		DrawHitbox(&ball)

		camera.StopRendering()
		ball.MoveBy(GetVecForKeyboard(10))

	}
}
