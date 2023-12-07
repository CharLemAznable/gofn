package supplier

import "github.com/CharLemAznable/gofn/common"

type Supplier[T any] func() (T, error)

func Of[T any](fn func() (T, error)) Supplier[T] {
	return fn
}

func Cast[T any](fn func() T) Supplier[T] {
	return func() (T, error) {
		return fn(), nil
	}
}

func (fn Supplier[T]) Fn() (T, error) {
	return fn()
}

func (fn Supplier[T]) Get() T {
	t, _ := fn()
	return t
}

func (fn Supplier[T]) MustGet() T {
	t, err := fn()
	common.PanicIfError(err)
	return t
}

func (fn Supplier[T]) Execute(ctx common.Context) {
	t, err := fn()
	ctx.SetErr(err)
	ctx.Set(t)
}
