package main

import (
	"fmt"
	"time"
)

// Resource component
type Resource struct {
	Name  string
	Count int
}

// StarSystem component
type StarSystem struct {
	Name string
}

type Entity int

type World struct {
	entities     []Entity
	resources    []*Resource    // Use slice for contiguous memory
	starSystems  []*StarSystem  // Use slice for contiguous memory
	entityIndex  map[Entity]int // Map entity to index in slices
	nextEntityID Entity
}

func NewWorld() *World {
	return &World{
		resources:   make([]*Resource, 0),
		starSystems: make([]*StarSystem, 0),
		entityIndex: make(map[Entity]int),
	}
}

func (w *World) NewEntity() Entity {
	entity := w.nextEntityID
	w.nextEntityID++
	w.entities = append(w.entities, entity)
	index := len(w.entities) - 1
	w.entityIndex[entity] = index
	w.resources = append(w.resources, nil)
	w.starSystems = append(w.starSystems, nil)
	return entity
}

func (w *World) AddResource(entity Entity, resource Resource) {
	index := w.entityIndex[entity]
	w.resources[index] = &resource
}

func (w *World) AddStarSystem(entity Entity, starSystem StarSystem) {
	index := w.entityIndex[entity]
	w.starSystems[index] = &starSystem
}

func (w *World) Update(delta float64) {
	for _, entity := range w.entities {
		index := w.entityIndex[entity]
		res := w.resources[index]
		if res != nil {
			// Simulate resource change (e.g., increment count)
			res.Count += int(delta)
		}
	}
}

func main() {
	world := NewWorld()

	// Create a new entity
	entity := world.NewEntity()

	// Add components to the entity
	world.AddResource(entity, Resource{Name: "Gold", Count: 0})
	world.AddStarSystem(entity, StarSystem{Name: "Alpha Centauri"})

	// Define the tick duration
	tickDuration := time.Second // 1 second per tick

	// Run the update loop
	for i := 0; i < 10; i++ { // Run the loop 10 times for demonstration
		world.Update(1.0) // Update the world with a delta time of 1.0
		fmt.Printf("Tick %d\n", i)
		fmt.Printf("Entity %d Resource: %+v\n", entity, *world.resources[world.entityIndex[entity]])
		fmt.Printf("Entity %d Star System: %+v\n", entity, *world.starSystems[world.entityIndex[entity]])
		time.Sleep(tickDuration)
	}

	// Update the StarSystem component
	world.AddStarSystem(entity, StarSystem{Name: "Proxima Centauri"})
	fmt.Printf("Entity %d Updated Star System: %+v\n", entity, *world.starSystems[world.entityIndex[entity]])
}
