package bifunction

type BiFunction[T any, U any, R any] func(T, U) R
type CheckedBiFunction[T any, U any, R any] func(T, U) (R, error)

func Of[T any, U any, R any](fn func(T, U) R) BiFunction[T, U, R] {
	return fn
}

func Checked[T any, U any, R any](fn func(T, U) (R, error)) CheckedBiFunction[T, U, R] {
	return fn
}

func Unchecked[T any, U any, R any](fn func(T, U) (R, error)) BiFunction[T, U, R] {
	return func(t T, u U) R {
		r, _ := fn(t, u)
		return r
	}
}
