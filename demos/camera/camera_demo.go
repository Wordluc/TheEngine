package main

import (
	"game/core"
	"game/core/base"
	"game/core/utils"

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

	ball := core.NewCircle(10)
	ball.MoveTo(base.NewVec[float32](1260, 40))

	wall := core.NewRectangle(20, 20)

	camera := core.NewCamera(base.UVec[int32]{
		X: W_RESOLUTION,
		Y: H_RESOLUTION,
	}, base.UVec[int32]{
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

		camera.StartRendering(base.CastVec[int32, float32](base.UVec[int32]{}))

		ball.Draw()
		wall.Draw()
		utils.DrawHitbox(&ball)

		camera.StopRendering()
		ball.MoveBy(utils.GetVecForKeyboard(10))

	}
}
