package bifunction

import "github.com/CharLemAznable/gofn/function"

func (fn BiFunction[T, U, R]) Curry() func(T) function.Function[U, R] {
	return func(t T) function.Function[U, R] {
		return func(u U) (R, error) {
			return fn(t, u)
		}
	}
}

func (fn BiFunction[T, U, R]) Partial(t T) function.Function[U, R] {
	return func(u U) (R, error) {
		return fn(t, u)
	}
}
