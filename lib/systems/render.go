package systems

import (
	"image/color"
	"reflect"

	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RenderSystem struct {
	Screen *ebiten.Image
}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

func (rs *RenderSystem) Update(registry *core.Registry) {
	transfType := reflect.TypeOf(&components.TransformComponent{})

	for _, p := range registry.GetAllComponentsOfType(transfType) {
		transf := p.(*components.TransformComponent)
		vector.DrawFilledRect(rs.Screen, transf.Position.X, transf.Position.Y, 10, 10, color.White, false)
	}

}
