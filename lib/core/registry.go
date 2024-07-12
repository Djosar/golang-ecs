package core

import (
	"fmt"
	"reflect"
	"slices"
)

type Registry struct {
	nextEntityId int
	systems      map[reflect.Type]System
	components   map[reflect.Type]map[Entity]Component
}

func NewRegistry() *Registry {
	return &Registry{
		nextEntityId: 0,
		systems:      make(map[reflect.Type]System),
		components:   make(map[reflect.Type]map[int]interface{}),
	}
}

func (r *Registry) NewEntity() Entity {
	r.nextEntityId++
	return r.nextEntityId
}

func (r *Registry) AddComponent(entity Entity, component Component) {
	identifier := reflect.TypeOf(component)
	if r.components[identifier] == nil {
		r.components[identifier] = make(map[int]interface{})
	}
	r.components[identifier][entity] = component
}

func (r *Registry) GetAllComponentsOfType(componentType reflect.Type) map[Entity]Component {
	return r.components[componentType]
}

func (r *Registry) GetComponent(componentType reflect.Type, entity Entity) Component {
	return r.components[componentType][entity]
}

func (r *Registry) AddSystem(system System) {
	if system != nil {
		identifier := reflect.TypeOf(system)
		r.systems[identifier] = system
	} else {
		fmt.Println("SYSTEM IS NIL")
	}
}

func (r *Registry) UpdateSystems(excludedSystems []reflect.Type) {
	for systemType, system := range r.systems {
		if !slices.Contains(excludedSystems, systemType) {
			system.Update(r)
		}
	}
}

func (r *Registry) GetSystem(systemType reflect.Type) System {
	return r.systems[systemType]
}

func (r *Registry) GetSystems() map[reflect.Type]System {
	return r.systems
}
