package function

type Function[T any, R any] func(T) R
type CheckedFunction[T any, R any] func(T) (R, error)

func Of[T any, R any](fn func(T) R) Function[T, R] {
	return fn
}

func Checked[T any, R any](fn func(T) (R, error)) CheckedFunction[T, R] {
	return fn
}

func Unchecked[T any, R any](fn func(T) (R, error)) Function[T, R] {
	return func(t T) R {
		r, _ := fn(t)
		return r
	}
}
