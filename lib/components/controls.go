package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ControlsComponent struct {
	Controls       map[ebiten.Key]func(*TransformComponent)
	ControlsBuffer []ebiten.Key
}
