package operators

import (
	"github.com/octohelm/go-observable/observable"
)

func SwitchMap[I any, O any](src observable.Observable[I], project func(x I) observable.Observable[O]) observable.Observable[O] {
	return observable.From(func(yield func(O, error) bool) {
		upstream := src.Observe()

		for {
			v, err, ok := upstream.Next()
			if !ok {
				break
			}

			if err != nil {
				if !yield(*new(O), err) {
					return
				}
				break
			}

			inner := project(v).Observe()

			for {
				x, err, ok := inner.Next()
				if !ok {
					break
				}

				if err != nil {
					if !yield(*new(O), err) {
						return
					}
					break
				}

				if !yield(x, nil) {
					return
				}
			}
		}
	})
}
