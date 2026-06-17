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

var SPRITE_SIZE_TERRAIN = base.Vec[float32]{X: 160, Y: 162}
var SPRITE_SIZE_CHARACTER = base.Vec[float32]{X: 156.5, Y: 156.5}

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	characterSprites, err := utils.ResultsMap(map[string]utils.Result[core.SpriteSheet]{
		"stop": core.NewSpriteSheet("demos/sprite_2/walk.png", SPRITE_SIZE_CHARACTER, base.Vec[float32]{}, 0, 1),
		"walk": core.NewSpriteSheet("demos/sprite_2/walk.png", SPRITE_SIZE_CHARACTER, base.Vec[float32]{}, 0, 0),
	})
	if err != nil {
		panic(err)
	}
	character := core.NewSprite(characterSprites)
	character.Hitbox = core.GetRectangleHitbox(SPRITE_SIZE_CHARACTER.X, SPRITE_SIZE_CHARACTER.Y)
	character.SpeedSpriteLoop = 6
	character.SetModifier(new(base.NewRigidBody(true, false, 50)))
	terrainSprite, err := core.NewSpriteSheet("demos/sprite_2/terrain.png", SPRITE_SIZE_TERRAIN, base.Vec[float32]{X: 1}, 0, 0).Open()
	if err != nil {
		panic(err)
	}
	terrain := core.NewMultiSprite(terrainSprite, base.NewVec[int32](5, 1))
	terrain.SelectedSprite = 13
	terrain.SetModifier(new(base.NewRigidBody(true, true, 0)))
	terrain.MoveTo(base.Vec[float32]{X: 0, Y: 500})
	camera := core.NewCamera(base.Vec[int32]{
		X: W_RESOLUTION,
		Y: H_RESOLUTION,
	}, base.Vec[int32]{
		X: W_WINDOW,
		Y: H_WINDOW,
	})
	rl.SetTargetFPS(30)
	quad := base.NewQuadTree(base.NewVec[float32](0, 0), base.NewVec[float32](W_WINDOW, H_WINDOW), nil)

	for {
		if rl.WindowShouldClose() {
			return
		}
		quad.Clear()
		quad.Insert(&terrain)
		quad.Insert(&character)
		base.UseModifierRef(&character, func(o base.Modifier) {
			switch r := o.(type) {
			case *base.RigidBody:
				r.ApplyAcceleration(utils.GetVecForKeyboard(100))
				if r.GetVelocity().X != 0 {
					character.ChanceSpriteSheet("walk")
					if r.GetVelocity().X < 0 {
						character.FlippedX = true
					} else {
						character.FlippedX = false
					}
				} else {
					character.ChanceSpriteSheet("stop")
				}
				r.ApplyAcceleration(base.Vec[float32]{Y: 8})
				r.Integrate(rl.GetFrameTime())
			}

		})
		quad.Foreach(func(elements []base.QuadTreeElement) {
			for i := range elements {
				for j := range elements {
					if i == j {
						continue
					}
					a, okA := elements[i].(base.Object)
					b, okB := elements[j].(base.Object)
					if !okA || !okB {
						continue
					}
					ra := base.GetModifierRef[*base.RigidBody](a)
					rb := base.GetModifierRef[*base.RigidBody](b)
					if ra == nil || rb == nil {
						continue
					}
					ra.Collide(rb)
				}
			}
		})

		camera.StartRendering(base.CastVec[int32, float32](base.Vec[int32]{}))

		character.Draw()
		terrain.Draw()
		//	utils.DrawHitbox(&terrain)
		//		utils.DrawHitbox(&character)

		camera.StopRendering()
		character.SpriteLoop()
	}
}
