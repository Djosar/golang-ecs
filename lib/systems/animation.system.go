package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

type AnimationSystem struct{}

func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{}
}

func (as *AnimationSystem) Update(registry *core.Registry) {
	transfType := reflect.TypeOf(&components.TransformComponent{})
	animType := reflect.TypeOf(&components.AnimationComponent{})

	for entity, component := range registry.GetAllComponentsOfType(transfType) {
		transform := component.(*components.TransformComponent)
		animationComp := registry.GetComponent(animType, entity).(*components.AnimationComponent)

		for identifier, handler := range animationComp.AnimationHandlers {
			if handler(transform) {
				animationComp.CurrentAnimation = identifier
			}
		}

		if currentAnimation := animationComp.Animations[animationComp.CurrentAnimation]; currentAnimation != nil {
			if currentAnimation.Counter < currentAnimation.Marker {
				if currentAnimation.Counter%currentAnimation.AnimationSpeed == 0 {
					if currentAnimation.FrameIndex < (len(currentAnimation.Frames) - 1) {
						currentAnimation.FrameIndex += 1
					} else {
						currentAnimation.FrameIndex = 0
					}
				}
				currentAnimation.Counter += 1
			} else {
				currentAnimation.Counter = 0
			}
		}
	}
}
