package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

// MovementSystem is responsible for updating the position and velocity
// of entities within the entity-component-system (ECS) architecture.
// It processes movement controls and applies the resulting transformations.
type MovementSystem struct{}

// NewMovementSystem creates and returns a new instance of MovementSystem.
//
// Returns:
//
//	*MovementSystem: A pointer to the newly created MovementSystem instance.
func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

// Update iterates through all entities that have both a TransformComponent and
// a ControlsComponent. It updates the velocity and position of each entity based
// on the current controls and applies the transformations.
//
// Parameters:
//
//	registry (*core.Registry): The registry containing all entities and components in the ECS.
func (ms *MovementSystem) Update(registry *core.Registry) {
	ctrlType := reflect.TypeOf(&components.ControlsComponent{})

	transfType := reflect.TypeOf(&components.TransformComponent{})
	for entity, transf := range registry.GetAllComponentsOfType(transfType) {
		transformComponent := transf.(*components.TransformComponent)
		controls := registry.GetComponent(ctrlType, entity).(*components.ControlsComponent)
		transformComponent.Velocity.DX = 0
		transformComponent.Velocity.DY = 0
		transformComponent.Speed = 1

		for _, key := range controls.ControlsBuffer {
			if ctrl := controls.Controls[key]; ctrl != nil {
				ctrl(transformComponent)
			}
		}

		transformComponent.Position.X += transformComponent.Speed * transformComponent.Velocity.DX
		transformComponent.Position.Y += transformComponent.Speed * transformComponent.Velocity.DY
	}
}
