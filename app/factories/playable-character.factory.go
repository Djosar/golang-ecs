package factories

import (
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func PlayableCharacterFactory(registry *core.Registry) (core.Entity, error) {
	entity := registry.NewEntity()
	animation, err := PlayerAnimationComponentFactory()
	if err != nil {
		return -1, err
	}
	transform := &components.TransformComponent{
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

	controls := &components.ControlsComponent{
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
	registry.AddComponent(entity, animation)
	registry.AddComponent(entity, transform)
	registry.AddComponent(entity, controls)
	return entity, nil
}
