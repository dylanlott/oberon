package main

import (
	"context"
	"time"

	"oberon/pkg/game"
)

func main() {
	config := game.Config{}
	looper, _ := game.New(config)

	initialState := []game.Observer{
		&game.Player{
			Username: "shakezula",
			State:    map[string]interface{}{},
		},
		&game.Player{
			Username: "meatwad",
			State:    map[string]interface{}{},
		},
		&game.Player{
			Username: "frylock",
			State:    map[string]interface{}{},
		},
		&game.Player{
			Username: "carl",
			State:    map[string]interface{}{},
		},
	}

	// register everything
	for _, item := range initialState {
		looper.Register(item)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	looper.Run(ctx, time.Second*1)
}
