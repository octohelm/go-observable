package util

import (
	"github.com/octohelm/go-observable/observable"
)

func Wait[T any](src observable.Observable[T]) error {
	ob := src.Observe()

	for {
		_, err, ok := ob.Next()
		if !ok {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func WaitUnit[T any](done <-chan struct{}, src observable.Observable[T]) error {
	ob := src.Observe()
	defer ob.CancelCause(nil)
	chErr := make(chan error)

	go func() {
		for {
			_, err, ok := ob.Next()
			if !ok {
				chErr <- nil
				return
			}
			if err != nil {
				chErr <- err
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return observable.Completed
		case err := <-chErr:
			return err
		}
	}
}
