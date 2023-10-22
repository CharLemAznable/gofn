package bipredicate

import (
	"testing"
)

func TestOf(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		fn       BiPredicate[int, string]
		input1   int
		input2   string
		expected bool
	}{
		{
			name:     "Test case 1",
			fn:       Of(func(i int, s string) (bool, error) { return i == 10 && s == "hello", nil }),
			input1:   10,
			input2:   "hello",
			expected: true,
		},
		{
			name:     "Test case 2",
			fn:       Of(func(i int, s string) (bool, error) { return i > 5 && len(s) > 3, nil }),
			input1:   3,
			input2:   "world",
			expected: false,
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn.Fn(tc.input1, tc.input2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestCast(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		fn       BiPredicate[int, string]
		input1   int
		input2   string
		expected bool
	}{
		{
			name:     "Test case 1",
			fn:       Cast(func(i int, s string) bool { return i == 10 && s == "hello" }),
			input1:   10,
			input2:   "hello",
			expected: true,
		},
		{
			name:     "Test case 2",
			fn:       Cast(func(i int, s string) bool { return i > 5 && len(s) > 3 }),
			input1:   3,
			input2:   "world",
			expected: false,
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn.Fn(tc.input1, tc.input2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestBiPredicate_Fn(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		fn       BiPredicate[int, string]
		input1   int
		input2   string
		expected bool
	}{
		{
			name:     "Test case 1",
			fn:       Of(func(i int, s string) (bool, error) { return i == 10 && s == "hello", nil }),
			input1:   10,
			input2:   "hello",
			expected: true,
		},
		{
			name:     "Test case 2",
			fn:       Of(func(i int, s string) (bool, error) { return i > 5 && len(s) > 3, nil }),
			input1:   3,
			input2:   "world",
			expected: false,
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn.Fn(tc.input1, tc.input2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestBiPredicate_Test(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		fn       BiPredicate[int, string]
		input1   int
		input2   string
		expected bool
	}{
		{
			name:     "Test case 1",
			fn:       Of(func(i int, s string) (bool, error) { return i == 10 && s == "hello", nil }),
			input1:   10,
			input2:   "hello",
			expected: true,
		},
		{
			name:     "Test case 2",
			fn:       Of(func(i int, s string) (bool, error) { return i > 5 && len(s) > 3, nil }),
			input1:   3,
			input2:   "world",
			expected: false,
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn.Test(tc.input1, tc.input2)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
