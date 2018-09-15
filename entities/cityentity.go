package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (city *City) Initialize(position *engo.Point) {
	city.SpaceComponent = common.SpaceComponent{
		Position: *position,
		Width:    30,
		Height:   64,
	}

	texture, err := common.LoadedSprite("textures/city.png")
	if err != nil {
		panic(err)
	}

	city.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{0.1, 0.1},
	}
}

func (city *City) LoadInto(world *ecs.World) {

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
		}
	}
}
