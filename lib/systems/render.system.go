package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/hajimehoshi/ebiten/v2"
)

// RenderSystem is responsible for rendering entities within the entity-component-system (ECS) architecture.
// It draws the current animation frame of each entity onto the screen.
type RenderSystem struct {
	Screen *ebiten.Image
}

// NewRenderSystem creates and returns a new instance of RenderSystem.
//
// Returns:
//
//	*RenderSystem: A pointer to the newly created RenderSystem instance.
func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

// Update iterates through all entities that have both a TransformComponent and
// an AnimationComponent. It renders the current frame of the entity's animation
// to the screen based on the entity's position.
//
// Parameters:
//
//	registry (*core.Registry): The registry containing all entities and components in the ECS.
func (rs *RenderSystem) Update(registry *core.Registry) {
	transfType := reflect.TypeOf(&components.TransformComponent{})
	animType := reflect.TypeOf(&components.AnimationComponent{})

	for entity, p := range registry.GetAllComponentsOfType(transfType) {
		transf := p.(*components.TransformComponent)
		animationComp := registry.GetComponent(animType, entity).(*components.AnimationComponent)
		currentAnimation := animationComp.GetCurrentAnimation()
		if currentAnimation != nil {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(transf.Position.X), float64(transf.Position.Y))
			currentFrame := currentAnimation.GetCurrentFrame()
			rs.Screen.DrawImage(currentFrame, opts)
		}
	}

}
