package systems

import (
	"fmt"
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

		fmt.Println(animationComp.CurrentAnimation)
	}
}
