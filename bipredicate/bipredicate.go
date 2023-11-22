package bipredicate

import "github.com/CharLemAznable/gofn/common"

type BiPredicate[T any, U any] func(T, U) (bool, error)

func Of[T any, U any](fn func(T, U) (bool, error)) BiPredicate[T, U] {
	return fn
}

func Cast[T any, U any](fn func(T, U) bool) BiPredicate[T, U] {
	return func(t T, u U) (bool, error) {
		return fn(t, u), nil
	}
}

func (fn BiPredicate[T, U]) Fn(t T, u U) (bool, error) {
	return fn(t, u)
}

func (fn BiPredicate[T, U]) Test(t T, u U) bool {
	b, _ := fn(t, u)
	return b
}

func (fn BiPredicate[T, U]) MustTest(t T, u U) bool {
	b, err := fn(t, u)
	common.PanicIfError(err)
	return b
}
