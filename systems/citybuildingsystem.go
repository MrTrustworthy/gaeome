package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"errors"
	"github.com/MrTrustworthy/gaeome/entities"
	"math/rand"
	"time"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type CityBuildingSystem struct {
	mouseTracker MouseTracker
	world        *ecs.World
	usedTiles    []int
	elapsed      float32
	built        int
}

func (cbs *CityBuildingSystem) Update(dt float32) {
	cbs.elapsed += dt
	if cbs.elapsed >= 2.0/(float32(cbs.built)+1.0) {
		cbs.generateCity()
		cbs.elapsed = 0
		cbs.built++
	}

}

func (cbs *CityBuildingSystem) New(world *ecs.World) {

	cbs.world = world

	cbs.mouseTracker.BasicEntity = ecs.NewBasic()
	cbs.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cbs.mouseTracker.BasicEntity, &cbs.mouseTracker.MouseComponent, nil, nil)

		}
	}
	entities.LoadSpritesheet()
	rand.Seed(time.Now().UnixNano())

}

func (cbs *CityBuildingSystem) generateCity() {

	x, y, err := cbs.reserveFreeTile()
	if err != nil {
		return
	}
	randIdx := rand.Intn(len(entities.CitySpriteMap))

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			city := entities.NewCity(x, y, i, j, randIdx)
			cbs.LoadIntoWorld(city)
		}
	}

	engo.Mailbox.Dispatch(HUDTextMessage{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{X: float32((x + 1) * 64), Y: float32((y + 1) * 64)},
			Width:    64,
			Height:   64,
		},
		MouseComponent: common.MouseComponent{},
		Line:           "Town, just built, generates $100 per day",
	})
}

func (cbs *CityBuildingSystem) LoadIntoWorld(city *entities.City) {
	for _, system := range cbs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
		}
	}
}

func (cbs *CityBuildingSystem) reserveFreeTile() (x, y int, err error) {
	x = rand.Intn(18)
	y = rand.Intn(18)
	t := x + y*18
	err = nil

	if len(cbs.usedTiles) > 300 {
		err = errors.New("Too many cities")
		return
	}

	for cbs.isTileUsed(t) { // effectively "while current tile choice is used, assign new tile choice
		x = rand.Intn(18)
		y = rand.Intn(18)
		t = x + y*18
	}
	cbs.usedTiles = append(cbs.usedTiles, t)
	return
}

func (cbs *CityBuildingSystem) isTileUsed(tile int) bool {
	for _, t := range cbs.usedTiles {
		if tile == t {
			return true
		}
	}
	return false
}

func (cbs *CityBuildingSystem) Remove(entity ecs.BasicEntity) {}
