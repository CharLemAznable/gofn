package errs

import (
	"errors"
	"testing"
)

func TestDefaultError(t *testing.T) {
	// Test case 1: err is nil, defMsg is not empty
	err1 := DefaultError(nil, "default message")
	if err1 == nil {
		t.Error("Expected an error, but got nil")
	} else if err1.Error() != "default message" {
		t.Errorf("Expected error message 'default message', but got '%s'", err1.Error())
	}

	// Test case 2: err is not nil, defMsg is empty
	err2 := errors.New("original error")
	err2 = DefaultError(err2, "")
	if err2 != err2 {
		t.Error("Expected the same error, but got a different error")
	}

	// Test case 3: err is not nil, defMsg is not empty
	err3 := errors.New("original error")
	err3 = DefaultError(err3, "default message")
	if err3 == nil {
		t.Error("Expected an error, but got nil")
	} else if err3.Error() != "original error" {
		t.Errorf("Expected error message 'original error', but got '%s'", err3.Error())
	}
}
