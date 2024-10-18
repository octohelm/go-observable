package agent

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/go-courier/logr"
	"github.com/octohelm/go-observable/observable"
	observablejoin "github.com/octohelm/go-observable/observable/join"
	observableutil "github.com/octohelm/go-observable/observable/util"
)

type Agent[T any] struct {
	done    chan struct{}
	closed  atomic.Bool
	sources []observable.Observable[T]
	wg      sync.WaitGroup
}

func (x *Agent[T]) Init(ctx context.Context) error {
	x.done = make(chan struct{})
	return nil
}

func (x *Agent[T]) Subscribe(o observable.Observable[T]) {
	x.sources = append(x.sources, o)
}

func (x *Agent[T]) Go(ctx context.Context, action func(ctx context.Context) error) {
	x.wg.Add(1)

	go func() {
		defer x.wg.Done()

		if err := action(ctx); err != nil {
			logr.FromContext(ctx).Error(err)
		}
	}()
}

func (x *Agent[T]) Disabled(ctx context.Context) bool {
	return len(x.sources) == 0
}

func (x *Agent[T]) Serve(ctx context.Context) error {
	if x.Disabled(ctx) {
		return nil
	}

	return observableutil.WaitUnit(x.done, observablejoin.Merge(x.sources...))
}

func (x *Agent[T]) Shutdown(ctx context.Context) error {
	if x.closed.Swap(true) {
		return nil
	}

	// close to stop all observable
	close(x.done)

	done := make(chan struct{})

	go func() {
		// graceful
		x.wg.Wait()

		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return nil
}
