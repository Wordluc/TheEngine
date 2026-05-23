package main

import (
	"game/core"
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W_RESOLUTION = 640
	H_RESOLUTION = 360
	W_WINDOW     = 1280
	H_WINDOW     = 720
)

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	character := base.NewNode(base.Vec[float32]{})
	sp, err := core.NewSprite(base.Vec[float32]{X: 156.5, Y: 156.5}, map[string]string{
		"walk": "demos/sprite/walk.png",
	})
	sp.SpeedSpriteLoop = 6
	if err != nil {
		panic(err)
	}
	character.AddObject(&sp)
	camera := core.NewCamera(base.Vec[int32]{
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
		camera.StartRendering(base.CastVec[int32, float32](base.Vec[int32]{}))
		_ = character.ForEachObjects(
			func(o base.Object) error {
				o.Draw()
				return nil
			},
		)

		character.MoveBy(core.GetVecForKeyboard(100))
		camera.StopRendering()
		sp.SpriteLoop()
	}
}
