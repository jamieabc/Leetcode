package main

import "fmt"

// Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to find the number of connected components in an undirected graph.
//
// Example 1:
//
// Input: n = 5 and edges = [[0, 1], [1, 2], [3, 4]]
//
//      0          3
//      |          |
//      1 --- 2    4
//
// Output: 2
//
// Example 2:
//
// Input: n = 5 and edges = [[0, 1], [1, 2], [2, 3], [3, 4]]
//
//      0           4
//      |           |
//      1 --- 2 --- 3
//
// Output:  1
//
// Note:
// You can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.

func countComponents(n int, edges [][]int) int {
	parents, ranks := make([]int, n), make([]int, n)
	for i := range parents {
		parents[i] = i
		ranks[i] = 1
	}

	count := n

	for _, edge := range edges {
		p1, p2 := find(parents, edge[0]), find(parents, edge[1])

		if p1 != p2 {
			count--

			if ranks[p1] >= ranks[p2] {
				parents[p2] = p1
				ranks[p1]++
			} else {
				parents[p1] = p2
				ranks[p2]++
			}
		}
	}

	return count
}

func find(parents []int, idx int) int {
	if parents[idx] != idx {
		parents[idx] = find(parents, parents[idx])
	}

	return parents[idx]
}

func countComponents3(n int, edges [][]int) int {
	visited := make([]bool, n)

	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	var ans int

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		ans++

		bfs(visited, graph, i)
	}

	return ans
}

func bfs(visited []bool, graph [][]int, point int) {
	queue := []int{point}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if visited[p] {
			continue
		}
		visited[p] = true

		for _, to := range graph[p] {
			if !visited[to] {
				queue = append(queue, to)
			}
		}
	}
}

func countComponents2(n int, edges [][]int) int {
	if n == 0 {
		return 0
	}

	arr := make([]int, n)
	for i := range arr {
		arr[i] = -1
	}
	stack := make([]int, 0)

	for _, e := range edges {
		small, large := smallLarge(e[0], e[1])

		if arr[large] == -1 && arr[small] == -1 {
			// no data, set one
			arr[small] = small
			arr[large] = small
		} else if arr[large] != -1 && arr[small] != -1 {
			target := min(parent(arr, large), parent(arr, small))
			if arr[large] != target {
				stack = append(stack, arr[large])
			}
			if arr[small] != target {
				stack = append(stack, arr[small])
			}

			for len(stack) != 0 {
				s := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if arr[s] != target {
					stack = append(stack, arr[s])
					arr[s] = target
				}
			}
		} else if arr[large] != -1 {
			arr[small] = arr[large]
		} else {
			arr[large] = arr[small]
		}
	}

	var count int
	for i := range arr {
		if arr[i] == i || arr[i] == -1 {
			count++
		}
	}

	return count
}

func parent(arr []int, n int) int {
	if arr[n] != n {
		arr[n] = parent(arr, arr[n])
	}
	return arr[n]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func countComponents1(n int, edges [][]int) int {
	mapping := make(map[int][]int)

	for _, e := range edges {
		mapping[e[0]] = append(mapping[e[0]], e[1])
		mapping[e[1]] = append(mapping[e[1]], e[0])
	}

	var count int
	stack := make([]int, 0)
	arr := make([]bool, n)

	for i := range arr {
		if !arr[i] {
			stack := append(stack, i)

			for len(stack) != 0 {
				s := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !arr[s] {
					arr[s] = true
					stack = append(stack, mapping[s]...)
				}
			}
			count++
		}
	}

	return count
}

//	Notes
//	1.	traversing route should not be order dependent

//	2.	from test case, even if node is not connected, still treat it as
//		one component

//	3.	the algo is still order dependent

//	4.	I think the algo is really bad, need to rethink from beginning. route
//		could be circular, so need a way to detect circular (already
//		traveled place). The point is that parsing data implies algo should
//		be order independent, cause it might exist 0-1-2  3-4-5 0-4, first
//		and second interval is connected when 0-4 is encountered.

//	5.	inspired from https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/discuss/77574/Easiest-2ms-Java-Solution

//		the problem is actually connect number together, so it's quite
//		straight forward to first assume all of them are disjoint, and
//		for each edge, find it's first & second's root and connect these
//		two roots together, reduce count by 1

//	6. 	when updating root1 & root2, it should update root1 & root2's value
//		but not child (e0, e1)

//	7.	inspired from https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/discuss/77578/Java-concise-DFS

//		when traversing, it's important to mark all connect points, e.g.
//		[1, 5] means 1 -> 5 & 5 -> 1, with this data, I can avoid order
//		traverse dependent problem.

//		time complexity is O(V+E), every node and edges are visited
//		space complexity O(V+E)

//	8.	reference https://www.cs.princeton.edu/~rs/AlgsDS07/01UnionFind.pdf

//		explains find & union

//	9.	inspired from solution, can use bfs/dfs to traverse, for every vertex not visited,
//		mark all connected nodes as visited
