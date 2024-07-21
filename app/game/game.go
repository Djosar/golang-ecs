package game

import (
	"log"
	"reflect"

	"github.com/Djosar/kro-ecs/app/factories"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game represents the main game structure. It holds the registry of all entities
// and systems, and the player entity.
type Game struct {
	Registry     *core.Registry
	PlayerEntity core.Entity
}

// NewGame initializes and returns a new Game instance. It sets up the registry,
// adds systems to it, and creates the player entity.
//
// Returns:
//
//	*Game: A pointer to the newly created Game instance.
//	error: An error if the player entity cannot be created.
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
	renderer := g.Registry.GetSystem(reflect.TypeOf(&systems.RenderSystem{}))
	renderer.(*systems.RenderSystem).Screen = screen
	renderer.Update(g.Registry)
}
