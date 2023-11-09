package bipredicate_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/bipredicate"
	"testing"
)

func TestOf(t *testing.T) {
	fn := func(a int, b string) (bool, error) {
		// Test logic here
		return true, nil
	}

	predicate := bipredicate.Of(fn)

	result, err := predicate(10, "test")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !result {
		t.Error("Expected true, but got false")
	}

	fn = func(a int, b string) (bool, error) {
		return false, errors.New("error")
	}

	func() {
		defer func() {
			rec := recover()
			if rec == nil {
				t.Error("Expected recover error, but got nil")
			}
			recErr, ok := rec.(error)
			if !ok {
				t.Errorf("Expected recover error, but got %v", rec)
			} else if recErr.Error() != "error" {
				t.Errorf("Expected error message 'error', but got '%s'", recErr.Error())
			}
		}()
		bipredicate.Of(fn).MustTest(10, "test")
	}()
}

func TestCast(t *testing.T) {
	fn := func(a int, b string) bool {
		// Test logic here
		return true
	}

	predicate := bipredicate.Cast(fn)

	result := predicate.Test(10, "test")
	if !result {
		t.Error("Expected true, but got false")
	}
}
