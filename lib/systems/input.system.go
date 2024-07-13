package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputSystem struct{}

func NewInputSystem() *InputSystem {
	return &InputSystem{}
}

func (iss *InputSystem) Update(registry *core.Registry) {
	controlsType := reflect.TypeOf(&components.ControlsComponent{})
	for _, component := range registry.GetAllComponentsOfType(controlsType) {
		controlsComponent := component.(*components.ControlsComponent)
		for key, _ := range controlsComponent.Controls {
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
