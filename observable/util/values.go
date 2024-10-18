package util

import (
	"errors"
	"iter"

	"github.com/octohelm/go-observable/observable"
)

func Values[T any](src observable.Observable[T]) (iter.Seq2[T, error], error) {
	upstream := src.Observe()

	return func(yield func(T, error) bool) {
		defer upstream.CancelCause(nil)

		for {
			v, err, ok := upstream.Next()
			if !ok {
				return
			}

			if !yield(v, err) {
				return
			}
		}
	}, nil
}

func Collect[T any](src observable.Observable[T]) (ret []T, err error) {
	values, err := Values(src)
	if err != nil {
		return nil, err
	}

	for v, err := range values {
		if err != nil {
			if !errors.Is(err, observable.Completed) {
				return nil, err
			}
			continue
		}
		ret = append(ret, v)
	}

	return
}

func FirstValue[T any](src observable.Observable[T]) (T, error) {
	values, err := Values(src)
	if err != nil {
		return *new(T), err
	}

	for v, err := range values {
		if err != nil {
			return *new(T), err
		}
		return v, nil
	}

	return *new(T), observable.Completed
}
