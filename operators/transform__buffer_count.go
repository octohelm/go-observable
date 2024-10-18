package operators

import (
	"github.com/octohelm/go-observable/observable"
)

func BufferCount[T any](n int) observable.Operator[T, []T] {
	return func(src observable.Observable[T]) observable.Observable[[]T] {
		return observable.From(func(yield func([]T, error) bool) {
			upstream := src.Observe()

			values := make([]T, n)
			i := 0

			emit := func() []T {
				x := values
				values = make([]T, n)
				i = 0
				return x
			}

			for {
				v, err, ok := upstream.Next()
				if !ok {
					return
				}

				values[i] = v
				i++

				if i == n {
					if !yield(emit(), err) {
						return
					}
				}
			}
		})
	}
}
