package game

import (
	"log"
	"reflect"

	"github.com/Djosar/kro-ecs/app/factories"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Registry     *core.Registry
	PlayerEntity core.Entity
}

func NewGame() (*Game, error) {
	registry := core.NewRegistry()

	game := &Game{
		Registry: registry,
	}

	systems := []core.System{
		systems.NewRenderSystem(),
		systems.NewMovementSystem(),
		systems.NewInputSystem(),
		systems.NewAnimationSystem(),
	}
	for _, system := range systems {
		game.Registry.AddSystem(system)
	}

	entity, err := factories.PlayableCharacterFactory(game.Registry)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	game.PlayerEntity = entity

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) Update() error {
	excludedTypes := []reflect.Type{
		reflect.TypeOf(&systems.RenderSystem{}),
	}
	g.Registry.UpdateSystems(excludedTypes)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	renderer := g.Registry.GetSystem(reflect.TypeOf(&systems.RenderSystem{}))

	renderer.(*systems.RenderSystem).Screen = screen
	renderer.Update(g.Registry)
}
