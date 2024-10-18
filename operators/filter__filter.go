package operators

import (
	"github.com/octohelm/go-observable/observable"
)

func Filter[T any](filter func(x T) bool) observable.Operator[T, T] {
	return func(src observable.Observable[T]) observable.Observable[T] {
		return observable.From(func(yield func(T, error) bool) {
			upstream := src.Observe()

			for {
				v, err, ok := upstream.Next()
				if !ok || err != nil {
					return
				}

				if filter(v) {
					if !yield(v, nil) {
						return
					}
				}
			}
		})
	}
}
