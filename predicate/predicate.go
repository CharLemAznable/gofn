package predicate

import (
	"github.com/CharLemAznable/gofn/common"
)

type Predicate[T any] func(T) (bool, error)

func Of[T any](fn func(T) (bool, error)) Predicate[T] {
	return fn
}

func Cast[T any](fn func(T) bool) Predicate[T] {
	return func(t T) (bool, error) {
		return fn(t), nil
	}
}

func (fn Predicate[T]) Test(t T) bool {
	b, _ := fn(t)
	return b
}

func (fn Predicate[T]) MustTest(t T) bool {
	b, err := fn(t)
	common.PanicIfError(err)
	return b
}

func (fn Predicate[T]) Execute(ctx common.Context) {
	t, err := common.Cast[T](ctx.Get())
	if err != nil {
		ctx.SetErr(err)
		ctx.Set(nil)
		return
	}
	b, err := fn(t)
	ctx.SetErr(err)
	ctx.SetInterrupt(!b)
	if !b {
		ctx.Set(nil)
	}
}
