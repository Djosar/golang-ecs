package systems

import (
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderSystem struct {
	Screen *ebiten.Image
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

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
