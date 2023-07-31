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

import "fmt"

func main() {
	array := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	activitySelection(array)
}

// Activity Selection Problem: Given a set of activities with start and finish times, select the maximum number of non-overlapping activities.
func activitySelection(activities [][]int) {

	fmt.Println(activities)
}

func sortByIndex(arr [][]int) {
	sorted := make([][]int, 0, len(arr))

}
