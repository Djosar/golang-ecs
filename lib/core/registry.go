package core

import (
	"fmt"
	"reflect"
	"slices"
)

// Registry manages entities, components, and systems within the ECS architecture.
// It provides methods to add entities, components, and systems, and to update systems.
type Registry struct {
	nextEntityId int
	systems      map[reflect.Type]System
	components   map[reflect.Type]map[Entity]Component
}

// NewRegistry creates and returns a new instance of Registry.
//
// Returns:
//
//	*Registry: A pointer to the newly created Registry instance.
func NewRegistry() *Registry {
	return &Registry{
		nextEntityId: 0,
		systems:      make(map[reflect.Type]System),
		components:   make(map[reflect.Type]map[int]interface{}),
	}
}

// NewEntity creates a new entity and returns its identifier.
//
// Returns:
//
//	Entity: The identifier of the newly created entity.
func (r *Registry) NewEntity() Entity {
	r.nextEntityId++
	return r.nextEntityId
}

// AddComponent adds a component to a specified entity.
//
// Parameters:
//
//	entity (Entity): The entity to which the component is added.
//	component (Component): The component to be added.
func (r *Registry) AddComponent(entity Entity, component Component) {
	identifier := reflect.TypeOf(component)
	if r.components[identifier] == nil {
		r.components[identifier] = make(map[int]interface{})
	}
	r.components[identifier][entity] = component
}

// GetAllComponentsOfType returns all components of a specified type.
//
// Parameters:
//
//	componentType (reflect.Type): The type of components to retrieve.
//
// Returns:
//
//	map[Entity]Component: A map of entities to their respective components of the specified type.
func (r *Registry) GetAllComponentsOfType(componentType reflect.Type) map[Entity]Component {
	return r.components[componentType]
}

// GetComponent returns the component of a specified type for a given entity.
//
// Parameters:
//
//	componentType (reflect.Type): The type of component to retrieve.
//	entity (Entity): The entity whose component is to be retrieved.
//
// Returns:
//
//	Component: The component of the specified type for the given entity.
func (r *Registry) GetComponent(componentType reflect.Type, entity Entity) Component {
	return r.components[componentType][entity]
}

// AddSystem adds a system to the registry.
//
// Parameters:
//
//	system (System): The system to be added.
func (r *Registry) AddSystem(system System) {
	if system != nil {
		identifier := reflect.TypeOf(system)
		r.systems[identifier] = system
	} else {
		fmt.Println("SYSTEM IS NIL")
	}
}

// UpdateSystems updates all systems in the registry, excluding the specified types.
//
// Parameters:
//
//	excludedSystems ([]reflect.Type): A slice of system types to be excluded from the update.
func (r *Registry) UpdateSystems(excludedSystems []reflect.Type) {
	for systemType, system := range r.systems {
		if !slices.Contains(excludedSystems, systemType) {
			system.Update(r)
		}
	}
}

// GetSystem returns the system of a specified type.
//
// Parameters:
//
//	systemType (reflect.Type): The type of system to retrieve.
//
// Returns:
//
//	System: The system of the specified type.
func (r *Registry) GetSystem(systemType reflect.Type) System {
	return r.systems[systemType]
}

// GetSystems returns all systems in the registry.
//
// Returns:
//
//	map[reflect.Type]System: A map of all systems in the registry.
func (r *Registry) GetSystems() map[reflect.Type]System {
	return r.systems
}
