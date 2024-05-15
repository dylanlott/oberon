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

	ticker := time.NewTicker(tickDuration)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("initiating graceful shutdown...")
				return
			case t := <-ticker.C:
				world.Update(1.0) // Update the world with a delta time of 1.0
				fmt.Printf("Tick at %v\n", t)
				fmt.Printf("Entity %d Resource: %+v\n", entity, *world.resources[world.entityIndex[entity]])
				fmt.Printf("Entity %d Star System: %+v\n", entity, *world.starSystems[world.entityIndex[entity]])
			}
		}
	}()

	// Run the ticker for 10 ticks for example purposes
	time.Sleep(10 * tickDuration)
	done <- true
}
