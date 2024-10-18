package operators

import (
	"sync"
	"time"

	"github.com/octohelm/go-observable/observable"
)

func BufferTime[T any](duration time.Duration) observable.Operator[T, []T] {
	return func(src observable.Observable[T]) observable.Observable[[]T] {
		return observable.From(func(yield func([]T, error) bool) {
			upstream := src.Observe()

			m := &sync.Mutex{}
			values := make([]T, 0)

			appendValue := func(v T) {
				m.Lock()
				defer m.Unlock()

				values = append(values, v)
			}

			emit := func() []T {
				m.Lock()
				defer m.Unlock()

				x := values
				values = make([]T, 0)
				return x
			}

			timer := time.NewTicker(duration)
			defer timer.Stop()

			for {
				select {
				case <-timer.C:
					if !yield(emit(), nil) {
						return
					}
				default:
					v, err, ok := upstream.Next()
					if !ok {
						return
					}
					if err != nil {
						if !yield(nil, err) {
							return
						}
					}
					appendValue(v)
				}
			}
		})
	}
}
