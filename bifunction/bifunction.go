package bifunction

import "github.com/CharLemAznable/gofn/function"

type BiFunction[T any, U any, R any] func(T, U) (R, error)

func Of[T any, U any, R any](fn func(T, U) (R, error)) BiFunction[T, U, R] {
	return fn
}

func Cast[T any, U any, R any](fn func(T, U) R) BiFunction[T, U, R] {
	return func(t T, u U) (R, error) {
		return fn(t, u), nil
	}
}

func (fn BiFunction[T, U, R]) Fn(t T, u U) R {
	r, _ := fn(t, u)
	return r
}

func (fn BiFunction[T, U, R]) Apply(t T, u U) R {
	return fn.Fn(t, u)
}

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
