package core

import (
	"github.com/Wordluc/TheEngine/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	base.ObjectBase
	resolution base.Vec[int32]
	screenSize base.Vec[int32]
	camera     *rl.Camera2D
	view       rl.RenderTexture2D
}

func NewCamera(resolution base.Vec[int32], screenSize base.Vec[int32]) Camera {
	return Camera{
		resolution: resolution,
		screenSize: screenSize,
		camera: &rl.Camera2D{
			Offset: rl.Vector2{},
			Target: rl.Vector2{},
			Zoom:   float32(resolution.X) / float32(screenSize.X),
		},
		view: rl.LoadRenderTexture(resolution.X, resolution.Y),
	}
}

func (c *Camera) SetResolution(w, h int32) {
	c.resolution = base.Vec[int32]{X: w, Y: h}
	c.camera.Zoom = float32(c.resolution.X) / float32(c.screenSize.X)
	rl.UnloadRenderTexture(c.view)
	c.view = rl.LoadRenderTexture(w, h)
}

func (c *Camera) SetScreenSize(w, h int32) {
	c.screenSize = base.Vec[int32]{X: w, Y: h}
	c.camera.Zoom = float32(c.resolution.X) / float32(c.screenSize.X)
}

func (c *Camera) StartRendering() {
	pos := c.GetPos()
	c.camera.Target = rl.Vector2{X: pos.X, Y: pos.Y}
	rl.BeginTextureMode(c.view)
	rl.ClearBackground(rl.White)
	rl.BeginMode2D(*c.camera)
}

func (c *Camera) StopRendering() {
	rl.EndMode2D()
	rl.EndTextureMode()

	rl.BeginDrawing()

	rl.DrawTexturePro(
		c.view.Texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(c.resolution.X), Height: -float32(c.resolution.Y)},
		rl.Rectangle{X: 0, Y: 0, Width: float32(c.screenSize.X), Height: -float32(c.screenSize.Y)},
		rl.Vector2{},
		c.Angle(),
		rl.White,
	)
	rl.EndDrawing()
}
