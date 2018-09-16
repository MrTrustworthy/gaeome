package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/gaeome/entities"
	"github.com/MrTrustworthy/gaeome/systems"
	"github.com/MrTrustworthy/gaeome/ui"
)

type BaseScene struct{}

// Type uniquely defines your game type
func (*BaseScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*BaseScene) Preload() {
	engo.Files.Load("textures/city.png", "tilemap/TrafficMap.tmx", "textures/citySheet.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (scene *BaseScene) Setup(updater engo.Updater) {

	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	world, _ := updater.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})

	scene.SetupInputs(world)

	ui.LoadUI(world)
	levelBounds := entities.LoadTilemap("tilemap/TrafficMap.tmx", world)

	// bounds the camera to the map, but only works if the camera zoom is 1
	//common.CameraBounds = engo.AABB{
	//	Min: engo.Point{X: levelBounds.Min.X + (1200)/2, Y: levelBounds.Min.Y + (800)/2},
	//	Max: engo.Point{X: levelBounds.Max.X - (1200)/2, Y: levelBounds.Max.Y - (800)/2},
	//}
	common.CameraBounds = levelBounds

	// add this last so all dependencies on other systems are resolved
	world.AddSystem(&systems.CityBuildingSystem{})
}

func (*BaseScene) SetupInputs(world *ecs.World) {
	engo.Input.RegisterButton("AddCity", engo.KeyC)

	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(common.NewKeyboardScroller(400, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	world.AddSystem(&common.EdgeScroller{400, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})
}

func main() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          1200,
		Height:         800,
		StandardInputs: true,
	}
	engo.Run(opts, &BaseScene{})
}
