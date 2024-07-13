package util

import "github.com/hajimehoshi/ebiten/v2"

type Animation struct {
	Frames  []*ebiten.Image
	Marker  int
	Counter int
}
