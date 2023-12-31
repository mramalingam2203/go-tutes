/*

+ Activity Selection Problem: Given a set of activities with start and finish times, select the maximum number of non-overlapping activities.

+ Fractional Knapsack Problem: Given items with weights and values, fill a knapsack with a maximum weight capacity while maximizing the total value. Items can be taken partially.

+ Huffman Coding: Given a set of characters and their frequencies, construct a binary tree to minimize the total encoding length.

+ Coin Change Problem: Given a set of coin denominations and a target amount, find the minimum number of coins needed to make the target amount.

+ Interval Scheduling Problem: Given a set of intervals with start and finish times, select the maximum number of non-overlapping intervals.

+ Minimum Spanning Tree (MST): Given a connected, undirected graph with edge weights, find the minimum weight spanning tree that connects all vertices.

+ Task Scheduling Problem: Given tasks with their durations and deadlines, schedule the tasks to meet the deadlines while minimizing the total completion time.

+ Job Sequencing Problem: Given jobs with their deadlines and profits, schedule the jobs to maximize the total profit while meeting the deadlines.

+ Greedy Coloring Algorithm: Given an undirected graph, assign colors to vertices such that no two adjacent vertices have the same color, using the fewest number of colors.

+ Dijkstra's Algorithm: Given a graph with non-negative edge weights, find the shortest path from a source vertex to all other vertices.

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	/* Activity Selection */
	// //array := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	// start := []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2}
	// finish := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	// activitySelection(start, finish)

	/* Fractional Knapsack problem */
	// w := []int{2, 3, 5, 7, 1}
	// v := []int{10, 5, 15, 7, 6}
	// kc := 15
	// fmt.Println(knapsack(w, v, kc))

	/* Coin change problem
	coinDenominations := []int{25, 10, 5, 1}
	targetAmount := 63

	minCoins := coinChangeGreedy(coinDenominations, targetAmount)
	fmt.Printf("Minimum coins needed to make %d cents: %d\n", targetAmount, minCoins) */

	/* Inteval Scheduling Problem */
	intervals := [][]int{{1, 3}, {2, 5}, {4, 6}, {7, 8}, {7, 9}}
	greedyIntervalScheduling(intervals)

	/* Task Scheduling Problem */
	times := [][]int{{5, 2}, {3, 3}, {7, 1}, {2, 4}}

	greedyTaskScheduling(times)

}

func knapsack(weights []int, values []int, capacity int) float64 {

	table := make([][]int, len(weights))

	for i := range table {
		table[i] = make([]int, 3)
	}

	for i := 0; i < len(weights); i++ {
		table[i][0] = weights[i]
		table[i][1] = values[i]
		table[i][2] = values[i] / weights[i]
	}

	// Sort the 2D slice based on the second column (index 1)
	sort.Slice(table, func(i, j int) bool {
		return table[i][2] > table[j][2]
	})
	fmt.Println(table)
	totalValue := 0.0

	for item := 0; item < len(table); item++ {
		if capacity >= table[item][0] {
			totalValue += float64(table[item][1])
			capacity -= table[item][0]
		} else {
			totalValue += (float64(table[item][1]) / float64(table[item][1])) * float64(capacity)
			break
		}
	}

	return totalValue
}

func activitySelection(start_time []int, finish_time []int) {
	timetable := make([][]int, len(start_time))

	for i := range timetable {
		timetable[i] = make([]int, 3)
	}

	for i := 0; i < len(timetable); i++ {
		timetable[i][0] = start_time[i]
		timetable[i][1] = finish_time[i]
		timetable[i][2] = i
	}

	// Sort the 2D slice based on the second column (index 1)
	sort.Slice(timetable, func(i, j int) bool {
		return timetable[i][1] < timetable[j][1]
	})

	fmt.Println(timetable)

	for i := 0; i < len(timetable)-1; i++ {
		if timetable[i][1] < timetable[i+1][0] {
			fmt.Println(timetable[i][2])
		}
	}

}

func coinChangeGreedy(coins []int, target int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Sort in descending order

	coinCount := 0

	for _, coin := range coins {
		for target >= coin {
			target -= coin
			coinCount++
		}
	}

	return coinCount
}

func greedyIntervalScheduling(intervals [][]int) {

	// Sort the 2D slice based on the end times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	selectedIntervals := [][]int{{}}
	selectedIntervals[0] = intervals[0]

	for i := 1; i <= len(intervals); i++ {
		if selectedIntervals[len(selectedIntervals)-1][0] < intervals[i-1][0] {
			selectedIntervals = append(selectedIntervals, intervals[i])
		}
	}

	fmt.Println(selectedIntervals)
}

func greedyTaskScheduling(intervals [][]int) {

	// Sort the 2D slice based on the end times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	currentTime := 0
	totalLateness := 0

	for i := 0; i < len(intervals); i++ {

		completionTime := currentTime + intervals[i][0]
		lateness := completionTime - intervals[i][1]
		totalLateness += lateness
		currentTime = completionTime
	}

	fmt.Println(totalLateness)

}
