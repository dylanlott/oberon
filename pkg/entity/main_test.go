package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	world := NewWorld()
	tickDuration := time.Second // 1 second per tick

	entity := world.NewEntity()

	world.AddResource(entity, Resource{Name: "Gold", Count: 0})
	world.AddStarSystem(entity, StarSystem{Name: "Alpha Centauri"})

	for i := 0; i < 10; i++ {
		world.Update(1.0)
		fmt.Printf("Tick %d\n", i)
		fmt.Printf("Entity %d Resource: %+v\n", entity, *world.resources[world.entityIndex[entity]])
		fmt.Printf("Entity %d Star System: %+v\n", entity, *world.starSystems[world.entityIndex[entity]])
		time.Sleep(tickDuration)
	}

	world.AddStarSystem(entity, StarSystem{Name: "Proxima Centauri"})
	fmt.Printf("Entity %d Updated Star System: %+v\n", entity, *world.starSystems[world.entityIndex[entity]])
}
