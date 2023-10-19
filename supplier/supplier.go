package supplier

type Supplier[T any] func() T
type CheckedSupplier[T any] func() (T, error)

func Of[T any](fn func() T) Supplier[T] {
	return fn
}

func Checked[T any](fn func() (T, error)) CheckedSupplier[T] {
	return fn
}

func Unchecked[T any](fn func() (T, error)) Supplier[T] {
	return func() T {
		t, _ := fn()
		return t
	}
}
