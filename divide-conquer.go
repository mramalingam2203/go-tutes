/*

1. Binary Search:
Implement binary search to find the index of a target element in a sorted array.

2. Merge Sort:
Write a merge sort algorithm to sort an array or list of elements.

3. Closest Pair of Points:
Given an array of 2D points, find the two points with the smallest Euclidean distance between them.

4. Maximum Subarray Sum:
Find the contiguous subarray within an array that has the largest sum.

5. Matrix Multiplication:
Multiply two matrices using the Strassen's algorithm or other divide and conquer matrix multiplication methods.

6. Counting Inversions:
Given an array, count the number of inversions, where an inversion occurs if A[i] > A[j] and i < j.

7. Karatsuba Multiplication:
Implement integer multiplication using the Karatsuba algorithm, which reduces the number of subproblems.

8. Fast Fourier Transform (FFT):
Implement the FFT algorithm to efficiently compute the discrete Fourier transform of a sequence.

9. Closest Pair in Strips:
Extend the closest pair problem to find the closest pair of points where one point lies in each of two disjoint vertical strips.

10. The Skyline Problem:
Given a list of building rectangles, find the skyline formed by these buildings when viewed from a distance.

11. Median of Two Sorted Arrays:
Find the median of two sorted arrays of equal size using a divide and conquer approach.

12. Longest Common Subsequence (LCS):
Find the longest common subsequence between two strings using dynamic programming and divide and conquer.

13. The Game of Nim:
Implement the game of Nim and develop a strategy for optimal moves using the divide and conquer approach.

14. Counting Triangles:
Given a set of points in a 2D plane, count the number of triangles that can be formed using these points as vertices.

15. Finding Majority Element:
Determine if there is a majority element (appearing more than n/2 times) in an array using divide and conquer.

16. Convex Hull:
Find the convex hull of a set of points in a 2D plane using algorithms like Graham's scan or Jarvis march

*/

package main

import (
	"strconv"
)

type Point struct {
	X, Y int
}

func main() {
	// coords := [][]int{{2, 1}, {1, 0}, {1, 2}, {4, 5}, {3, 2}, {5, 4}}
	/* coords := []Point{{2, 1}, {1, 0}, {1, 2}, {4, 5}, {3, 2}, {5, 4}}
	findConvexHull(coords) */

	num1 := 1010120
	num2 := 231301
	karatsubaMultiplication(num1, num2)

}

func binarySearch() {}
func mergeSort()    {}

func closestPoints()      {}
func maximumSubarraySum() {}

func karatsubaMultiplication(x int, y int) int {
	if x < 10 && y < 10 {
		return x * y
	}

	x_str := strconv.Itoa(x)
	y_str := strconv.Itoa(y)

	n := max(len(x_str), len(y_str))
	m := n / 2

	a := x_str[:m]
	b := x_str[n-m:]

	c := y_str[:m]
	d := y_str[n-m:]

	/*
	   a = first m digits of x
	   b = last n - m digits of x
	   c = first m digits of y
	   d = last n - m digits of y
	*/

	return 0

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func FFT()                 {}
func skyline()             {}
func medianOfSortedArray() {}
func LCS()                 {}

/*
func findConvexHull(points []Point) []Point {
	n := len(points)

	if n <= 3 {
		return points
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	leftHalf := points[:n/2]
	rightHalf := points[n/2:]

	leftHull := findConvexHull(leftHalf)
	rightHull := findConvexHull(rightHalf)

	return mergeHulls(leftHull, rightHull)

}

func mergeHulls(leftHull, rightHull []Point) []Point {
	lowerTangent := findLowerTangent(leftHull, rightHull)
	upperTangent := findUpperTangent(leftHull, rightHull)

	mergedHull := append(lowerTangent, upperTangent...)
	return mergedHull
}

func findLowerTangent(leftHull, rightHull []Point) []Point {
	leftIdx := findLeftMostPointIndex(rightHull)
	rightIdx := findRightMostPointIndex(leftHull)

	for {
		newLeftIdx := (leftIdx + 1) % len(rightHull)
		newRightIdx := (rightIdx + 1) % len(leftHull)

		if !isCounterClockwise(rightHull[leftIdx], leftHull[rightIdx], rightHull[newLeftIdx]) {
			leftIdx = newLeftIdx
		} else if isCounterClockwise(leftHull[rightIdx], rightHull[leftIdx], leftHull[newRightIdx]) {
			rightIdx = newRightIdx
		} else {
			break
		}
	}

	return []Point{rightHull[leftIdx], leftHull[rightIdx]}
}

func isCounterClockwise(a, b, c Point) bool {
	return (b.X-a.X)*(c.Y-a.Y)-(b.Y-a.Y)*(c.X-a.X) > 0
}
*/
