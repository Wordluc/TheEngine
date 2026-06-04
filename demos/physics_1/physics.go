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
	rBall := new(base.NewRigidBody(true, false))
	ball.SetModifier(rBall)
	terrain := core.NewRectangle(100, 5)
	terrain.MoveTo(base.NewVec[float32](0, 100))
	rTerrain := new(base.NewRigidBody(true, true))
	terrain.SetModifier(rTerrain)

	block := core.NewRectangle(20, 20)
	block.MoveTo(base.NewVec[float32](40, 50))
	rBlock := new(base.NewRigidBody(true, false))
	block.SetModifier(rBlock)

	block1 := core.NewRectangle(20, 20)
	block1.MoveTo(base.NewVec[float32](40, 50))
	rBlock1 := new(base.NewRigidBody(true, false))
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
		ball.MoveBy(base.NewVec(0, 20*rl.GetFrameTime()))
		block.MoveBy(base.NewVec(0, 20*rl.GetFrameTime()))
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
		quad.Query(func(elements []base.QuadTreeElement[float32]) {
			for i := range elements {
				for j := range elements {
					if i == j {
						continue
					}
					if elements[i] == elements[j] { //See why this happen!!! Error
						continue
					}
					a, okA := elements[i].(base.Object)
					b, okB := elements[j].(base.Object)
					if !okA || !okB {
						continue
					}
					a.GetModifiers()[0].(*base.RigidBody).Collide(b.GetModifiers()[0].(*base.RigidBody))
				}
			}
		})
		quad.DrawBorder()
		camera.StopRendering()
		if rl.IsKeyDown(rl.KeyS) {
			terrain.MoveBy(base.NewVec[float32](0, 2))
		}
		if rl.IsKeyDown(rl.KeyD) {
			terrain.MoveBy(base.NewVec[float32](2, 0))
		}
		if rl.IsKeyDown(rl.KeyA) {
			terrain.MoveBy(base.NewVec[float32](-2, 0))
		}
		if rl.IsKeyDown(rl.KeyW) {
			terrain.MoveBy(base.NewVec[float32](0, -2))
		}
	}
}
