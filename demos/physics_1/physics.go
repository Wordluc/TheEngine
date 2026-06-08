package main

import (
	"game/core"
	"game/core/base"

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
	ball := core.NewCircle(10)
	rBall := new(base.NewRigidBody(true, false, 5))
	ball.SetModifier(rBall)
	terrain := core.NewRectangle(500, 5)
	terrain.MoveTo(base.NewVec[float32](0, 200))
	rTerrain := new(base.NewRigidBody(true, true, 0))
	terrain.SetModifier(rTerrain)

	block := core.NewRectangle(20, 20)
	block.MoveTo(base.NewVec[float32](60, 180))
	rBlock := new(base.NewRigidBody(true, true, 30))
	block.SetModifier(rBlock)

	block1 := core.NewRectangle(20, 20)
	block1.MoveTo(base.NewVec[float32](40, 50))
	rBlock1 := new(base.NewRigidBody(true, true, 30))
	block1.SetModifier(rBlock1)

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
		base.UseModifierRef(&ball, func(rb *base.RigidBody) {
			rb.Touch()
			rb.ApplyAcceleration(base.NewVec[float32](0, 8))
			rb.Integrate(rl.GetFrameTime())
			rb.GetVelocity().CapAt(base.Vec[float32]{X: 10, Y: 20})
		})

		base.UseModifierRef(&block, func(rb *base.RigidBody) {
			rb.ApplyAcceleration(base.NewVec[float32](0, 8))
			rb.Integrate(rl.GetFrameTime())
			rb.GetVelocity().CapAt(base.Vec[float32]{X: 10, Y: 20})
		})
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

		r := base.GetModifierRef[*base.RigidBody](&ball)
		isTouchingDown := r.Collision.CheckIf(func(cd base.CollisionDetail) bool { return cd.Y < 0 })
		if rl.IsKeyDown(rl.KeyS) {
			r.ApplyAcceleration(base.NewVec[float32](0, 10))
		}
		if rl.IsKeyDown(rl.KeyD) {
			var speed float32 = 10
			if !isTouchingDown {
				speed = 5
			}
			r.ApplyAcceleration(base.NewVec(speed, 0))
		}
		if rl.IsKeyDown(rl.KeyA) {
			var speed float32 = 10
			if !isTouchingDown {
				speed = 5
			}
			r.ApplyAcceleration(base.NewVec(-speed, 0))
		}
		if isTouchingDown {
			if rl.IsKeyPressed(rl.KeyW) {
				r.ApplyImpulse(base.NewVec[float32](0, -7))
			}
		}
		if r.GetForce().X == 0 {
			func() {
				f := r.GetVelocity().Clone()
				f.Y = 0

				dt := rl.GetFrameTime()
				if dt == 0 {
					return
				}

				t := f.X / dt
				if isTouchingDown {
					f.X = t * -0.15
				} else {
					f.X = t * -0.1
				}
				r.ApplyAcceleration(*f)
			}()
		}

		quad.Query(func(elements []base.QuadTreeElement) {
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
