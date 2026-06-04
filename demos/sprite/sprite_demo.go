package main

import (
	"game/core"
	"game/core/base"
	"game/core/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W_RESOLUTION = 640
	H_RESOLUTION = 360
	W_WINDOW     = 1280
	H_WINDOW     = 720
)

var SPRITE_SIZE = base.Vec[float32]{X: 156.5, Y: 156.5}

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	character := base.NewNode(base.Vec[float32]{})
	characterSprites, err := utils.ResultsMap(map[string]utils.Result[core.SpriteSheet]{
		"walk": core.NewSpriteSheet("demos/sprite/walk.png", SPRITE_SIZE, 0),
	})
	if err != nil {
		panic(err)
	}
	sp := core.NewSprite(characterSprites)
	sp.SpeedSpriteLoop = 6
	character.AddObject(&sp)
	terrainSprite, err := core.NewSpriteSheet("demos/sprite/terrain.png", SPRITE_SIZE, 2).Open()
	if err != nil {
		panic(err)
	}
	terrain := core.NewMultiSprite(terrainSprite)
	terrain.BlockSize.X = 5
	terrain.BlockSize.Y = 1
	terrain.SelectedSprite = 13
	terrain.SpriteSheetOffset = 2
	terrain.MoveTo(base.Vec[float32]{X: 1, Y: 720 - 156.5})
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
				if d, ok := o.(base.Drawable); ok {
					d.Draw()
				}
				return nil
			},
		)
		terrain.Draw()

		character.MoveBy(utils.GetVecForKeyboard(100))
		camera.StopRendering()
		sp.SpriteLoop()
	}
}
