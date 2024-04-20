// Package game is responsible for setting up the framework
// for a simple tick-looped terminal game.
package game

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Observer defines an interface for observers to the game loop.
type Observer interface {
	// Tick advances the game loop of the observer
	Tick(state map[string]interface{}, input chan WriteOp)
}

// WriteOp is used to update the game state with a key value operation
type WriteOp struct {
	Key   string
	Value interface{}
}

// Looper starts a game loop that can exposes a plugin interface
type Looper struct {
	obs []Observer
	in  chan WriteOp
}

type Config struct {
	bufSize  int64
	interval int64
}

type State map[string]interface{}

type Player struct {
	Observer

	Username string
	State    State
}

// New returns a Looper and the handle to its input channel.
// * All write ops must be passed in to this channel.
func New(cfg Config) (*Looper, chan WriteOp) {
	in := make(chan WriteOp, cfg.bufSize)
	looper := &Looper{
		obs: make([]Observer, cfg.bufSize),
		in:  in,
	}
	return looper, in
}

// Run starts a game loop that ticks at every interval and accepts
// transactions to update the game state, which can only ever be
// updated by passing in a transaction to update a key & value.
func (l *Looper) Run(ctx context.Context, interval time.Duration) {
	var gameState = map[string]interface{}{}
	ticker := time.NewTicker(interval)

	go func() {
		for inputTx := range l.in {
			gameState[inputTx.Key] = inputTx.Value
		}
	}()

	for {
		<-ticker.C
		for _, o := range l.obs {
			o.Tick(gameState, l.in)
		}
	}
}

// Register adds a plugin to the set of plugins that are hooked into
// the looper.
func (l *Looper) Register(obs Observer) error {
	if obs == nil {
		return fmt.Errorf("must supply Observer")
	}
	l.obs = append(l.obs, obs)
	log.Printf("registered observer: %+v", obs)
	return nil
}
