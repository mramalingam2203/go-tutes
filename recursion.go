// recursion

package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(Fibonacci(7))

	inputArray := []int{1, 2, 3, 4, 5}
	fmt.Println(recursiveSum(inputArray, 0))

}

func factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func Fibonacci(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

func recursiveSum(array []int, startIndex int) int {

	// Base case: If the array is empty, return 0
	if startIndex >= len(array) {
		return 0
	}

	currentElement := array[startIndex]
	restOfArraySum := recursiveSum(array, startIndex+1)
	return currentElement + restOfArraySum
}
