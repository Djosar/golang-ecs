package main

import (
	"log"

	"github.com/Djosar/kro-ecs/app/animations"
	"github.com/Djosar/kro-ecs/app/factories"
	"github.com/Djosar/kro-ecs/app/game"
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/systems"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame(core.NewRegistry())
	renderer := systems.NewRenderSystem()
	movement := systems.NewMovementSystem()
	input := systems.NewInputSystem()
	animation := systems.NewAnimationSystem()

	transformComponent := &components.TransformComponent{
		Position: util.Coordinate[float32]{
			X: 0,
			Y: 0,
		},
		Speed:     1,
		Direction: "down",
		Velocity: util.Velocity{
			DX: 0,
			DY: 0,
		},
	}
	controlsComponent := &components.ControlsComponent{
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

	animations, err := animations.CreatePlayerAnimations()
	if err != nil {
		log.Fatal(err)
		return
	}
	animationComponent := &components.AnimationComponent{
		CurrentAnimation: "idle_down",
		Animations:       animations,
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

	game.Registry.AddSystem(renderer)
	game.Registry.AddSystem(input)
	game.Registry.AddSystem(movement)
	game.Registry.AddSystem(animation)

	factories.PlayableCharacterFactory(game.Registry, transformComponent, animationComponent, controlsComponent)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
