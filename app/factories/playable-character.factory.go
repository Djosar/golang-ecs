package factories

import (
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

func PlayableCharacterFactory(
	registry *core.Registry,
	transform *components.TransformComponent,
	animation *components.AnimationComponent,
	controls *components.ControlsComponent,
) (entity core.Entity) {
	entity = CharacterFactory(registry, transform, animation)
	registry.AddComponent(entity, controls)
	return entity
}
