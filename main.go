package main

import (
	"log"
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/systems"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	registry *core.Registry
}

func (g *Game) Update() error {
	excludedTypes := []reflect.Type{
		reflect.TypeOf(&systems.RenderSystem{}),
	}
	g.registry.UpdateSystems(excludedTypes)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	renderer := g.registry.GetSystem(reflect.TypeOf(&systems.RenderSystem{}))

	renderer.(*systems.RenderSystem).Screen = screen
	renderer.Update(g.registry)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		registry: core.NewRegistry(),
	}
	renderer := systems.NewRenderSystem()
	movement := systems.NewMovementSystem()
	input := systems.NewInputSystem()
	transform1 := &components.TransformComponent{
		Position: util.Coordinate[float32]{
			X: 0,
			Y: 0,
		},
		Speed: 1,
		Velocity: util.Velocity{
			DX: 0,
			DY: 0,
		},
	}
	ctrls := &components.ControlsComponent{
		Controls: map[ebiten.Key]func(*components.TransformComponent){
			ebiten.KeyW:     func(transformComponent *components.TransformComponent) { transformComponent.Velocity.DY = -1 },
			ebiten.KeyD:     func(transformComponent *components.TransformComponent) { transformComponent.Velocity.DX = 1 },
			ebiten.KeyS:     func(transformComponent *components.TransformComponent) { transformComponent.Velocity.DY = 1 },
			ebiten.KeyA:     func(transformComponent *components.TransformComponent) { transformComponent.Velocity.DX = -1 },
			ebiten.KeyShift: func(transformComponent *components.TransformComponent) { transformComponent.Speed = 2 },
		},
	}

	game.registry.AddSystem(renderer)
	game.registry.AddSystem(input)
	game.registry.AddSystem(movement)

	game.registry.AddComponent(0, transform1)
	game.registry.AddComponent(0, ctrls)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
