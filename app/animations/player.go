package animations

import (
	"bytes"
	"image"

	"github.com/Djosar/kro-ecs/assets"
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func CreatePlayerAnimations() (map[components.AnimationIdentifier]*util.Animation, error) {
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

	animations := map[components.AnimationIdentifier]*util.Animation{
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
	}

	return animations, nil
}
