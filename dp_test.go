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

func TestLCS(t *testing.T) {
	tests := []struct {
		X        string
		Y        string
		expected string
	}{
		{"ABCDGH", "AEDFHR", "ADH"},
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"ABCDEF", "XYZ", ""},
		{"", "XYZ", ""},
		{"", "", ""},
	}

	for _, test := range tests {
		result := LCS(test.X, test.Y)
		if result != test.expected {
			t.Errorf("LCS(%s, %s) returned %s, expected %s", test.X, test.Y, result, test.expected)
		}
	}
}

func TestKnapsack(t *testing.T) {
	tests := []struct {
		weights           []int
		values            []int
		knapsack_capacity int
		expected          int
	}{
		{[]int{}, []int{}, 10, 0},
		{[]int{2, 3, 4}, []int{5, 6, 7}, 0, 0},
		{[]int{6}, []int{10}, 5, 0},
		{[]int{2, 3, 4}, []int{5, 6, 7}, 10, 18},
		{[]int{0, 2, 0, 3}, []int{0, 5, 0, 6}, 5, 11},
		{[]int{4, 5, 2, 1, 7, 3, 6}, []int{10, 12, 6, 4, 15, 8, 11}, 20, 40},
	}

	for _, test := range tests {
		result := knapsack(test.weights, test.values, test.knapsack_capacity)
		if result != test.expected {
			t.Errorf("coinChange(%v, %d) returned %d, expected %d", test.weights, test.values, result, test.expected)
		}
	}
}

func TestMaxSumSubArray(t *testing.T) {
	tests := []struct {
		array    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{-5, -3, -8, -12, -9}, -3},
		{[]int{2, 4, 7, 1, 3, 6}, 23},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{5}, 5},
		{[]int{-3, -2, 1, -5, 4, 0, -1, 2}, 4},
		{[]int{1, -2, 3, 4, -1, 2, 1, -5, 4, 6, -8, 10}, 16},
	}

	for _, test := range tests {
		result := maxSumSubarray(test.array)
		if result != test.expected {
			t.Errorf("maxSumSubArray(%v, %d) returned %d, expected %d", test.array, result, test.expected)
		}
	}
}
