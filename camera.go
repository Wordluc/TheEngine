package main

import (
	"game/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	sizeView   base.Vec[int32]
	resolution base.Vec[int32]
	camera     *rl.Camera2D
	view       rl.RenderTexture2D
}

func NewCamera(sizeView base.Vec[int32], resolution base.Vec[int32]) Camera {
	return Camera{
		sizeView:   sizeView,
		resolution: resolution,
		camera: &rl.Camera2D{
			Offset: rl.Vector2{X: float32(resolution.X / 2), Y: float32(resolution.Y / 2)},
			Target: rl.Vector2{X: float32(resolution.X / 2), Y: float32(resolution.Y / 2)},
			Zoom:   float32(resolution.X / sizeView.X),
		},
		view: rl.LoadRenderTexture(sizeView.X, sizeView.Y),
	}
}

func (c *Camera) SetResolution(w, h int32) {
	c.resolution = base.Vec[int32]{X: w, Y: h}
	c.camera.Zoom = float32(c.resolution.X / c.sizeView.X)
	c.camera.Offset = rl.Vector2{X: float32(c.resolution.X) / 2, Y: float32(c.resolution.Y) / 2}
	rl.UnloadRenderTexture(c.view)
	c.view = rl.LoadRenderTexture(w, h)
}

func (c *Camera) StartRendering(pos base.Vec[float32]) {
	c.camera.Target = rl.Vector2{X: pos.X, Y: pos.Y}
	rl.BeginTextureMode(c.view)
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(*c.camera)
}

func (c *Camera) StopRendering() {
	rl.EndMode2D()
	rl.EndTextureMode()

	rl.BeginDrawing()
	scale := float32(rl.GetScreenWidth()) / float32(c.resolution.X)

	offsetX := (float32(rl.GetScreenWidth()) - float32(c.resolution.X)*scale) / 2
	offsetY := (float32(rl.GetScreenHeight()) - float32(c.resolution.Y)*scale) / 2

	rl.DrawTexturePro(
		c.view.Texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(c.resolution.X), Height: -float32(c.resolution.Y)},
		rl.Rectangle{X: offsetX, Y: offsetY, Width: float32(c.resolution.X) * scale, Height: float32(c.resolution.Y) * scale},
		rl.Vector2{},
		0.0,
		rl.White,
	)
	rl.EndDrawing()
}
