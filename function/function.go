package function

import (
	"github.com/CharLemAznable/gofn/common"
)

type Function[T any, R any] func(T) (R, error)

func Of[T any, R any](fn func(T) (R, error)) Function[T, R] {
	return fn
}

func Cast[T any, R any](fn func(T) R) Function[T, R] {
	return func(t T) (R, error) {
		return fn(t), nil
	}
}

func (fn Function[T, R]) Apply(t T) R {
	r, _ := fn(t)
	return r
}

func (fn Function[T, R]) MustApply(t T) R {
	r, err := fn(t)
	common.PanicIfError(err)
	return r
}

func (fn Function[T, R]) Execute(ctx common.Context) {
	t, err := common.Cast[T](ctx.Get())
	if err != nil {
		ctx.SetErr(err)
		ctx.Set(nil)
		return
	}
	r, err := fn(t)
	ctx.SetErr(err)
	ctx.Set(r)
}
