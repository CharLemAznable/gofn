package biconsumer

import (
	"testing"
)

func TestOf(t *testing.T) {
	fn := func(t int, u string) error {
		// test implementation
		return nil
	}

	bc := Of(fn)

	err := bc(10, "test")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestCast(t *testing.T) {
	fn := func(t int, u string) {
		// test implementation
	}

	bc := Cast(fn)

	err := bc(10, "test")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestFn(t *testing.T) {
	fn := func(t int, u string) error {
		// test implementation
		return nil
	}

	bc := Of(fn)

	bc.Fn(10, "test")
}

func TestAccept(t *testing.T) {
	fn := func(t int, u string) error {
		// test implementation
		return nil
	}

	bc := Of(fn)

	bc.Accept(10, "test")
}
