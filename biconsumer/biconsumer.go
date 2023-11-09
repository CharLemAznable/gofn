package biconsumer

import "github.com/CharLemAznable/gofn/common"

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

func (fn BiConsumer[T, U]) Accept(t T, u U) {
	_ = fn(t, u)
}

func (fn BiConsumer[T, U]) MustAccept(t T, u U) {
	common.PanicIfError(fn(t, u))
}
