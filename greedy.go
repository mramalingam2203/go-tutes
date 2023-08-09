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
	/*
		array := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		activitySelection(array)
	*/

	w := []int{2, 3, 4, 5}
	v := []int{3, 4, 5, 6}
	kc := 8
	knapsack(w, v, kc)

}

func knapsack(weights []int, values []int, knapcap int) int {

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
		return table[i][2] < table[j][2]
	})

	fmt.Println(table)

	return 0

}
