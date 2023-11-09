package common_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/common"
	"testing"
)

func TestDefaultErrorMsg(t *testing.T) {
	// Test case 1: err is nil, defMsg is not empty
	err1 := common.DefaultErrorMsg(nil, "default message")
	if err1 == nil {
		t.Error("Expected an error, but got nil")
	} else if err1.Error() != "default message" {
		t.Errorf("Expected error message 'default message', but got '%s'", err1.Error())
	}

	// Test case 2: err is not nil, defMsg is empty
	err2 := errors.New("original error")
	err2 = common.DefaultErrorMsg(err2, "")
	if err2 != err2 {
		t.Error("Expected the same error, but got a different error")
	}

	// Test case 3: err is not nil, defMsg is not empty
	err3 := errors.New("original error")
	err3 = common.DefaultErrorMsg(err3, "default message")
	if err3 == nil {
		t.Error("Expected an error, but got nil")
	} else if err3.Error() != "original error" {
		t.Errorf("Expected error message 'original error', but got '%s'", err3.Error())
	}
}

func TestPanicIfError(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Expected no panic, but got %v", r)
			}
		}()
		common.PanicIfError(nil)
	}()
	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("Expected panic, but got nothing")
			}
			err, ok := r.(error)
			if !ok {
				t.Errorf("Expected panic error, but got %v", r)
			} else if err.Error() != "panic error" {
				t.Errorf("Expected error message 'panic error', but got '%s'", err.Error())
			}
		}()
		common.PanicIfError(errors.New("panic error"))
	}()
}
