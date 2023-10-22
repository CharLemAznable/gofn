package consumer

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

func (fn Consumer[T]) Fn(t T) {
	_ = fn(t)
}

func (fn Consumer[T]) Accept(t T) {
	fn.Fn(t)
}
