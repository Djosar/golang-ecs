package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputSystem is responsible for handling input within the entity-component-system (ECS) architecture.
// It updates control components by detecting key presses and releases, maintaining a buffer of active controls.
type InputSystem struct{}

// NewInputSystem creates and returns a new instance of InputSystem.
//
// Returns:
//
//	*InputSystem: A pointer to the newly created InputSystem instance.
func NewInputSystem() *InputSystem {
	return &InputSystem{}
}

// Update iterates through all entities that have a ControlsComponent. It updates the input state
// by detecting key presses and releases and maintaining a buffer of active controls.
//
// Parameters:
//
//	registry (*core.Registry): The registry containing all entities and components in the ECS.
func (iss *InputSystem) Update(registry *core.Registry) {
	controlsType := reflect.TypeOf(&components.ControlsComponent{})
	for _, component := range registry.GetAllComponentsOfType(controlsType) {
		controlsComponent := component.(*components.ControlsComponent)
		for key := range controlsComponent.Controls {
			if inpututil.IsKeyJustPressed(key) {
				controlsComponent.ControlsBuffer = append(controlsComponent.ControlsBuffer, key)
			}
			if inpututil.IsKeyJustReleased(key) {
				releasedKeyIdx := -1
				for idx, k := range controlsComponent.ControlsBuffer {
					if k == key {
						releasedKeyIdx = idx
					}
				}
				if releasedKeyIdx >= 0 {
					controlsComponent.ControlsBuffer = append(controlsComponent.ControlsBuffer[:releasedKeyIdx], controlsComponent.ControlsBuffer[releasedKeyIdx+1:]...)
				}
			}
		}
	}
}
