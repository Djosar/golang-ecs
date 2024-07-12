package main

import (
	"log"
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/systems"
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
	pos1 := &components.PositionComponent{
		X: 100,
		Y: 100,
	}
	velo1 := &components.VelocityComponent{
		DX: 0,
		DY: 0,
	}
	ctrls := &components.ControlsComponent{
		Controls: map[ebiten.Key]func(*components.VelocityComponent){
			ebiten.KeyW: func(vc *components.VelocityComponent) { vc.DY = -1; vc.DX = 0 },
			ebiten.KeyD: func(vc *components.VelocityComponent) { vc.DX = 1; vc.DY = 0 },
			ebiten.KeyS: func(vc *components.VelocityComponent) { vc.DY = 1; vc.DX = 0 },
			ebiten.KeyA: func(vc *components.VelocityComponent) { vc.DX = -1; vc.DY = 0 },
		},
	}

	game.registry.AddSystem(renderer)
	game.registry.AddSystem(input)
	game.registry.AddSystem(movement)

	game.registry.AddComponent(0, pos1)
	game.registry.AddComponent(0, velo1)
	game.registry.AddComponent(0, ctrls)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
