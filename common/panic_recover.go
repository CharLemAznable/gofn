package common

import "fmt"

type Panicked chan any

func (pr Panicked) Recover() {
	if err := recover(); err != nil {
		pr <- err
	}
}

func (pr Panicked) Caught() <-chan any {
	return pr
}

func WrapPanic(v any) error {
	return &PanicError{error: v}
}

type PanicError struct {
	error any
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panicked with %v", e.error)
}
