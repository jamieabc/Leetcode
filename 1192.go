package main

// There are n servers numbered from 0 to n-1 connected by undirected server-to-server connections forming a network where connections[i] = [a, b] represents a connection between servers a and b. Any server can reach any other server directly or indirectly through the network.
//
// A critical connection is a connection that, if removed, will make some server unable to reach some other server.
//
// Return all critical connections in the network in any order.
//
//
//
// Example 1:
//
// Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
// Output: [[1,3]]
// Explanation: [[3,1]] is also accepted.
//
//
//
// Constraints:
//
// 1 <= n <= 10^5
// n-1 <= connections.length <= 10^5
// connections[i][0] != connections[i][1]
// There are no repeated connections.

func criticalConnections(n int, connections [][]int) [][]int {
	ranks := make([]int, n)
	for i := range ranks {
		ranks[i] = -1
	}

	adjList := make([][]int, n)

	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, connection := range connections {
		adjList[connection[0]] = append(adjList[connection[0]], connection[1])
		adjList[connection[1]] = append(adjList[connection[1]], connection[0])
	}

	removable := make(map[[2]int]bool)

	var rank int
	dfs(adjList, ranks, removable, -1, 0, &rank)

	critical := make([][]int, 0)
	for _, connection := range connections {
		if _, ok := removable[[2]int{connection[0], connection[1]}]; !ok {
			critical = append(critical, connection)
		}
	}

	return critical
}

func dfs(adjList [][]int, ranks []int, removable map[[2]int]bool, parent, cur int, rank *int) int {
	if ranks[cur] != -1 {
		return ranks[cur]
	}
	ranks[cur] = *rank
	*rank++
	minRank := ranks[cur]

	for _, to := range adjList[cur] {
		// no need to go back to parent
		if to == parent {
			continue
		}

		nextRank := dfs(adjList, ranks, removable, cur, to, rank)

		if nextRank <= ranks[cur] {
			removable[[2]int{cur, to}] = true
			removable[[2]int{to, cur}] = true
		}

		minRank = min(minRank, nextRank)
	}

	return minRank
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, an bridge edge (remove this edges causes whoel graph
//		separated) is the edges not in loop

//		using dfs, if a node is visited second time, this node is in a loop

//		to find edges in loop, need another value to track (ranks for time), this value
//		is increasing, so if any adjacent nodes that has smaller value, it means a loop
//		is found

//		return minimum ranks in reachable nodes, parent can determine which edge is
//		in loop (able to be removed)

//		very brilliant solution, it needs fully understanding of dfs

//	2.	inspired form https://leetcode.com/problems/critical-connections-in-a-network/discuss/382638/DFS-detailed-explanation-O(orEor)-solution

//		author has a pretty good explanation of the problem
