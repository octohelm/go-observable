package operators

import "github.com/octohelm/go-observable/observable"

func IgnoreElements[T any]() observable.Operator[T, struct{}] {
	return func(src observable.Observable[T]) observable.Observable[struct{}] {
		return observable.From(func(yield func(struct{}, error) bool) {
			upstream := src.Observe()

			for {
				_, err, ok := upstream.Next()
				if !ok {
					return
				}

				if err != nil {
					if !yield(struct{}{}, err) {
						return
					}
					return
				}
			}
		})
	}
}
