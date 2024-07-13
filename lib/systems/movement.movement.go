package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

type MovementSystem struct{}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

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
