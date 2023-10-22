package compose

import (
	"errors"
	"strconv"
	"testing"
)

func TestThenApply(t *testing.T) {
	fn1 := func(t int) (int, error) {
		if t < 0 {
			return 0, errors.New("negative number")
		}
		return t, nil
	}

	fn2 := func(t int) (string, error) {
		return "result: " + strconv.Itoa(t), nil
	}

	errVal := "error"
	applyFn := ThenApply(fn1, fn2, errVal)

	result, err := applyFn(10)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}
	expected := "result: 10"
	if result != expected {
		t.Errorf("Expected result: %s, but got: %s", expected, result)
	}

	result, err = applyFn(-5)
	if err == nil {
		t.Error("Expected non-nil error, but got nil")
	}
	expectedErr := "negative number"
	if err.Error() != expectedErr {
		t.Errorf("Expected error: %s, but got: %v", expectedErr, err)
	}
	if result != errVal {
		t.Errorf("Expected result: %s, but got: %s", errVal, result)
	}
}

func TestTestThenSupply(t *testing.T) {
	predicateFn := func(t int) (bool, error) {
		return t > 0, nil
	}

	supplierFn := func() (int, error) {
		return 10, nil
	}

	errVal := 0
	supplyFn := TestThenSupply(predicateFn, supplierFn, errVal)

	result, err := supplyFn(5)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}
	expected := 10
	if result != expected {
		t.Errorf("Expected result: %d, but got: %d", expected, result)
	}

	result, err = supplyFn(-5)
	if err == nil {
		t.Error("Expected non-nil error, but got nil")
	}
	expectedErr := "predicate failed"
	if err.Error() != expectedErr {
		t.Errorf("Expected error: %s, but got: %v", expectedErr, err)
	}
	if result != errVal {
		t.Errorf("Expected result: %d, but got: %d", errVal, result)
	}
}
