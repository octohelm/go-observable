package observable

func Empty[T any]() Observable[T] {
	return From(func(yield func(T, error) bool) {})
}
