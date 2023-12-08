package biconsumer

import "github.com/CharLemAznable/ge"

type BiConsumer[T any, U any] func(T, U) error

func Of[T any, U any](fn func(T, U) error) BiConsumer[T, U] {
	return fn
}

func Cast[T any, U any](fn func(T, U)) BiConsumer[T, U] {
	return func(t T, u U) error {
		fn(t, u)
		return nil
	}
}

func (fn BiConsumer[T, U]) Fn(t T, u U) error {
	return fn(t, u)
}

func (fn BiConsumer[T, U]) Accept(t T, u U) {
	_ = fn(t, u)
}

func (fn BiConsumer[T, U]) MustAccept(t T, u U) {
	ge.PanicIfError(fn(t, u))
}
