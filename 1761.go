package main

import "math"

// You are given an undirected graph. You are given an integer n which is the number of nodes in the graph and an array edges, where each edges[i] = [ui, vi] indicates that there is an undirected edge between ui and vi.
//
// A connected trio is a set of three nodes where there is an edge between every pair of them.
//
// The degree of a connected trio is the number of edges where one endpoint is in the trio, and the other is not.
//
// Return the minimum degree of a connected trio in the graph, or -1 if the graph has no connected trios.
//
//
//
// Example 1:
//
// Input: n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
// Output: 3
// Explanation: There is exactly one trio, which is [1,2,3]. The edges that form its degree are bolded in the figure above.
//
// Example 2:
//
// Input: n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
// Output: 0
// Explanation: There are exactly three trios:
// 1) [1,4,3] with degree 0.
// 2) [2,5,6] with degree 2.
// 3) [5,6,7] with degree 2.
//
//
//
// Constraints:
//
//     2 <= n <= 400
//     edges[i].length == 2
//     1 <= edges.length <= n * (n-1) / 2
//     1 <= ui, vi <= n
//     ui != vi
//     There are no repeated edges.

// tc: O(n^3)
func minTrioDegree(n int, edges [][]int) int {
	graph := make([][]bool, n+1)
	for i := range graph {
		graph[i] = make([]bool, n+1)
	}

	degrees := make([]int, n+1)

	for _, e := range edges {
		graph[e[0]][e[1]] = true
		graph[e[1]][e[0]] = true

		degrees[e[0]]++
		degrees[e[1]]++
	}

	minDegree := math.MaxInt32

	for i := 1; i <= n; i++ {
		to := graph[i]

		for j := range to {
			if to[j] {
				for k := range graph[j] {
					if graph[j][k] && k > j && k > i && graph[i][k] {
						minDegree = min(minDegree, degrees[i]+degrees[j]+degrees[k]-6)
					}
				}
			}
		}
	}

	if minDegree == math.MaxInt32 {
		return -1
	}

	return minDegree
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	no need to try already visited nodes

//	2.	inspired from sample code, can be optimized by [][]bool
