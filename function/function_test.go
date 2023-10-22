package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf(t *testing.T) {
	fn := Of(func(t int) (int, error) {
		return t * 2, nil
	})

	result, err := fn(5)

	assert.NoError(t, err)
	assert.Equal(t, 10, result)
}

func TestCast(t *testing.T) {
	fn := Cast(func(t int) int {
		return t * 2
	})

	result, err := fn(5)

	assert.NoError(t, err)
	assert.Equal(t, 10, result)
}

func TestFn(t *testing.T) {
	fn := Of(func(t int) (int, error) {
		return t * 2, nil
	})

	result := fn.Fn(5)

	assert.Equal(t, 10, result)
}

func TestApply(t *testing.T) {
	fn := Of(func(t int) (int, error) {
		return t * 2, nil
	})

	result := fn.Apply(5)

	assert.Equal(t, 10, result)
}

func TestIdentity(t *testing.T) {
	fn := Identity[int]()

	result, err := fn(5)

	assert.NoError(t, err)
	assert.Equal(t, 5, result)
}

func TestYCombinator(t *testing.T) {
	f1 := func(f func(int) int) func(int) int {
		return func(n int) int {
			if n <= 1 {
				return 1
			}
			return n * f(n-1)
		}
	}

	fac := YCombinator(f1)

	result := fac(5)

	assert.Equal(t, 120, result)

	f2 := func(f func(int) int) func(int) int {
		return func(n int) int {
			if n <= 2 {
				return 1
			}
			return f(n-1) + f(n-2)
		}
	}

	fib := YCombinator(f2)

	result2 := fib(10)

	assert.Equal(t, 55, result2)
}
