package bifunction_test

import (
	"github.com/CharLemAznable/gofn/bifunction"
	"testing"
)

func TestCurry(t *testing.T) {
	fn := bifunction.Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	curriedFn := fn.Curry()(1)
	r, err := curriedFn("test")
	if r != true || err != nil {
		t.Errorf("TestCurry failed")
	}
}

func TestPartial(t *testing.T) {
	fn := bifunction.Of(func(t int, u string) (bool, error) {
		return true, nil
	})

	partialFn := fn.Partial(1)
	r, err := partialFn("test")
	if r != true || err != nil {
		t.Errorf("TestPartial failed")
	}
}
