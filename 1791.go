package main

// There is an undirected star graph consisting of n nodes labeled from 1 to n. A star graph is a graph where there is one center node and exactly n - 1 edges that connect the center node with every other node.
//
// You are given a 2D integer array edges where each edges[i] = [ui, vi] indicates that there is an edge between the nodes ui and vi. Return the center of the given star graph.
//
//
//
// Example 1:
//
// Input: edges = [[1,2],[2,3],[4,2]]
// Output: 2
// Explanation: As shown in the figure above, node 2 is connected to every other node, so 2 is the center.
//
// Example 2:
//
// Input: edges = [[1,2],[5,1],[1,3],[1,4]]
// Output: 1
//
//
//
// Constraints:
//
//     3 <= n <= 105
//     edges.length == n - 1
//     edges[i].length == 2
//     1 <= ui, vi <= n
//     ui != vi
//     The given edges represent a valid star graph.

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}

	return edges[0][1]
}

func findCenter1(edges [][]int) int {
	size := len(edges)
	inDegree := make([]int, size+2)

	for _, e := range edges {
		inDegree[e[0]]++
		inDegree[e[1]]++

		if inDegree[e[0]] == size {
			return e[0]
		}

		if inDegree[e[1]] == size {
			return e[1]
		}
	}

	return 0
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/find-center-of-star-graph/discuss/1108319/C%2B%2BJava-O(1)-or-O(1)-or-1-liner

//		a center node occurs at ever edge...
