package observable

type Operator[I any, O any] func(source Observable[I]) Observable[O]

func Pipe[T any](source Observable[T], operators ...Operator[T, T]) Observable[T] {
	ret := source
	for _, op := range operators {
		ret = op(ret)
	}
	return ret
}

func Pipe2[S, O any](source Observable[S], operator Operator[S, O]) Observable[O] {
	return operator(source)
}

func Pipe3[S, A, O any](source Observable[S], operator0 Operator[S, A], operator1 Operator[A, O]) Observable[O] {
	return operator1(operator0(source))
}

func Pipe4[S, A, B, O any](source Observable[S], operator0 Operator[S, A], operator1 Operator[A, B], operator2 Operator[B, O]) Observable[O] {
	return operator2(operator1(operator0(source)))
}

func Pipe5[S, A, B, C, O any](source Observable[S], operator0 Operator[S, A], operator1 Operator[A, B], operator2 Operator[B, C], operator3 Operator[C, O]) Observable[O] {
	return operator3(operator2(operator1(operator0(source))))
}
