package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type LevelTile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func LoadTilemap(tileURL string, world *ecs.World) engo.AABB {
	resource, err := engo.Files.Resource(tileURL)
	if err != nil {
		panic(err)
	}

	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level
	tiles := createTiles(levelData)
	addToRenderSystem(tiles, world)

	return levelData.Bounds()
}

func createTiles(levelData *common.Level) (tiles []*LevelTile) {
	tiles = make([]*LevelTile, 0)
	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &LevelTile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return
}

func addToRenderSystem(tiles []*LevelTile, world *ecs.World) {
	// add the tiles to the RenderSystem
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, tile := range tiles {
				sys.Add(&tile.BasicEntity, &tile.RenderComponent, &tile.SpaceComponent)
			}
		}
	}
}
