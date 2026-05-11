package main

import (
	b "game/object/base"

	"github.com/gen2brain/raylib-go/raylib"
)

const W_WINDOW = 1200
const H_WINDOW = 700

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	//	rl.ToggleFullscreen()

	ball := NewBall(10)
	ball.MoveTo(b.NewVec(40, 40))
	target := rl.LoadRenderTexture(W_WINDOW, H_WINDOW)
	for {
		if rl.WindowShouldClose() {
			return
		}
		{
			rl.BeginTextureMode(target)
			rl.ClearBackground(rl.White)
			ball.Draw()
			DrawHitbox(&ball)
			rl.EndTextureMode()
		}

		rl.BeginDrawing()
		rl.DrawTexturePro(
			target.Texture,
			rl.Rectangle{X: 0, Y: 0, Width: W_WINDOW, Height: -H_WINDOW},
			rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: float32(rl.GetScreenHeight())},
			rl.Vector2{X: 0, Y: 0},
			0.0,
			rl.White,
		)
		ball.MoveBy(GetVecForKeyboard())
		rl.EndDrawing()
	}
}
