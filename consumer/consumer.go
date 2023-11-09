package consumer

import (
	"github.com/CharLemAznable/gofn/common"
)

type Consumer[T any] func(T) error

func Of[T any](fn func(T) error) Consumer[T] {
	return fn
}

func Cast[T any](fn func(T)) Consumer[T] {
	return func(t T) error {
		fn(t)
		return nil
	}
}

func (fn Consumer[T]) Accept(t T) {
	_ = fn(t)
}

func (fn Consumer[T]) MustAccept(t T) {
	common.PanicIfError(fn(t))
}

func (fn Consumer[T]) Execute(ctx common.Context) {
	t, err := common.Cast[T](ctx.Get())
	if err != nil {
		ctx.SetErr(err)
		ctx.Set(nil)
		return
	}
	err = fn(t)
	ctx.SetErr(err)
	if err != nil {
		ctx.Set(nil)
	}
}
