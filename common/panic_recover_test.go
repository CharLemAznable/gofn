package common_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/common"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, "error", actualError.Error())

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
	assert.True(t, ok)
	assert.Equal(t, "panicked with panicked", panicError.Error())
}
