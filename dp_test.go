package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := fibonacci(test.n)
		if result != test.expected {
			t.Errorf("fibonacci(%d) returned %d, expected %d", test.n, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, test := range tests {
		result := factorial(test.n)
		if result != test.expected {
			t.Errorf("fibonacci(%d) returned %d, expected %d", test.n, result, test.expected)
		}
	}
}
