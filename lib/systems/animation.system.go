package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

// AnimationSystem is responsible for updating animation components
// within the entity-component-system (ECS) architecture. It handles the
// transition of animation states based on the provided handlers and updates
// the current frame of each animation.
type AnimationSystem struct{}

// NewAnimationSystem creates and returns a new instance of AnimationSystem.
//
// Returns:
//
//	*AnimationSystem: A pointer to the newly created AnimationSystem instance.
func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{}
}

// Update iterates through all entities that have both a TransformComponent and
// an AnimationComponent. It updates the animation state of each entity based on
// the provided handlers and advances the animation frames.
//
// Parameters:
//
//	registry (*core.Registry): The registry containing all entities and components in the ECS.
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
