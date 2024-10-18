package internal

import (
	"sync"
	"sync/atomic"
)

type Group struct {
	Completer

	wg      sync.WaitGroup
	created atomic.Bool
}

func (g *Group) Wait() {
	g.wg.Wait()
	g.CancelCause(nil)
}

func (g *Group) initOnce() {
	if g.created.Swap(true) {
		return
	}

	g.Completer.Init()
}

func (g *Group) Go(fn func() error) {
	g.initOnce()

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		if err := fn(); err != nil {
			g.CancelCause(err)
		}
	}()
}
