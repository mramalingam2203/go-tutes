// some exercises in dynamic programming

package main

import (
	"fmt"
	"math"
)

/*

1. Fibonacci Series: Implement the Fibonacci series using dynamic programming to efficiently calculate the nth Fibonacci number.

2. Factorial: Calculate the factorial of a given number using dynamic programming.

3. Coin Change Problem: Given a set of coin denominations and a target amount, find the minimum number of coins needed to make up that amount.

4. Longest Common Subsequence (LCS): Given two strings, find the length of the longest common subsequence using dynamic programming.

5. Knapsack Problem: Given a set of items with weights and values, and a maximum weight capacity for a knapsack, determine the maximum value that can be obtained by including items in the knapsack.

6, Maximum Sum Subarray: Given an array of integers, find the contiguous subarray with the largest sum using dynamic programming.

7. Rod Cutting Problem: Given a rod of length n and a list of prices for different rod lengths, find the maximum revenue that can be obtained by cutting and selling the rod.

8. Palindrome Partitioning: Given a string, partition it into substrings such that each substring is a palindrome, and the number of partitions is minimized.

9. Unique Paths: Given a grid with m rows and n columns, find the number of unique paths from the top-left corner to the bottom-right corner. You can only move right or down at any point.

10. Minimum Path Sum: Given a grid with non-negative numbers, find the minimum path sum from the top-left corner to the bottom-right corner. You can only move right or down at any point.

11. Longest Increasing Subsequence (LIS): Given an array of integers, find the length of the longest increasing subsequence using dynamic programming.

12. Maximum Product Subarray: Given an array of integers, find the contiguous subarray with the largest product using dynamic programming.

13. Coin Change 2: Given a set of coin denominations and a target amount, find the number of combinations that make up that amount.

14. Word Break: Given a non-empty string and a dictionary of words, determine if the string can be segmented into a space-separated sequence of dictionary words.

*/

func main() {

	/* Fibonacci and Fatorial functions call */
	fmt.Println(fibonacci(10))
	fmt.Println(factorial(10))

	/* Coin change fnuction call */
	coins := make([]int, 1)
	coins[0] = 2
	// coins[1] = 2
	// coins[2] = 5
	target := 3

	fmt.Println(coinChange(coins, target))

	/*	Longest Common Subsequence(LCS) function call */
	str1 := "abcabcdcd"
	str2 := "abcdddadc"
	LCS(str1, str2)

	/* Knapsack problem */
	w := []int{2, 3, 4, 5}
	v := []int{3, 4, 5, 6}
	kc := 8
	knapsack(w, v, kc)

	/* Maximm sum subarray */
	array := []int{2, 4, 7, 1, 3, 6}
	maxSumSubarray(array)
}

// Fibonacci Series: Implement the Fibonacci series using dynamic programming to efficiently calculate the nth Fibonacci number

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	// Create an array to store the results of subproblems.
	fib := make([]int, n+1)

	// Base cases.
	fib[0] = 0
	fib[1] = 1

	// Compute and store the Fibonacci numbers from 2 to n.
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

// Factorial: Calculate the factorial of a given number using dynamic programming.

func factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	fac := make([]int, n+1)

	fac[0] = 1
	fac[1] = 1

	for i := 2; i <= n; i++ {
		fac[i] = fac[i-1] * i
	}

	return fac[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b

}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a

}

// Coin Change Problem: Given a set of coin denominations and a target amount, find the minimum number of coins needed to make up that amount.

func coinChange(coins []int, target int) int {
	fmt.Println(coins, target)
	dp := make([]int, target+1)

	for i := 1; i <= target; i++ {
		dp[i] = math.MaxInt32 // Initialize each entry with a large value to represent "infinity".
	}

	dp[0] = 0

	for i := 1; i <= target; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
				// fmt.Println(dp[i])
			}
		}
	}

	// If dp[target] remains unchanged (still equal to MaxInt32), it means it's not possible to make up the target amount.
	if dp[target] == math.MaxInt32 {
		return -1
	}

	return dp[target]
}

// 4. Longest Common Subsequence (LCS): Given two strings, find the length of the longest common subsequence using dynamic programming.
func LCS(X, Y string) string {
	m := len(X)
	n := len(Y)

	// create a 2d slice of strings
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// Initialize the first row and first column of the dp array to 0.
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	//lcs_length := dp[m][n]

	lcs := ""
	i := m
	j := n

	for i > 0 && j > 0 {

		if X[i-1] == Y[j-1] {
			lcs += string(X[i-1])
			i--
			j--

		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	return reverseString(lcs)

}

func reverseString(s string) string {
	runes := []rune(s) // Convert the string to a slice of runes

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // Convert the reversed slice back to a string
}

// 5. Knapsack Problem: Given a set of items with weights and values, and a maximum weight capacity for a knapsack,
// determine the maximum value that can be obtained by including items in the knapsack.

func knapsack(weights []int, values []int, knapcap int) int {

	n := len(weights)

	//Create a 2D array to store the maximum values that can be achieved
	// for different knapsack capacities and different items taken.
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, knapcap+1)
	}

	for w := 0; w <= knapcap; w++ {
		dp[0][w] = 0
	}

	// Calculate the maximum value that can be achieved for each subproblem
	for i := 1; i <= n; i++ {
		for w := 0; w <= knapcap; w++ {

			// If the current item's weight exceeds the current knapsack capacity,
			// we cannot take it. So the maximum value remains the same as the previous item.
			if weights[i-1] > w {
				dp[i][w] = dp[i-1][w]
			} else {
				// Otherwise, we have two choices: take the current item or leave it.
				// We choose the maximum of these two options.
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			}

		}
	}

	fmt.Println(dp[n][knapcap])

	// The final result is stored in dp[n][knapsack_capacity]
	return dp[n][knapcap]

}

// 6, Maximum Sum Subarray: Given an array of integers, find the contiguous subarray with the largest sum using dynamic programming.

func maxSumSubarray(arr []int) int {
	n := len(arr)
	max_ending_here := arr[0]
	max_so_far := arr[0]

	for i := 1; i <= n-1; i++ {

		max_ending_here = max(arr[i], max_ending_here+arr[i])
		max_so_far = max(max_so_far, max_ending_here)
	}
	fmt.Println(max_so_far)
	return max_so_far
}
