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


func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1, 5, 10, 25}, 30, 2},
		{[]int{186, 419, 83, 408}, 6249, 20},
		{[]int{3, 7, 405, 436}, 8839, 25},
	}

	for _, test := range tests {
		result := coinChange(test.coins, test.target)
		if result != test.expected {
			t.Errorf("coinChange(%v, %d) returned %d, expected %d", test.coins, test.target, result, test.expected)
		}
	}
}
