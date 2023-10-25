package supplier_test

import (
	"github.com/CharLemAznable/gofn/supplier"
	"testing"
)

func TestConstant(t *testing.T) {
	// Test case 1: Test with integer value
	intSupplier := supplier.Constant(10)
	result, err := intSupplier()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if result != 10 {
		t.Errorf("Expected result to be 10, got %v", result)
	}

	// Test case 2: Test with string value
	strSupplier := supplier.Constant("Hello")
	result2, err := strSupplier()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if result2 != "Hello" {
		t.Errorf("Expected result to be 'Hello', got %v", result2)
	}

	// Test case 3: Test with custom struct
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John", Age: 30}
	personSupplier := supplier.Constant(person)
	result3, err := personSupplier()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if result3 != person {
		t.Errorf("Expected result to be %+v, got %+v", person, result3)
	}
}
