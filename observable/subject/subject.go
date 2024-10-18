package subject

import (
	"iter"
	"sync"
	"sync/atomic"

	"github.com/octohelm/go-observable/observable"
)

type Subject[T any] interface {
	observable.ValueNotifier[T]
	observable.Observable[T]
	observable.Cancelable
}

func NewSubject[T any]() Subject[T] {
	return &subject[T]{}
}

type subject[T any] struct {
	observers sync.Map
	closed    atomic.Bool
	err       atomic.Value
}

func (s *subject[T]) Send(value T) {
	if s.closed.Load() {
		return
	}

	for o := range s.observer() {
		if x, ok := o.(observable.ValueNotifier[T]); ok {
			x.Send(value)
		}
	}
}

func (s *subject[T]) Observe() observable.Observer[T] {
	ch := make(chan T)

	ob := observable.From(func(yield func(T, error) bool) {
		for v := range ch {
			if !yield(v, nil) {
				return
			}
		}
	})

	o := &notifier[T]{
		ch:       ch,
		Observer: ob.Observe(),
	}

	s.observers.Store(o, true)

	go func() {
		<-o.Done()

		s.observers.Delete(o)
	}()

	return o
}

type notifier[T any] struct {
	observable.Observer[T]
	ch chan T
}

func (n *notifier[T]) Send(v T) {
	select {
	case <-n.Done():
	case n.ch <- v:
	}
}

func (s *subject[T]) observer() iter.Seq[observable.Observer[T]] {
	obs := make([]observable.Observer[T], 0)

	for k, _ := range s.observers.Range {
		obs = append(obs, k.(observable.Observer[T]))
	}

	return func(yield func(observable.Observer[T]) bool) {
		for _, o := range obs {
			if !yield(o) {
				return
			}
		}
	}
}

func (s *subject[T]) CancelCause(err error) {
	if s.closed.Swap(true) {
		return
	}

	for o := range s.observer() {
		o.CancelCause(err)
	}

	return
}
