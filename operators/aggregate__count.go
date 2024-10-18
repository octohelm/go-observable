package operators

import "github.com/octohelm/go-observable/observable"

func Count[T any]() observable.Operator[T, int] {
	return func(src observable.Observable[T]) observable.Observable[int] {
		return observable.From(func(yield func(int, error) bool) {
			upstream := src.Observe()

			i := 0
			for {
				_, err, ok := upstream.Next()
				if !ok {
					return
				}

				if err != nil {
					if !yield(0, err) {
						return
					}
					return
				}

				if !yield(i, nil) {
					return
				}
				i++
			}
		})
	}
}
