package bifunction_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/bifunction"
	"testing"
)

func TestOf(t *testing.T) {
	fn := bifunction.Of(func(a int, b string) (int, error) {
		return a + len(b), nil
	})

	result, err := fn.Fn(5, "hello")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 10
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	fn = bifunction.Of(func(a int, b string) (int, error) {
		return a + len(b), errors.New("error")
	})
	func() {
		defer func() {
			rec := recover()
			if rec == nil {
				t.Error("Expected recover error, but got nil")
			}
			recErr, ok := rec.(error)
			if !ok {
				t.Errorf("Expected recover error, but got %v", rec)
			} else if recErr.Error() != "error" {
				t.Errorf("Expected error message 'error', but got '%s'", recErr.Error())
			}
		}()
		fn.MustApply(5, "hello")
	}()
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
