package bipredicate

type BiPredicate[T any, U any] func(T, U) bool
type CheckedBiPredicate[T any, U any] func(T, U) (bool, error)

func Of[T any, U any](fn func(T, U) bool) BiPredicate[T, U] {
	return fn
}

func Checked[T any, U any](fn func(T, U) (bool, error)) CheckedBiPredicate[T, U] {
	return fn
}

func Unchecked[T any, U any](fn func(T, U) (bool, error)) BiPredicate[T, U] {
	return func(t T, u U) bool {
		b, _ := fn(t, u)
		return b
	}
}
