package bifunction_test

import (
	"github.com/CharLemAznable/gofn/bifunction"
	"testing"
)

func TestOf(t *testing.T) {
	fn := bifunction.Of(func(a int, b string) (int, error) {
		return a + len(b), nil
	})

	result, err := fn(5, "hello")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCast(t *testing.T) {
	fn := bifunction.Cast(func(a int, b string) int {
		return a + len(b)
	})

	result := fn.Apply(5, "hello")

	expected := 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
