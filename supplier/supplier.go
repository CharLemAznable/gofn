package supplier

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
