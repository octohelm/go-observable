package join

import (
	"errors"
	"sync"

	"github.com/octohelm/go-observable/observable"
)

func Merge[T any](sources ...observable.Observable[T]) observable.Observable[T] {
	return observable.From(func(yield func(T, error) bool) {
		chValue := make(chan T)
		chDone := make(chan struct{})
		chErr := make(chan error)

		wg := &sync.WaitGroup{}

		for _, s := range sources {
			upstream := s.Observe()

			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					v, err, ok := upstream.Next()
					if !ok {
						break
					}
					if err != nil {
						chErr <- err
						break
					}
					chValue <- v
				}
			}()
		}

		go func() {
			wg.Wait()
			close(chDone)
		}()

		var err error
		for {
			select {
			case e := <-chErr:
				err = errors.Join(err, e)
			case <-chDone:
				if err != nil {
					if !yield(*new(T), err) {
						return
					}
				}
				return
			case v := <-chValue:
				if !yield(v, nil) {
					return
				}

			}
		}
	})
}
