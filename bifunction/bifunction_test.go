package bifunction

import (
	"testing"
)

func TestOf(t *testing.T) {
	fn := Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	r, err := fn(1, "test")
	if r != true || err != nil {
		t.Errorf("TestOf failed")
	}
}

func TestCast(t *testing.T) {
	fn := Cast(func(t int, u string) bool {
		return true
	})

	r, err := fn(1, "test")
	if r != true || err != nil {
		t.Errorf("TestCast failed")
	}
}

func TestFn(t *testing.T) {
	fn := Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	r := fn.Fn(1, "test")
	if r != true {
		t.Errorf("TestFn failed")
	}
}

func TestApply(t *testing.T) {
	fn := Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	r := fn.Apply(1, "test")
	if r != true {
		t.Errorf("TestApply failed")
	}
}

func TestCurry(t *testing.T) {
	fn := Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	curriedFn := fn.Curry()(1)
	r, err := curriedFn("test")
	if r != true || err != nil {
		t.Errorf("TestCurry failed")
	}
}

func TestPartial(t *testing.T) {
	fn := Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	partialFn := fn.Partial(1)
	r, err := partialFn("test")
	if r != true || err != nil {
		t.Errorf("TestPartial failed")
	}
}
