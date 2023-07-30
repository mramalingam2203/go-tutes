// some exercises in dynamic programming

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(factorial(10))

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
