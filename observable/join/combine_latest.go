package join

import (
	"github.com/octohelm/go-observable/internal"
	"github.com/octohelm/go-observable/observable"
)

func CombineLatest2[A any, B any, O any](sourceA observable.Observable[A], sourceB observable.Observable[B], project func(a A, b B) O) observable.Observable[O] {
	return observable.From(func(yield func(O, error) bool) {
		va := &internal.Latest[A]{}
		vb := &internal.Latest[B]{}
		chChanges := make(chan struct{})

		wg := &internal.Group{}

		wg.Go(func() error {
			return va.Watch(sourceA.Observe(), chChanges)
		})

		wg.Go(func() error {
			return vb.Watch(sourceB.Observe(), chChanges)
		})

		go wg.Wait()

		for {
			select {
			case <-wg.Done():
				if err := wg.Err(); err != nil {
					if !yield(*new(O), err) {
						return
					}
				}
				return
			case <-chChanges:
				a := va.Value()
				if a == nil {
					continue
				}
				b := vb.Value()
				if b == nil {
					continue
				}

				if !yield(project(*a, *b), nil) {
					return
				}
			}
		}
	})
}
