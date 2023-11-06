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
)

func TestRunnable(t *testing.T) {
	combinator := combinate.NewCombinator()
	fn := combinate.Runnable(combinator)
	err := fn()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	combinator.Append(runnable.Cast(func() {}))
	err = fn()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	combinator.Append(runnable.Of(func() error {
		return errors.New("error")
	}))
	err = fn()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestSupplier(t *testing.T) {
	combinator := combinate.NewCombinator(supplier.Cast(func() int {
		return 10
	}))
	fn := combinate.Supplier[int](combinator)
	i := fn.Get()
	if i != 10 {
		t.Errorf("Expected 10, but got %d", i)
	}

	combinator.Append(supplier.Cast(func() string {
		return "ok"
	}))
	fn2 := combinate.Supplier[string](combinator)
	s := fn2.Get()
	if s != "ok" {
		t.Errorf("Expected 'ok', but got '%s'", s)
	}
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
	if err.Error() != "even" {
		t.Errorf("Expected 'even', but got '%s'", err.Error())
	}

	err = fn(5)
	if err.Error() != "odd" {
		t.Errorf("Expected 'odd', but got '%s'", err.Error())
	}
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
	if str != "even" {
		t.Errorf("Expected 'even', but got '%s'", str)
	}

	str, err := fn(5)
	if str != "" {
		t.Errorf("Expected '', but got '%s'", str)
	}
	if err.Error() != "odd" {
		t.Errorf("Expected 'odd', but got '%s'", err.Error())
	}
}

func TestPredicate(t *testing.T) {
	combinator := combinate.NewCombinator(function.Of(strconv.Atoi)).
		Append(predicate.Cast(func(i int) bool {
			return i%2 == 0
		}))
	fn := combinate.Predicate[string](combinator)

	b := fn.Test("10")
	if !b {
		t.Errorf("Expected true, but got false")
	}

	b = fn.Test("5")
	if b {
		t.Errorf("Expected false, but got true")
	}
}
