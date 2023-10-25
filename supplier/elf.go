package supplier

func Constant[T any](t T) Supplier[T] {
	return func() (T, error) {
		return t, nil
	}
}
