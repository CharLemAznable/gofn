package predicate

type Predicate[T any] func(T) (bool, error)

func Of[T any](fn func(T) (bool, error)) Predicate[T] {
	return fn
}

func Cast[T any](fn func(T) bool) Predicate[T] {
	return func(t T) (bool, error) {
		return fn(t), nil
	}
}

func (fn Predicate[T]) Fn(t T) bool {
	b, _ := fn(t)
	return b
}

func (fn Predicate[T]) Test(t T) bool {
	return fn.Fn(t)
}
