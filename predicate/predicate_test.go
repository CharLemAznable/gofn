package predicate_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/predicate"
	"testing"
)

func TestOf(t *testing.T) {
	// Test case 1: fn returns true
	fn := predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result, err := fn(5)
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: fn returns false
	fn = predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result, err = fn(-5)
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}
	if result {
		t.Errorf("Expected false, but got true")
	}

	// Test case 3: fn returns error
	expectedErr := "Some error occurred"
	fn = predicate.Of(func(n int) (bool, error) {
		return false, errors.New(expectedErr)
	})
	result, err = fn(5)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErr, err.Error())
	}
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestCast(t *testing.T) {
	// Test case 1: fn returns true
	fn := predicate.Cast(func(n int) bool {
		return n > 0
	})
	result, err := fn(5)
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: fn returns false
	fn = predicate.Cast(func(n int) bool {
		return n > 0
	})
	result, err = fn(-5)
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestPredicateFn(t *testing.T) {
	// Test case 1: fn returns true
	fn := predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result := fn.Fn(5)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: fn returns false
	fn = predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result = fn.Fn(-5)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestPredicateTest(t *testing.T) {
	// Test case 1: fn returns true
	fn := predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result := fn.Test(5)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: fn returns false
	fn = predicate.Of(func(n int) (bool, error) {
		return n > 0, nil
	})
	result = fn.Test(-5)
	if result {
		t.Errorf("Expected false, but got true")
	}
}
