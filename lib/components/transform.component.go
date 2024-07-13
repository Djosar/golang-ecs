package components

import "github.com/Djosar/kro-ecs/lib/util"

type TransformComponent struct {
	Speed     float32
	Direction string
	Position  util.Coordinate[float32]
	Velocity  util.Velocity
}
