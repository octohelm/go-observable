package internal

import (
	"errors"
	"sync/atomic"
)

var Completed = errors.New("completed")

type Completer struct {
	OnComplete func()

	created atomic.Bool
	closed  atomic.Bool
	done    chan struct{}
	err     atomic.Value
}

func (g *Group) Err() error {
	err := g.err.Load()
	if err == nil {
		return nil
	}
	return err.(error)
}

func (o *Completer) Done() <-chan struct{} {
	return o.done
}

func (o *Completer) CancelCause(err error) {
	if o.closed.Swap(true) {
		return
	}

	if err == nil {
		err = Completed
	}
	if o.OnComplete != nil {
		o.OnComplete()
	}
	o.err.Store(err)
	close(o.done)
	return
}

func (o *Completer) Init() {
	if o.created.Swap(true) {
		return
	}
	o.done = make(chan struct{})
}
