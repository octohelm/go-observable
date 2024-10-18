package observable

func Of[T any](values ...T) Observable[T] {
	return From(func(yield func(T, error) bool) {
		for _, v := range values {
			if !yield(v, nil) {
				return
			}
		}
	})
}
