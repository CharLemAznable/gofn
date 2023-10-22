package supplier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	// Test case 1: Test Of function with a function that returns an integer
	fn := func() (int, error) {
		return 10, nil
	}
	s := Of(fn)
	result, err := s()
	assert.NoError(t, err)
	assert.Equal(t, 10, result)

	// Test case 2: Test Of function with a function that returns a string
	fn2 := func() (string, error) {
		return "hello", nil
	}
	s2 := Of(fn2)
	result2, err2 := s2()
	assert.NoError(t, err2)
	assert.Equal(t, "hello", result2)
}

func TestCast(t *testing.T) {
	// Test case 1: Test Cast function with a function that returns an integer
	fn := func() int {
		return 10
	}
	s := Cast(fn)
	result, err := s()
	assert.NoError(t, err)
	assert.Equal(t, 10, result)

	// Test case 2: Test Cast function with a function that returns a string
	fn2 := func() string {
		return "hello"
	}
	s2 := Cast(fn2)
	result2, err2 := s2()
	assert.NoError(t, err2)
	assert.Equal(t, "hello", result2)
}

func TestSupplier_Fn(t *testing.T) {
	// Test case 1: Test Fn method with a function that returns an integer
	fn := func() (int, error) {
		return 10, nil
	}
	s := Of(fn)
	result := s.Fn()
	assert.Equal(t, 10, result)

	// Test case 2: Test Fn method with a function that returns a string
	fn2 := func() (string, error) {
		return "hello", nil
	}
	s2 := Of(fn2)
	result2 := s2.Fn()
	assert.Equal(t, "hello", result2)
}

func TestSupplier_Get(t *testing.T) {
	// Test case 1: Test Get method with a function that returns an integer
	fn := func() (int, error) {
		return 10, nil
	}
	s := Of(fn)
	result := s.Get()
	assert.Equal(t, 10, result)

	// Test case 2: Test Get method with a function that returns a string
	fn2 := func() (string, error) {
		return "hello", nil
	}
	s2 := Of(fn2)
	result2 := s2.Get()
	assert.Equal(t, "hello", result2)
}
