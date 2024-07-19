package assets

import (
	_ "embed"
	_ "image/png"
)

var (
	//go:embed sprites/idle.png
	IdleSpriteSheet []byte

	//go:embed sprites/run.png
	RunSpriteSheet []byte

	//go:embed sprites/walk.png
	WalkSpriteSheet []byte
)
