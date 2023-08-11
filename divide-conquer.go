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
	"sort"
)

type Point struct {
	X, Y int
}

func main() {
	// coords := [][]int{{2, 1}, {1, 0}, {1, 2}, {4, 5}, {3, 2}, {5, 4}}
	coords := []Point{{2, 1}, {1, 0}, {1, 2}, {4, 5}, {3, 2}, {5, 4}}
	findConvexHull(coords)

}

func binarySearch() {}
func mergeSort()    {}

func closestPoints()      {}
func maximumSubarraySum() {}

func karatsubaMultiplication(a int, b int) int {
	if a < 10 && b < 10 {
		return a * b
	}

}

func FFT()                 {}
func skyline()             {}
func medianOfSortedArray() {}
func LCS()                 {}

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

/*
func findConvexHull(points [][]int) [][]int {
	n := len(points)

	if n <= 3 {
		return points
	}

	left_half := points[:n/2]
	right_half := points[n/2:]

	// Sort the 2D slice based on x-coord
	sort.Slice(left_half, func(i, j int) bool {
		return left_half[i][0] < left_half[j][0]
	})

	// Sort the 2D slice based on the x-coord
	sort.Slice(right_half, func(i, j int) bool {
		return right_half[i][0] < right_half[j][0]
	})

	fmt.Println(left_half)
	fmt.Println(right_half)

	left_hull := findConvexHull(left_half)   // Recurse on left subset
	right_hull := findConvexHull(right_half) // Recurse on right subset

	fmt.Println("called")

	_ = left_hull
	_ = right_hull

	return points
}

func merge_hulls(left_hull, right_hull) [][]int {
	lower_tangent := find_lower_tangent(left_hull, right_hull)
	upper_tangent := find_upper_tangent(left_hull, right_hull)

	merged_hull := lower_tangent + upper_tangent

	return merged_hull

}

func isCounterClockwise(a, b, c Point) bool {
	return (b.X-a.X)*(c.Y-a.Y)-(b.Y-a.Y)*(c.X-a.X) > 0
}

func find_lower_tangent(l_hull [][]int, r_hull [][]int) [][]int {
	// Sort the 2D slice based on x-coord
	sort.Slice(l_hull, func(i, j int) bool {
		return l_hull[i][0] < l_hull[j][0]
	})

	// Sort the 2D slice based on the x-coord
	sort.Slice(r_hull, func(i, j int) bool {
		return r_hull[i][0] < r_hull[j][0]
	})

}

func find_upper_tangent()
*/
