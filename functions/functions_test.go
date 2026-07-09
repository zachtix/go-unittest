package functions

import "testing"

func TestAdd(t *testing.T) {
	testCase := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "Add positive numbers", a: 2, b: 3, expected: 5},
		{name: "Add negative numbers", a: -2, b: -3, expected: -5},
		{name: "Add zero", a: 0, b: 0, expected: 0},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	testCase := []struct {
		name     string
		num      int
		expected int
	}{
		{name: "Factorial of 0", num: 0, expected: 1},
		{name: "Factorial of 1", num: 1, expected: 1},
		{name: "Factorial of 5", num: 5, expected: 120},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)
			if result != tc.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", tc.num, result, tc.expected)
			}
		})
	}
}
