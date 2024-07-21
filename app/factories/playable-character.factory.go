package factories

import (
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/Djosar/kro-ecs/lib/util"
	"github.com/hajimehoshi/ebiten/v2"
)

// PlayableCharacterFactory creates a new playable character entity with the necessary components
// such as animation, transform, and controls. It initializes the entity and registers it with the
// provided registry.
//
// Parameters:
//
//	registry (*core.Registry): The registry where the new entity and its components will be registered.
//
// Returns:
//
//	core.Entity: The identifier of the newly created entity.
//	error: An error if any component cannot be created or registered.
func PlayableCharacterFactory(registry *core.Registry) (core.Entity, error) {
	entity := registry.NewEntity()

	// Create the animation component using the PlayerAnimationComponentFactory function
	animation, err := PlayerAnimationComponentFactory()
	if err != nil {
		return -1, err
	}

	// Initialize the transform component with default position, speed, direction, and velocity
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

	// Initialize the controls component with key bindings for movement and speed
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

	// Register the created components with the entity in the registry
	registry.AddComponent(entity, animation)
	registry.AddComponent(entity, transform)
	registry.AddComponent(entity, controls)
	return entity, nil
}
