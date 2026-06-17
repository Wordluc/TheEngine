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
	character.Hitbox = core.GetRectangleHitbox(40, 0, 80, SPRITE_SIZE_CHARACTER.Y)
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
	camera.MoveBy(base.Vec[float32]{Y: -200})
	character.AddObject(&camera)
	ball := core.NewCircle(10)
	ball.SetModifier(new(base.NewRigidBody(true, false, 5)))
	ball.MoveBy(base.NewVec[float32](300, 0))
	rl.SetTargetFPS(30)
	quad := base.NewQuadTree(base.NewVec[float32](0, 0), base.NewVec[float32](W_WINDOW, H_WINDOW), nil)
	quad.DEBUG()

	for {
		if rl.WindowShouldClose() {
			return
		}
		quad.Clear()
		quad.Insert(&terrain)
		quad.Insert(&character)
		quad.Insert(&ball)
		base.UseModifierRef(&character, func(r *base.RigidBody) {
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
		})
		base.UseModifierRef(&ball, func(rb *base.RigidBody) {
			rb.Touch()
			rb.ApplyAcceleration(base.NewVec[float32](0, 8))
			rb.Integrate(rl.GetFrameTime())
			rb.GetVelocity().CapAt(base.Vec[float32]{X: 20, Y: 20})
		})
		if rl.IsKeyPressed(rl.KeyLeft) {
			camera.MoveBy(base.Vec[float32]{X: -10})
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			camera.MoveBy(base.Vec[float32]{X: 10})
		}
		if rl.IsKeyPressed(rl.KeyR) {
			camera.RotateBy(camera.Angle() + 10)
		}
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

		camera.StartRendering()

		character.Draw()
		terrain.Draw()
		ball.Draw()
		utils.DrawHitbox(&terrain)
		utils.DrawHitbox(&character)
		utils.DrawHitbox(&ball)

		camera.StopRendering()
	}
}
