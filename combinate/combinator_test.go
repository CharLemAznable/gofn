package combinate_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/consumer"
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/predicate"
	"github.com/CharLemAznable/gofn/runnable"
	"github.com/CharLemAznable/gofn/supplier"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunnable(t *testing.T) {
	combinator := combinate.NewCombinator()
	fn := combinate.Runnable(combinator)
	err := fn()
	assert.NoError(t, err)

	combinator.Append(runnable.Cast(func() {}))
	err = fn()
	assert.NoError(t, err)

	combinator.Append(runnable.Of(func() error {
		return errors.New("error")
	}))
	err = fn()
	assert.Error(t, err)
}

func TestSupplier(t *testing.T) {
	combinator := combinate.NewCombinator(supplier.Cast(func() int {
		return 10
	}))
	fn := combinate.Supplier[int](combinator)
	i := fn.Get()
	assert.Equal(t, 10, i)

	combinator.Append(supplier.Cast(func() string {
		return "ok"
	}))
	fn2 := combinate.Supplier[string](combinator)
	s := fn2.Get()
	assert.Equal(t, "ok", s)
}

func TestConsumer(t *testing.T) {
	combinator := combinate.NewCombinator(function.Cast(func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})).Append(consumer.Of(func(str string) error {
		return errors.New(str)
	}))
	fn := combinate.Consumer[int](combinator)

	err := fn(10)
	assert.Equal(t, "even", err.Error())

	err = fn(5)
	assert.Equal(t, "odd", err.Error())
}

func TestFunction(t *testing.T) {
	combinator := combinate.NewCombinator(predicate.Of(func(i int) (bool, error) {
		if i%2 == 0 {
			return true, nil
		}
		return false, errors.New("odd")
	})).Append(function.Cast(func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	}))
	fn := combinate.Function[int, string](combinator)

	str := fn.Apply(10)
	assert.Equal(t, "even", str)

	str, err := fn(5)
	assert.Equal(t, "", str)
	assert.Equal(t, "odd", err.Error())
}

func TestPredicate(t *testing.T) {
	combinator := combinate.NewCombinator(function.Of(strconv.Atoi)).
		Append(predicate.Cast(func(i int) bool {
			return i%2 == 0
		}))
	fn := combinate.Predicate[string](combinator)

	b := fn.Test("10")
	assert.True(t, b)

	b = fn.Test("5")
	assert.False(t, b)
}
