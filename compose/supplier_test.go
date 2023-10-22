package compose

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSupplyThenApply(t *testing.T) {
	// Test case 1: supplierFn returns a value, functionFn returns a value
	supplierFn1 := func() (int, error) {
		return 10, nil
	}
	functionFn1 := func(t int) (string, error) {
		return "Result: " + strconv.Itoa(t), nil
	}
	errVal1 := "Error"
	expectedResult1 := "Result: 10"

	supplier1 := SupplyThenApply(supplierFn1, functionFn1, errVal1)
	result1, err1 := supplier1()
	if err1 != nil {
		t.Errorf("Test case 1: Expected no error, got %v", err1)
	}
	if result1 != expectedResult1 {
		t.Errorf("Test case 1: Expected %s, got %s", expectedResult1, result1)
	}

	// Test case 2: supplierFn returns an error
	supplierFn2 := func() (int, error) {
		return 0, fmt.Errorf("supplier error")
	}
	functionFn2 := func(t int) (string, error) {
		return "Result: " + strconv.Itoa(t), nil
	}
	errVal2 := "Error"
	expectedError2 := "supplier error"

	supplier2 := SupplyThenApply(supplierFn2, functionFn2, errVal2)
	result2, err2 := supplier2()
	if result2 != errVal2 {
		t.Errorf("Test case 2: Expected result %s, got %s", errVal2, result2)
	}
	if err2 == nil {
		t.Errorf("Test case 2: Expected an error, got nil")
	}
	if err2.Error() != expectedError2 {
		t.Errorf("Test case 2: Expected %s, got %s", expectedError2, err2.Error())
	}

	// Test case 3: functionFn returns an error
	supplierFn3 := func() (int, error) {
		return 10, nil
	}
	functionFn3 := func(t int) (string, error) {
		return "", fmt.Errorf("function error")
	}
	errVal3 := "Error"
	expectedError3 := "function error"

	supplier3 := SupplyThenApply(supplierFn3, functionFn3, errVal3)
	result3, err3 := supplier3()
	if result3 == errVal3 {
		t.Errorf("Test case 3: Expected result is not %s", errVal3)
	}
	if result3 != "" {
		t.Errorf("Test case 3: Expected result is \"\", got %s", result3)
	}
	if err3 == nil {
		t.Errorf("Test case 3: Expected an error, got nil")
	}
	if err3.Error() != expectedError3 {
		t.Errorf("Test case 3: Expected %s, got %s", expectedError3, err3.Error())
	}
}
