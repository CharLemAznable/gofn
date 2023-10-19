package unaryoperator

type UnaryOperator[T any] func(T) T
type CheckedUnaryOperator[T any] func(T) (T, error)

func Of[T any](fn func(T) T) UnaryOperator[T] {
	return fn
}

func Checked[T any](fn func(T) (T, error)) CheckedUnaryOperator[T] {
	return fn
}

func Unchecked[T any](fn func(T) (T, error)) UnaryOperator[T] {
	return func(t T) T {
		u, _ := fn(t)
		return u
	}
}
