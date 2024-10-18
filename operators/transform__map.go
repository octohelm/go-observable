package operators

import (
	"github.com/octohelm/go-observable/observable"
)

func Map[I any, O any](project func(x I) O) observable.Operator[I, O] {
	return func(src observable.Observable[I]) observable.Observable[O] {
		return observable.From(func(yield func(O, error) bool) {
			upstream := src.Observe()

			for {
				v, err, ok := upstream.Next()
				if !ok {
					return
				}

				if !yield(project(v), err) {
					return
				}
			}
		})
	}
}
