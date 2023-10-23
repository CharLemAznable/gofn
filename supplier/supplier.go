package supplier

import (
	"github.com/CharLemAznable/gofn/runnable"
)

type Supplier[T any] func() (T, error)

func Of[T any](fn func() (T, error)) Supplier[T] {
	return fn
}

func Cast[T any](fn func() T) Supplier[T] {
	return func() (T, error) {
		return fn(), nil
	}
}

func (fn Supplier[T]) Fn() T {
	t, _ := fn()
	return t
}

func (fn Supplier[T]) Get() T {
	return fn.Fn()
}

func (fn Supplier[T]) To(consumerFn func(T) error) runnable.Runnable {
	return func() error {
		t, err := fn()
		if err != nil {
			return err
		}
		return consumerFn(t)
	}
}

func Constant[T any](t T) Supplier[T] {
	return func() (T, error) {
		return t, nil
	}
}
