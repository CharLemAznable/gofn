package compose_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/compose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyThenConsume(t *testing.T) {
	// Mock functionFn and consumerFn
	functionFn := func(t interface{}) (interface{}, error) {
		return t, nil
	}
	consumerFn := func(r interface{}) error {
		return nil
	}

	// Test case 1: functionFn and consumerFn return no error
	err := compose.ApplyThenConsume(functionFn, consumerFn)(1)
	assert.NoError(t, err)

	// Test case 2: functionFn returns an error
	functionFn = func(t interface{}) (interface{}, error) {
		return nil, errors.New("function error")
	}
	err = compose.ApplyThenConsume(functionFn, consumerFn)(1)
	assert.EqualError(t, err, "function error")

	// Test case 3: consumerFn returns an error
	functionFn = func(t interface{}) (interface{}, error) {
		return t, nil
	}
	consumerFn = func(r interface{}) error {
		return errors.New("consumer error")
	}
	err = compose.ApplyThenConsume(functionFn, consumerFn)(1)
	assert.EqualError(t, err, "consumer error")
}

func TestCheckThenRun(t *testing.T) {
	// Mock predicateFn and runnableFn
	predicateFn := func(t interface{}) (bool, error) {
		return true, nil
	}
	runnableFn := func() error {
		return nil
	}

	// Test case 1: predicateFn returns true and no error
	err := compose.CheckThenRun(predicateFn, runnableFn)(1)
	assert.NoError(t, err)

	// Test case 2: predicateFn returns false
	predicateFn = func(t interface{}) (bool, error) {
		return false, nil
	}
	err = compose.CheckThenRun(predicateFn, runnableFn)(1)
	assert.EqualError(t, err, "predicate failed")

	// Test case 3: predicateFn returns an error
	predicateFn = func(t interface{}) (bool, error) {
		return true, errors.New("predicate error")
	}
	err = compose.CheckThenRun(predicateFn, runnableFn)(1)
	assert.EqualError(t, err, "predicate error")

	// Test case 4: runnableFn returns an error
	predicateFn = func(t interface{}) (bool, error) {
		return true, nil
	}
	runnableFn = func() error {
		return errors.New("runnable error")
	}
	err = compose.CheckThenRun(predicateFn, runnableFn)(1)
	assert.EqualError(t, err, "runnable error")
}

func TestCheckThenConsume(t *testing.T) {
	// Mock predicateFn and consumerFn
	predicateFn := func(t interface{}) (bool, error) {
		return true, nil
	}
	consumerFn := func(t interface{}) error {
		return nil
	}

	// Test case 1: predicateFn returns true and no error
	err := compose.CheckThenConsume(predicateFn, consumerFn)(1)
	assert.NoError(t, err)

	// Test case 2: predicateFn returns false
	predicateFn = func(t interface{}) (bool, error) {
		return false, nil
	}
	err = compose.CheckThenConsume(predicateFn, consumerFn)(1)
	assert.EqualError(t, err, "predicate failed")

	// Test case 3: predicateFn returns an error
	predicateFn = func(t interface{}) (bool, error) {
		return true, errors.New("predicate error")
	}
	err = compose.CheckThenConsume(predicateFn, consumerFn)(1)
	assert.EqualError(t, err, "predicate error")

	// Test case 4: consumerFn returns an error
	predicateFn = func(t interface{}) (bool, error) {
		return true, nil
	}
	consumerFn = func(t interface{}) error {
		return errors.New("runnable error")
	}
	err = compose.CheckThenConsume(predicateFn, consumerFn)(1)
	assert.EqualError(t, err, "runnable error")
}
