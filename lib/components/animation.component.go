package components

import "github.com/Djosar/kro-ecs/lib/util"

type AnimationIdentifier string

type AnimationComponent struct {
	CurrentAnimation  AnimationIdentifier
	AnimationHandlers map[AnimationIdentifier]func(*TransformComponent) bool
	Animations        map[AnimationIdentifier]*util.Animation
}

func (ac *AnimationComponent) GetCurrentAnimation() *util.Animation {
	return ac.Animations[ac.CurrentAnimation]
}
