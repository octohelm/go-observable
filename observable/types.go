package observable

type Cancelable interface {
	CancelCause(err error)
}

type Observer[T any] interface {
	Next() (T, error, bool)
	Done() <-chan struct{}
	Cancelable
}

type ValueNotifier[T any] interface {
	Send(x T)
}

type Observable[T any] interface {
	Observe() Observer[T]
}

type Subscriber[T any] interface {
	ValueNotifier[T]
}
