package factories

import (
	"github.com/Djosar/kro-ecs/lib/components"
	"github.com/Djosar/kro-ecs/lib/core"
)

func CharacterFactory(registry *core.Registry, transform *components.TransformComponent, animation *components.AnimationComponent) (entity core.Entity) {
	entity = registry.NewEntity()
	registry.AddComponent(entity, transform)
	registry.AddComponent(entity, animation)

	return entity
}
