package compose

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunInSequence(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	fn1 := func() error { return nil }
	fn2 := func() error { return err1 }
	fn3 := func() error { return err2 }
	fn4 := func() error { return err3 }

	fn := RunInSequence(fn1, fn2, fn3, fn4)
	err := fn()

	assert.Equal(t, err1, err)

	fn = RunInSequence(fn1)
	err = fn()
	assert.NoError(t, err)
}

func TestSupplyThenConsume(t *testing.T) {
	e := errors.New("error")

	supplierFn := func() (int, error) { return 42, nil }
	consumerFn := func(t int) error { return nil }

	fn := SupplyThenConsume(supplierFn, consumerFn)
	err := fn()

	assert.NoError(t, err)

	supplierFn = func() (int, error) { return 0, e }
	fn = SupplyThenConsume(supplierFn, consumerFn)
	err = fn()

	assert.Equal(t, e, err)
}
