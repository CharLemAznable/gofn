package biconsumer_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/biconsumer"
	"testing"
)

func TestOf(t *testing.T) {
	fn := func(i int, s string) error {
		return nil
	}

	con := biconsumer.Of(fn)
	err := con(10, "test")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCast(t *testing.T) {
	fn := func(i int, s string) {
		// do something
	}

	con := biconsumer.Cast(fn)
	err := con(10, "test")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestBiConsumer_Accept(t *testing.T) {
	fn := func(i int, s string) error {
		return errors.New("error")
	}

	con := biconsumer.Of(fn)
	err := biconsumer.Cast(con.Accept).Fn(10, "test")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

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
		con.MustAccept(10, "test")
	}()
}
