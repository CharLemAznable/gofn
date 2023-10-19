package runnable

type Runnable func()
type CheckedRunnable func() error

func Of(fn func()) Runnable {
	return fn
}

func Checked(fn func() error) CheckedRunnable {
	return fn
}

func Unchecked(fn func() error) Runnable {
	return func() {
		_ = fn()
	}
}
