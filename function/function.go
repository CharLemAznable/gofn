package function

type Function[T any, R any] func(T) (R, error)

func Of[T any, R any](fn func(T) (R, error)) Function[T, R] {
	return fn
}

func Cast[T any, R any](fn func(T) R) Function[T, R] {
	return func(t T) (R, error) {
		return fn(t), nil
	}
}

func (fn Function[T, R]) Fn(t T) R {
	r, _ := fn(t)
	return r
}

func (fn Function[T, R]) Apply(t T) R {
	return fn.Fn(t)
}

func Identity[T any]() Function[T, T] {
	return Of(func(t T) (T, error) {
		return t, nil
	})
}

func YCombinator[T any](f func(func(T) T) func(T) T) func(T) T {
	return func(t T) T {
		return f(YCombinator(f))(t)
	}
}
