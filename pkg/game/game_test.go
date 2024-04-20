package game

import (
	"context"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestRun(t *testing.T) {
	l := &Looper{}
	ctx, cancel := context.WithCancel(context.Background())
	go l.Run(ctx, 1*time.Second)
	cancel()
	<-ctx.Done()
}

func TestPlugin(t *testing.T) {
	is := is.New(t)
	l := &Looper{
		obs: make([]Observer, 0),
		in:  make(chan WriteOp),
	}

	m := &mock{}
	l.Register(m)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go l.Run(ctx, time.Second*1)

	time.Sleep(time.Second * 2)
	is.Equal(m.called, 1)
}

func TestReadWrite(t *testing.T) {
	is := is.New(t)
	l := &Looper{
		obs: make([]Observer, 0),
		in:  make(chan WriteOp),
	}

	m := &mock{}
	l.Register(m)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go l.Run(ctx, time.Second*1)

	time.Sleep(time.Second * 2)
	is.Equal(m.called, 1)
}

type mock struct{ called int }

func (m *mock) Tick(state map[string]interface{}, in chan WriteOp) {
	m.called++
}
