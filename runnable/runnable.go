package runnable

type Runnable func() error

func Of(fn func() error) Runnable {
	return fn
}

func Cast(fn func()) Runnable {
	return func() error {
		fn()
		return nil
	}
}

func (fn Runnable) Fn() {
	_ = fn()
}

func (fn Runnable) Run() {
	fn.Fn()
}
