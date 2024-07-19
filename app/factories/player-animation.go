package factories

import (
	"bytes"
	"image"

	"github.com/Djosar/kro-ecs/app/assets"
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func PlayerAnimationComponentFactory() (*components.AnimationComponent, error) {
	idleSpriteFile, _, err := image.Decode(bytes.NewReader(assets.IdleSpriteSheet))
	if err != nil {
		return nil, err
	}
	walkSpriteFile, _, err := image.Decode(bytes.NewReader(assets.WalkSpriteSheet))
	if err != nil {
		return nil, err
	}
	sprintSpriteFile, _, err := image.Decode(bytes.NewReader(assets.RunSpriteSheet))
	if err != nil {
		return nil, err
	}
	return &components.AnimationComponent{
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
	}, nil
}
