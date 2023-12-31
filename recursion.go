// recursion

package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(Fibonacci(7))

	inputArray := []int{1, 2, 3, 4, 5}
	fmt.Println(recursiveSum(inputArray, 0))

	fmt.Println(recursivePower(2, 5))

	fmt.Println(isPalindrome("abcddcba", 0, 7))

}

/*
Factorial Calculation:
Write a recursive function to calculate the factorial of a given positive integer n.
*/

func factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

/*
Fibonacci Series:
Implement a recursive function to generate the Fibonacci sequence up to the nth term.
*/

func Fibonacci(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

/*
Sum of Array Elements:
Write a recursive function to calculate the sum of all elements in an array.
*/

func recursiveSum(array []int, startIndex int) int {

	// Base case: If the array is empty, return 0
	if startIndex >= len(array) {
		return 0
	}

	currentElement := array[startIndex]
	restOfArraySum := recursiveSum(array, startIndex+1)
	return currentElement + restOfArraySum
}

/*

Power Calculation:
Implement a recursive function to calculate the result of raising a number x to the power of a non-negative integer n.

*/
func recursivePower(x, n int) int {

	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	return x * recursivePower(x, n-1)
}

/*
Palindrome Check:
Create a recursive function to check if a given string is a palindrome (reads the same forwards and backwards).
*/

func isPalindrome(str string, leftIndex, rightIndex int) bool {

	//	fmt.Println(leftIndex, rightIndex, string(str[leftIndex]), string(str[rightIndex]))

	// Base case: If the indices cross each other, it's a palindrome
	if leftIndex >= rightIndex {
		return true
	}

	// Base case: If characters at the indices are different, it's not a palindrome
	if str[leftIndex] != str[rightIndex] {
		return false
	}

	// Recursive case: Check the next pair of characters
	return isPalindrome(str, leftIndex+1, rightIndex-1)

}
