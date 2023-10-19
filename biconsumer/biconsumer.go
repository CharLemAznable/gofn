package biconsumer

type BiConsumer[T any, U any] func(T, U)
type CheckedBiConsumer[T any, U any] func(T, U) error

func Of[T any, U any](fn func(T, U)) BiConsumer[T, U] {
	return fn
}

func Checked[T any, U any](fn func(T, U) error) CheckedBiConsumer[T, U] {
	return fn
}

func Unchecked[T any, U any](fn func(T, U) error) BiConsumer[T, U] {
	return func(t T, u U) {
		_ = fn(t, u)
	}
}
