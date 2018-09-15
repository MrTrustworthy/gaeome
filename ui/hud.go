package ui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image"
	"image/color"
)

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewHUD() (hud HUD) {
	width, height := float32(400), float32(100)

	spaceComponent := common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - height},
		Width:    width,
		Height:   height,
	}

	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, int(width), int(height))
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)
	renderComponent := common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
	}
	renderComponent.SetShader(common.HUDShader)
	renderComponent.SetZIndex(1)

	hud = HUD{
		BasicEntity:     ecs.NewBasic(),
		SpaceComponent:  spaceComponent,
		RenderComponent: renderComponent,
	}

	return
}

func LoadUI(world *ecs.World) {
	hud := NewHUD()
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
		}
	}
}
