package predicate

type Predicate[T any] func(T) bool
type CheckedPredicate[T any] func(T) (bool, error)

func Of[T any](fn func(T) bool) Predicate[T] {
	return fn
}

func Checked[T any](fn func(T) (bool, error)) CheckedPredicate[T] {
	return fn
}

func Unchecked[T any](fn func(T) (bool, error)) Predicate[T] {
	return func(t T) bool {
		b, _ := fn(t)
		return b
	}
}
