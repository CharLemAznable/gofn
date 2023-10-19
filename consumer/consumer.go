package consumer

type Consumer[T any] func(T)
type CheckedConsumer[T any] func(T) error

func Of[T any](fn func(T)) Consumer[T] {
	return fn
}

func Checked[T any](fn func(T) error) CheckedConsumer[T] {
	return fn
}

func Unchecked[T any](fn func(T) error) Consumer[T] {
	return func(t T) {
		_ = fn(t)
	}
}
