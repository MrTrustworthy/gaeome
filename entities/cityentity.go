package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var CitySpritesheet *common.Spritesheet

func LoadSpritesheet() {
	CitySpritesheet = common.NewSpritesheetWithBorderFromFile("textures/citySheet.png", 16, 16, 1, 1)

}

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewCity(x, y, i, j, randIdx int) *City {
	city := &City{BasicEntity: ecs.NewBasic()}
	city.SpaceComponent.Position = engo.Point{
		X: float32(((x+1)*64)+8) + float32(i*16),
		Y: float32((y+1)*64) + float32(j*16),
	}
	city.RenderComponent.Drawable = CitySpritesheet.Cell(CitySpriteMap[randIdx][i+3*j])
	city.RenderComponent.SetZIndex(1)

	return city
}

var CitySpriteMap = [...][12]int{
	{99, 100, 101,
		454, 269, 455,
		415, 195, 416,
		452, 306, 453,
	},
	{99, 100, 101,
		268, 269, 270,
		268, 269, 270,
		305, 306, 307,
	},
	{75, 76, 77,
		446, 261, 447,
		446, 261, 447,
		444, 298, 445,
	},
	{75, 76, 77,
		407, 187, 408,
		407, 187, 408,
		444, 298, 445,
	},
	{75, 76, 77,
		186, 150, 188,
		186, 150, 188,
		297, 191, 299,
	},
	{83, 84, 85,
		413, 228, 414,
		411, 191, 412,
		448, 302, 449,
	},
	{83, 84, 85,
		227, 228, 229,
		190, 191, 192,
		301, 302, 303,
	},
	{91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 946, 947,
	},
	{91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 803, 947,
	},
	{91, 92, 93,
		238, 239, 240,
		238, 239, 240,
		312, 313, 314,
	},
}
