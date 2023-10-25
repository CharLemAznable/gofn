package function

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
