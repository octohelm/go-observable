package internal

type Cancelable interface {
	CancelCause(err error)
}

type Observer[T any] interface {
	Next() (T, error, bool)
	Done() <-chan struct{}
	Cancelable
}
