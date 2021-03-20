package main

//  In this problem, a tree is an undirected graph that is connected and has no cycles.
//
// The given input is a graph that started as a tree with N nodes (with distinct values 1, 2, ..., N), with one additional edge added. The added edge has two different vertices chosen from 1 to N, and was not an edge that already existed.
//
// The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v] with u < v, that represents an undirected edge connecting nodes u and v.
//
// Return an edge that can be removed so that the resulting graph is a tree of N nodes. If there are multiple answers, return the answer that occurs last in the given 2D-array. The answer edge [u, v] should be in the same format, with u < v.
//
// Example 1:
//
// Input: [[1,2], [1,3], [2,3]]
// Output: [2,3]
// Explanation: The given undirected graph will be like this:
//   1
//  / \
// 2 - 3
//
// Example 2:
//
// Input: [[1,2], [2,3], [3,4], [1,4], [1,5]]
// Output: [1,4]
// Explanation: The given undirected graph will be like this:
// 5 - 1 - 2
//     |   |
//     4 - 3
//
// Note:
// The size of the input 2D-array will be between 3 and 1000.
// Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.
//
//
// Update (2017-09-26):
// We have overhauled the problem description + test cases and specified clearly the graph is an undirected graph. For the directed graph follow up please see Redundant Connection II). We apologize for any inconvenience caused.

func findRedundantConnection(edges [][]int) []int {
	size := len(edges)
	group := make([]int, size+1)
	rank := make([]int, size+1)

	for i := range group {
		group[i] = i
		rank[i] = 1
	}

	var ans []int
	for _, e := range edges {
		p1, p2 := find(group, e[0]), find(group, e[1])

		if p1 == p2 {
			ans = e
			continue
		}

		if rank[p1] >= rank[p2] {
			rank[p1] += rank[p2]
			group[p2] = p1
		} else {
			rank[p2] += rank[p1]
			group[p1] = p2
		}
	}

	return ans
}

func find(group []int, idx int) int {
	if group[idx] != idx {
		group[idx] = find(group, group[idx])
	}
	return group[idx]
}

// tc: O(n^2)
func findRedundantConnection1(edges [][]int) []int {
	size := len(edges)
	list := make([][]int, size+1)
	for i := range list {
		list[i] = make([]int, 0)
	}

	var ans []int

	for _, e := range edges {
		stack := make([]int, len(list[e[0]]))
		copy(stack, list[e[0]])
		visited := make([]bool, size+1)
		visited[e[0]] = true
		var found bool

		for len(stack) > 0 {
			to := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if to == e[1] {
				ans = e
				found = true
				break
			}

			if visited[to] {
				continue
			}
			visited[to] = true

			stack = append(stack, list[to]...)
		}

		if !found {
			list[e[0]] = append(list[e[0]], e[1])
			list[e[1]] = append(list[e[1]], e[0])
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/redundant-connection/discuss/277026/DFS-Java-Solution-With-Explanation
//
//		use dfs to check additional edge [u, v], start from u can reach v means this is
//		duplicate one

//		build adj list: reachable nodes, each time chkeck takes O(n), so total tc will
//		be O(n^2)

//	2.	becareful about slice, when assign and append, then original might be changed

//		e.g. list[e[0]] = [1, 2]
//			 stack := list[e[0]]
//			 stack = append(stack, 3)
//			 list[e[0]] = [1, 2, 3]
