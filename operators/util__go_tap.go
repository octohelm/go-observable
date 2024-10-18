package operators

import "github.com/octohelm/go-observable/observable"

func GoTap[T any](do func(x T)) observable.Operator[T, T] {
	return func(src observable.Observable[T]) observable.Observable[T] {
		return observable.From(func(yield func(T, error) bool) {
			upstream := src.Observe()

			for {
				v, err, ok := upstream.Next()
				if !ok {
					return
				}

				if err != nil {
					if !yield(*new(T), err) {
						return
					}
					return
				}

				go func() {
					do(v)
				}()

				if !yield(v, nil) {
					return
				}
			}
		})
	}
}
