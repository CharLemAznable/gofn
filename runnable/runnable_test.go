package runnable_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/runnable"
	"testing"
)

func TestOf(t *testing.T) {
	err := errors.New("test error")
	runnableFn := runnable.Of(func() error {
		return err
	})

	resultErr := runnableFn()
	if resultErr != err {
		t.Errorf("Expected error: %v, but got: %v", err, resultErr)
	}
}

func TestCast(t *testing.T) {
	fnCalled := false
	runnableFn := runnable.Cast(func() {
		fnCalled = true
	})

	resultErr := runnableFn()
	if resultErr != nil {
		t.Errorf("Expected nil error, but got: %v", resultErr)
	}

	if !fnCalled {
		t.Error("Expected function to be called, but it was not")
	}
}

func TestRunnable_Run(t *testing.T) {
	fnCalled := false
	runnableFn := runnable.Runnable(func() error {
		fnCalled = true
		return nil
	})

	runnableFn.Run()

	if !fnCalled {
		t.Error("Expected function to be called, but it was not")
	}
}
