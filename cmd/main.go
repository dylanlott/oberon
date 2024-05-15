package main

import (
	"fmt"
	"time"

	"oberon/pkg/entity"
)

func main() {
	world := entity.NewWorld()
	tickDuration := time.Second // 1 second per tick

	entityID := world.NewEntity()

	world.AddResource(entityID, entity.Resource{Name: "Gold", Count: 0})
	world.AddShip(entityID, entity.Ship{
		Name:    "Karakoum",
		OwnerID: "shakezula",
		Cargo: []entity.Resource{
			{Name: "Gold", Count: 1},
		},
	})
	world.AddStarSystem(entityID, entity.StarSystem{Name: "Alpha Centauri"})

	ticker := time.NewTicker(tickDuration)
	defer ticker.Stop()
	done := make(chan bool)

	go func(world *entity.World, ticker *time.Ticker, done chan bool) {
		for {
			select {
			case <-done:
				fmt.Printf("initiating graceful shutdown...")
				return
			case t := <-ticker.C:
				world.Update(1.0) // Update the world with a delta time of 1.0
				fmt.Printf("Tick at %v\n", t)
			}
		}
	}(world, ticker, done)
}
