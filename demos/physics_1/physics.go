package main

import (
	"game/core"
	"game/core/base"
	"game/core/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W_RESOLUTION = 1280
	H_RESOLUTION = 720
	W_WINDOW     = 1280
	H_WINDOW     = 720
)

var SPRITE_SIZE = base.Vec[float32]{X: 156.5, Y: 156.5}

func main() {
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	ball := core.NewRectangle(20, 20)
	rBall := new(base.NewRigidBody(true, false, 5))
	ball.SetModifier(rBall)
	rBall.Id = "ball"
	terrain := core.NewRectangle(500, 5)
	terrain.MoveTo(base.NewVec[float32](0, 200))
	rTerrain := new(base.NewRigidBody(true, true, 0))
	terrain.SetModifier(rTerrain)

	block := core.NewRectangle(150, 20)
	block.MoveTo(base.NewVec[float32](150, 180))
	rBlock := new(base.NewRigidBody(true, true, 30))
	rBlock.Friction = -1
	block.SetModifier(rBlock)

	block1 := core.NewRectangle(20, 20)
	block1.MoveTo(base.NewVec[float32](40, 50))
	rBlock1 := new(base.NewRigidBody(true, false, 30))
	block1.SetModifier(rBlock1)
	rBlock1.Id = "block"

	camera := core.NewCamera(base.Vec[int32]{
		X: W_RESOLUTION,
		Y: H_RESOLUTION,
	}, base.Vec[int32]{
		X: W_WINDOW,
		Y: H_WINDOW,
	})
	var FPS int32 = 30
	rl.SetTargetFPS(FPS)
	quad := base.NewQuadTree(base.NewVec[float32](0, 0), base.NewVec[float32](W_WINDOW, H_WINDOW), nil)
	for {
		if rl.WindowShouldClose() {
			return
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			FPS -= 30
			rl.SetTargetFPS(FPS)
			println("FPS ", FPS)
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			FPS += 30
			rl.SetTargetFPS(FPS)
			println("FPS ", FPS)
		}
		quad.Clear()
		quad.Insert(&ball)
		quad.Insert(&terrain)
		quad.Insert(&block)
		quad.Insert(&block1)

		camera.StartRendering(base.CastVec[int32, float32](base.Vec[int32]{}))
		ball.Draw()
		terrain.Draw()
		block.Draw()
		block1.Draw()

		utils.DrawHitbox(&ball)
		utils.DrawHitbox(&block)
		utils.DrawHitbox(&block1)
		utils.DrawHitbox(&terrain)

		r := base.GetModifierRef[*base.RigidBody](&ball)
		if rl.IsKeyPressed(rl.KeyUp) {
			r.Mass += 10
			println("Mass ", r.Mass)
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			r.Mass -= 10
			println("Mass ", r.Mass)
		}
		isTouchingDown := r.Collision.CheckIf(func(cd base.CollisionDetail) bool { return cd.Y < 0 })
		if rl.IsKeyDown(rl.KeyS) {
			r.ApplyAcceleration(base.NewVec[float32](0, 10))
		}
		if isTouchingDown {
			if rl.IsKeyPressed(rl.KeyW) {
				r.ApplyImpulse(base.NewVec[float32](0, -40))
			}
		}
		if rl.IsKeyDown(rl.KeyD) {
			var speed float32 = 20
			if !isTouchingDown {
				speed = 5
			}
			r.ApplyAcceleration(base.NewVec(speed, 0))
		}
		if rl.IsKeyDown(rl.KeyA) {
			var speed float32 = 20
			if !isTouchingDown {
				speed = 5
			}
			r.ApplyAcceleration(base.NewVec(-speed, 0))
		}

		base.UseModifierRef(&ball, func(rb *base.RigidBody) {
			rb.Touch()
			rb.ApplyAcceleration(base.NewVec[float32](0, 8))
			rb.Integrate(rl.GetFrameTime())
			rb.GetVelocity().CapAt(base.Vec[float32]{X: 20, Y: 20})
		})

		base.UseModifierRef(&block1, func(rb *base.RigidBody) {
			rb.Touch()
			rb.ApplyAcceleration(base.NewVec[float32](0, 8))
			rb.Integrate(rl.GetFrameTime())
			rb.GetVelocity().CapAt(base.Vec[float32]{X: 20, Y: 20})
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
		quad.DrawBorder()
		camera.StopRendering()

	}
}
