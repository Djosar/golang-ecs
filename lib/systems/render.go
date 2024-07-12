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
	posType := reflect.TypeOf(&components.PositionComponent{})

	for _, p := range registry.GetAllComponentsOfType(posType) {
		pos := p.(*components.PositionComponent)
		vector.DrawFilledRect(rs.Screen, pos.X, pos.Y, 10, 10, color.White, false)
	}

}
