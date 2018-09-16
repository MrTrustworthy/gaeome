package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"image/color"
)

// Text is an entity containing text printed to the screen
type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
type HUDTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.MouseComponent
	Line string
}

const HUDTextMessageType string = "HUDTextMessage"

func (HUDTextMessage) Type() string {
	return HUDTextMessageType
}

// HUDTextSystem prints the text to our HUD based on the current state of the game
type HUDTextSystem struct {
	text, money Text
	entities    []HUDTextMessage
}

func (hts *HUDTextSystem) Add(entity *HUDTextMessage) {
	hts.entities = append(hts.entities, *entity)
}

// New is called when the system is added to the world.
// Adds text to our HUD that will update based on the state of the game.
func (hts *HUDTextSystem) New(world *ecs.World) {
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.Black,
		Size: 24,
	}
	fnt.CreatePreloaded()

	hts.text = Text{BasicEntity: ecs.NewBasic()}
	hts.text.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Nothing Selected!",
	}
	hts.text.SetShader(common.TextHUDShader)
	hts.text.RenderComponent.SetZIndex(1001)
	hts.text.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 200},
	}
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hts.text.BasicEntity, &hts.text.RenderComponent, &hts.text.SpaceComponent)
		}
	}

	hts.money = Text{BasicEntity: ecs.NewBasic()}
	hts.money.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "$0",
	}
	hts.money.SetShader(common.TextHUDShader)
	hts.money.RenderComponent.SetZIndex(1001)
	hts.money.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 40},
	}
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hts.money.BasicEntity, &hts.money.RenderComponent, &hts.money.SpaceComponent)
		}
	}

	hts.listenToMessages(world)
}

func (hts *HUDTextSystem) listenToMessages(world *ecs.World) {
	engo.Mailbox.Listen(HUDTextMessageType, func(message engo.Message) {
		msg, ok := message.(HUDTextMessage)
		if !ok {
			return
		}
		for _, system := range world.Systems() {
			switch sys := system.(type) {
			case *common.MouseSystem:
				sys.Add(&msg.BasicEntity, &msg.MouseComponent, &msg.SpaceComponent, nil)
			case *HUDTextSystem:
				sys.Add(&msg)
				txt := hts.text.RenderComponent.Drawable.(common.Text)
				txt.Text = msg.Line
				hts.text.RenderComponent.Drawable = txt
				fmt.Println("We have a total of ", len(sys.entities), "cities now")
			}
		}

	})
}

// Update is called each frame to update the system.
func (hts *HUDTextSystem) Update(dt float32) {}

// Remove takes an enitty out of the system.
// It does nothing as HUDTextSystem has no entities.
func (hts *HUDTextSystem) Remove(entity ecs.BasicEntity) {}
