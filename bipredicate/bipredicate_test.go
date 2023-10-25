package bipredicate_test

import (
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
