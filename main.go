package main

import (
	"bytes"
	"image"
	"log"
	"reflect"

	"github.com/Djosar/kro-ecs/assets"
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

	idleSpriteFile, _, err := image.Decode(bytes.NewReader(assets.IdleSpriteSheet))
	if err != nil {
		log.Fatal(err)
	}
	walkSpriteFile, _, err := image.Decode(bytes.NewReader(assets.WalkSpriteSheet))
	if err != nil {
		log.Fatal(err)
	}
	sprintSpriteFile, _, err := image.Decode(bytes.NewReader(assets.RunSpriteSheet))
	if err != nil {
		log.Fatal(err)
	}

	animations := &components.AnimationComponent{
		CurrentAnimation: "idle_down",
		Animations: map[components.AnimationIdentifier]*util.Animation{
			"idle_up": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(idleSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 3),
						util.NewCoordinate(1, 3),
						util.NewCoordinate(2, 3),
						util.NewCoordinate(3, 3),
					},
				),
				60,
				10,
			),
			"idle_down": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(idleSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 2),
						util.NewCoordinate(1, 2),
						util.NewCoordinate(2, 2),
						util.NewCoordinate(3, 2),
					},
				),
				60,
				10,
			),
			"idle_left": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(idleSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 1),
						util.NewCoordinate(1, 1),
						util.NewCoordinate(2, 1),
						util.NewCoordinate(3, 1),
					},
				),
				60,
				10,
			),
			"idle_right": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(idleSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 0),
						util.NewCoordinate(1, 0),
						util.NewCoordinate(2, 0),
						util.NewCoordinate(3, 0),
					},
				),
				60,
				10,
			),
			"walk_up": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(walkSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 3),
						util.NewCoordinate(1, 3),
						util.NewCoordinate(2, 3),
						util.NewCoordinate(3, 3),
						util.NewCoordinate(4, 3),
						util.NewCoordinate(5, 3),
						util.NewCoordinate(6, 3),
						util.NewCoordinate(7, 3),
					},
				),
				60,
				10,
			),
			"walk_down": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(walkSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 2),
						util.NewCoordinate(1, 2),
						util.NewCoordinate(2, 2),
						util.NewCoordinate(3, 2),
						util.NewCoordinate(4, 2),
						util.NewCoordinate(5, 2),
						util.NewCoordinate(6, 2),
						util.NewCoordinate(7, 2),
					},
				),
				60,
				10,
			),
			"walk_left": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(walkSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 1),
						util.NewCoordinate(1, 1),
						util.NewCoordinate(2, 1),
						util.NewCoordinate(3, 1),
						util.NewCoordinate(4, 1),
						util.NewCoordinate(5, 1),
						util.NewCoordinate(6, 1),
						util.NewCoordinate(7, 1),
					},
				),
				60,
				10,
			),
			"walk_right": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(walkSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 0),
						util.NewCoordinate(1, 0),
						util.NewCoordinate(2, 0),
						util.NewCoordinate(3, 0),
						util.NewCoordinate(4, 0),
						util.NewCoordinate(5, 0),
						util.NewCoordinate(6, 0),
						util.NewCoordinate(7, 0),
					},
				),
				60,
				10,
			),
			"sprint_up": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(sprintSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 3),
						util.NewCoordinate(1, 3),
						util.NewCoordinate(2, 3),
						util.NewCoordinate(3, 3),
						util.NewCoordinate(4, 3),
						util.NewCoordinate(5, 3),
						util.NewCoordinate(6, 3),
						util.NewCoordinate(7, 3),
					},
				),
				60,
				10,
			),
			"sprint_down": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(sprintSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 2),
						util.NewCoordinate(1, 2),
						util.NewCoordinate(2, 2),
						util.NewCoordinate(3, 2),
						util.NewCoordinate(4, 2),
						util.NewCoordinate(5, 2),
						util.NewCoordinate(6, 2),
						util.NewCoordinate(7, 2),
					},
				),
				60,
				10,
			),
			"sprint_left": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(sprintSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 1),
						util.NewCoordinate(1, 1),
						util.NewCoordinate(2, 1),
						util.NewCoordinate(3, 1),
						util.NewCoordinate(4, 1),
						util.NewCoordinate(5, 1),
						util.NewCoordinate(6, 1),
						util.NewCoordinate(7, 1),
					},
				),
				60,
				10,
			),
			"sprint_right": util.NewAnimation(
				util.GenerateFrames(
					ebiten.NewImageFromImage(sprintSpriteFile),
					80, 80,
					[]*util.Coordinate[int]{
						util.NewCoordinate(0, 0),
						util.NewCoordinate(1, 0),
						util.NewCoordinate(2, 0),
						util.NewCoordinate(3, 0),
						util.NewCoordinate(4, 0),
						util.NewCoordinate(5, 0),
						util.NewCoordinate(6, 0),
						util.NewCoordinate(7, 0),
					},
				),
				60,
				10,
			),
		},
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
