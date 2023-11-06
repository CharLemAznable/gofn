package common_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/common"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	finished := make(chan error)
	panicked := make(common.Panicked)

	go func() {
		defer panicked.Recover()
		finished <- errors.New("error")
	}()

	var actualError error
	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = common.WrapPanic(v)
	}
	if actualError.Error() != "error" {
		t.Errorf("Expected error message 'error', but got '%s'", actualError.Error())
	}

	go func() {
		defer panicked.Recover()
		panic("panicked")
	}()

	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = common.WrapPanic(v)
	}
	panicError, ok := (actualError).(*common.PanicError)
	if !ok {
		t.Errorf("Expected error is common.PanicError, but got %T", actualError)
	}
	if panicError.Error() != "panicked with panicked" {
		t.Errorf("Expected error message 'panicked with panicked', but got '%s'", panicError.Error())
	}
}
