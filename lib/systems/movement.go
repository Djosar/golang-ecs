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
	posType := reflect.TypeOf(&components.PositionComponent{})
	velType := reflect.TypeOf(&components.VelocityComponent{})
	ctrlType := reflect.TypeOf(&components.ControlsComponent{})

	for entity, velo := range registry.GetAllComponentsOfType(velType) {
		velocity := velo.(*components.VelocityComponent)
		position := registry.GetComponent(posType, entity).(*components.PositionComponent)
		controls := registry.GetComponent(ctrlType, entity).(*components.ControlsComponent)
		velocity.DX = 0
		velocity.DY = 0

		if len(controls.ControlsBuffer) > 0 {
			currentKey := controls.ControlsBuffer[len(controls.ControlsBuffer)-1]
			if currentControl := controls.Controls[currentKey]; currentControl != nil {
				currentControl(velocity)
			}
		}

		position.X += velocity.DX
		position.Y += velocity.DY
	}

}
