package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/gaeome/entities"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type CityBuildingSystem struct {
	mouseTracker MouseTracker
	world *ecs.World
}

func (cbs *CityBuildingSystem) Update(dt float32){
	if engo.Input.Button("AddCity").JustPressed() {
		cbs.AddCity()
	}

}

func (cbs *CityBuildingSystem) AddCity() {
	city := entities.City{BasicEntity: ecs.NewBasic()}
	city.Initialize(&engo.Point{cbs.mouseTracker.MouseX, cbs.mouseTracker.MouseY})
	city.LoadInto(cbs.world)

	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis: common.XAxis,
		Value: cbs.mouseTracker.MouseX,
		Incremental: false,
	})
	engo.Mailbox.Dispatch(common.CameraMessage{
		Axis: common.YAxis,
		Value: cbs.mouseTracker.MouseY,
		Incremental: false,
	})
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
}

func (cbs *CityBuildingSystem) Remove(entity ecs.BasicEntity) {

}
