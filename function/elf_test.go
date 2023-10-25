package function_test

import (
	"github.com/CharLemAznable/gofn/function"
	"testing"
)

func TestIdentity(t *testing.T) {
	// Test case 1
	result, err := function.Identity[int]()(5)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result to be 5, but got: %v", result)
	}

	// Test case 2
	result2, err := function.Identity[string]()("hello")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result2 != "hello" {
		t.Errorf("Expected result to be 'hello', but got: %v", result2)
	}
}

func TestYCombinator(t *testing.T) {
	// Test case 1
	fab := func(g func(int) int) func(int) int {
		return func(n int) int {
			if n == 0 {
				return 1
			}
			return n * g(n-1)
		}
	}
	result := function.YCombinator(fab)(5)
	if result != 120 {
		t.Errorf("Expected result to be 120, but got: %v", result)
	}

	// Test case 2
	fib := func(g func(int) int) func(int) int {
		return func(i int) int {
			if i <= 2 {
				return 1
			}
			return g(i-1) + g(i-2)
		}
	}
	result2 := function.YCombinator(fib)(10)
	if result2 != 55 {
		t.Errorf("Expected result to be 55, but got: %v", result2)
	}
}
