package observable

import (
	"errors"
	"iter"

	"github.com/octohelm/go-observable/internal"
)

var Completed = internal.Completed

func From[T any](seq iter.Seq2[T, error]) Observable[T] {
	return &observable[T]{
		seq: seq,
	}
}

type observable[T any] struct {
	seq iter.Seq2[T, error]
}

func (o *observable[T]) Observe() Observer[T] {
	next, stop := iter.Pull2(o.seq)

	s := &stream[T]{
		next: next,
	}

	s.OnComplete = func() {
		go stop()
	}

	s.Init()

	return s
}

type stream[T any] struct {
	next func() (T, error, bool)
	internal.Completer
}

func (o *stream[T]) Next() (T, error, bool) {
	v, err, ok := o.next()
	if err != nil && errors.Is(err, Completed) {
		return v, nil, false
	}
	return v, err, ok
}
