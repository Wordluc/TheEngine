package main

import (
	"fmt"
	"game/core"
	"game/core/base"
	"math/rand"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	W_RESOLUTION = 1280
	H_RESOLUTION = 720
	W_WINDOW     = 1280
	H_WINDOW     = 720
)

type character struct {
	core.Ball
	speed base.Vec[float32]
}

const m_speed = 1
const size_ball = 3

func main() {
	balls := []*character{}
	rl.InitWindow(W_WINDOW, H_WINDOW, "Ciao")
	ball := new(character{
		Ball:  core.NewBall(size_ball),
		speed: getRandomSpeed(m_speed, m_speed),
	})
	ball.MoveTo(base.NewVec[float32](1260, 40))
	balls = append(balls, ball)

	wall := core.NewSquare(20)

	camera := core.NewCamera(base.Vec[int32]{
		X: W_RESOLUTION,
		Y: H_RESOLUTION,
	}, base.Vec[int32]{
		X: W_WINDOW,
		Y: H_WINDOW,
	})

	rl.SetTargetFPS(60)
	for {
		if rl.WindowShouldClose() {
			return
		}

		camera.StartRendering(base.CastVec[int32, float32](base.Vec[int32]{}))

		core.DrawHitbox(ball)
		if rl.IsKeyDown(rl.KeyN) {
			ball := new(character{
				Ball:  core.NewBall(size_ball),
				speed: getRandomSpeed(m_speed, m_speed),
			})
			p := rl.GetMousePosition()
			ball.MoveTo(base.NewVec(p.X, p.Y))
			balls = append(balls, ball)
		}
		for i := range balls {
			balls[i].MoveBy(balls[i].speed)
			checkBoundaris(balls[i])
			balls[i].Draw()
		}
		for i := range balls {
			for j := range balls {
				if i == j {
					continue
				}
				if collide(*balls[i], *balls[j]) {
					invert(&balls[i].speed)
					invert(&balls[j].speed)
				}

			}
			balls[i].MoveBy(balls[i].speed)
			checkBoundaris(balls[i])
			balls[i].Draw()
		}
		rl.DrawText(fmt.Sprintf("FPS: %v\n", 1/rl.GetFrameTime()), 10, 40, 20, rl.Red)
		rl.DrawText(fmt.Sprintf("ELEMENT: %v", len(balls)), 10, 80, 20, rl.Red)

		wall.Draw()
		camera.StopRendering()
	}
}

func invert(v *base.Vec[float32]) {
	v.X *= -1
	v.Y *= -1
}
func collide(a, b character) bool {
	xa, ya := a.GetPos().Get()
	xb, yb := b.GetPos().Get()
	dx := xa - xb
	dy := ya - yb
	return dx*dx+dy*dy <= size_ball*size_ball*size_ball*size_ball
}
func getRandomSpeed(xMax, yMax float32) (res base.Vec[float32]) {
	res.X = (rand.Float32() * xMax) - xMax/2
	res.Y = (rand.Float32() * yMax) - yMax/2
	return res
}

func checkBoundaris(c *character) {
	ps := c.GetPos()
	if ps.X < size_ball {
		ps.X = size_ball
		c.speed.X *= -1
	}
	if ps.Y < size_ball {
		ps.Y = size_ball
		c.speed.Y *= -1
	}
	if ps.X > 1280-size_ball {
		ps.X = 1280 - size_ball
		c.speed.X *= -1
	}
	if ps.Y > 720-size_ball {
		ps.Y = 720 - size_ball
		c.speed.Y *= -1
	}
}
