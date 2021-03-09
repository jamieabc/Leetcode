package main

import (
	"container/heap"
	"math"
	"sort"
)

// There is an undirected weighted connected graph. You are given a positive integer n which denotes that the graph has n nodes labeled from 1 to n, and an array edges where each edges[i] = [ui, vi, weighti] denotes that there is an edge between nodes ui and vi with weight equal to weighti.
//
// A path from node start to node end is a sequence of nodes [z0, z1, z2, ..., zk] such that z0 = start and zk = end and there is an edge between zi and zi+1 where 0 <= i <= k-1.
//
// The distance of a path is the sum of the weights on the edges of the path. Let distanceToLastNode(x) denote the shortest distance of a path between node n and node x. A restricted path is a path that also satisfies that distanceToLastNode(zi) > distanceToLastNode(zi+1) where 0 <= i <= k-1.
//
// Return the number of restricted paths from node 1 to node n. Since that number may be too large, return it modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: n = 5, edges = [[1,2,3],[1,3,3],[2,3,1],[1,4,2],[5,2,2],[3,5,1],[5,4,10]]
// Output: 3
// Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The three restricted paths are:
// 1) 1 --> 2 --> 5
// 2) 1 --> 2 --> 3 --> 5
// 3) 1 --> 3 --> 5
//
// Example 2:
//
// Input: n = 7, edges = [[1,3,1],[4,1,2],[7,3,4],[2,5,3],[5,6,1],[6,7,2],[7,5,3],[2,6,4]]
// Output: 1
// Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The only restricted path is 1 --> 3 --> 7.
//
//
//
// Constraints:
//
// 1 <= n <= 2 * 104
// n - 1 <= edges.length <= 4 * 104
// edges[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 1 <= weighti <= 105
// There is at most one edge between any two nodes.
// There is at least one path between any two nodes.

type MinHeap [][]int // dist, node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int        { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countRestrictedPaths(n int, edges [][]int) int {
	// dist[i]: min cost from n to i
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[n] = 0

	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}

	for _, e := range edges {
		graph[e[0]][e[1]] = e[2]
		graph[e[1]][e[0]] = e[2]
	}

	mh := &MinHeap{}
	heap.Init(mh)
	heap.Push(mh, []int{0, n})

	for mh.Len() > 0 {
		pop := heap.Pop(mh).([]int)

		for to, cost := range graph[pop[1]] {
			if cost != 0 {
				if tmp := dist[pop[1]] + cost; dist[to] > tmp {
					dist[to] = tmp
					heap.Push(mh, []int{tmp, to})
				}
			}
		}
	}

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	return dfs(graph, memo, dist, 1, n)
}

func dfs(graph, memo [][]int, dist []int, cur, target int) int {
	mod := int(1e9 + 7)
	if cur == target {
		return 1
	}

	if memo[cur][target] == 0 {
		var paths int
		for to, weight := range graph[cur] {
			if weight > 0 && dist[cur] > dist[to] {
				paths = (paths + dfs(graph, memo, dist, to, target)) % mod
			}
		}

		memo[cur][target] = paths
	}

	return memo[cur][target]
}

//	Notes
//	1.	not able to finish during contest, TLE

//	2.	from alex 39, final part can be solved by dp, cause distance from start
//		is always decreasing, and if a -> b, then # to b += # to a, very clever

//		topological sort, some kind of

//	3.	inspired from https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/discuss/1097204/PythonJava-Dijkstra-and-Cached-DFS-Clean-and-Concise

//		dijkstra takes O(E * log(V)), E: edges, V: vertex
//		dfs, every edge visited once, takes O(E)

//	4.	inspired from https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/discuss/1097478/C%2B%2B-Dijkstra-%2B-DP

//		restricted path can be built when running dijkstra, since the condition
//		is dijkstra
