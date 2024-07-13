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
	animation := systems.NewAnimationSystem()
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
			ebiten.KeyW: func(transformComponent *components.TransformComponent) {
				transformComponent.Velocity.DY = -1
				transformComponent.Direction = "up"
			},
			ebiten.KeyD: func(transformComponent *components.TransformComponent) {
				transformComponent.Velocity.DX = 1
				transformComponent.Direction = "right"
			},
			ebiten.KeyS: func(transformComponent *components.TransformComponent) {
				transformComponent.Velocity.DY = 1
				transformComponent.Direction = "down"
			},
			ebiten.KeyA: func(transformComponent *components.TransformComponent) {
				transformComponent.Velocity.DX = -1
				transformComponent.Direction = "left"
			},
			ebiten.KeyShift: func(transformComponent *components.TransformComponent) { transformComponent.Speed = 2 },
		},
	}
	animations := &components.AnimationComponent{
		CurrentAnimation: "idle_down",
		Animations:       map[components.AnimationIdentifier]*util.Animation{},
		AnimationHandlers: map[components.AnimationIdentifier]func(*components.TransformComponent) bool{
			"idle_up": func(tc *components.TransformComponent) bool {
				return tc.Direction == "up" && tc.Velocity.DX == 0 && tc.Velocity.DY == 0
			},
			"idle_down": func(tc *components.TransformComponent) bool {
				return tc.Direction == "down" && tc.Velocity.DX == 0 && tc.Velocity.DY == 0
			},
			"idle_left": func(tc *components.TransformComponent) bool {
				return tc.Direction == "left" && tc.Velocity.DX == 0 && tc.Velocity.DY == 0
			},
			"idle_right": func(tc *components.TransformComponent) bool {
				return tc.Direction == "right" && tc.Velocity.DX == 0 && tc.Velocity.DY == 0
			},
			"walk_up": func(tc *components.TransformComponent) bool {
				return tc.Direction == "up" && tc.Velocity.DX == 0 && tc.Velocity.DY == -1 && tc.Speed == 1
			},
			"walk_down": func(tc *components.TransformComponent) bool {
				return tc.Direction == "down" && tc.Velocity.DX == 0 && tc.Velocity.DY == 1 && tc.Speed == 1
			},
			"walk_left": func(tc *components.TransformComponent) bool {
				return tc.Direction == "left" && tc.Velocity.DX == -1 && tc.Velocity.DY == 0 && tc.Speed == 1
			},
			"walk_right": func(tc *components.TransformComponent) bool {
				return tc.Direction == "right" && tc.Velocity.DX == 1 && tc.Velocity.DY == 0 && tc.Speed == 1
			},
			"sprint_up": func(tc *components.TransformComponent) bool {
				return tc.Direction == "up" && tc.Velocity.DX == 0 && tc.Velocity.DY == -1 && tc.Speed == 2
			},
			"sprint_down": func(tc *components.TransformComponent) bool {
				return tc.Direction == "down" && tc.Velocity.DX == 0 && tc.Velocity.DY == 1 && tc.Speed == 2
			},
			"sprint_left": func(tc *components.TransformComponent) bool {
				return tc.Direction == "left" && tc.Velocity.DX == -1 && tc.Velocity.DY == 0 && tc.Speed == 2
			},
			"sprint_right": func(tc *components.TransformComponent) bool {
				return tc.Direction == "right" && tc.Velocity.DX == 1 && tc.Velocity.DY == 0 && tc.Speed == 2
			},
		},
	}

	game.registry.AddSystem(renderer)
	game.registry.AddSystem(input)
	game.registry.AddSystem(movement)
	game.registry.AddSystem(animation)

	game.registry.AddComponent(0, transform1)
	game.registry.AddComponent(0, ctrls)
	game.registry.AddComponent(0, animations)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
