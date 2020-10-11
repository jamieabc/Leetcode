package main

// There is an infrastructure of n cities with some number of roads connecting these cities. Each roads[i] = [ai, bi] indicates that there is a bidirectional road between cities ai and bi.
//
// The network rank of two different cities is defined as the total number of directly connected roads to either city. If a road is directly connected to both cities, it is only counted once.
//
// The maximal network rank of the infrastructure is the maximum network rank of all pairs of different cities.
//
// Given the integer n and the array roads, return the maximal network rank of the entire infrastructure.
//
//
//
// Example 1:
//
// Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]
// Output: 4
// Explanation: The network rank of cities 0 and 1 is 4 as there are 4 roads that are connected to either 0 or 1. The road between 0 and 1 is only counted once.
//
// Example 2:
//
// Input: n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]
// Output: 5
// Explanation: There are 5 roads that are connected to cities 1 or 2.
//
// Example 3:
//
// Input: n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]
// Output: 5
// Explanation: The network rank of 2 and 5 is 5. Notice that all the cities do not have to be connected.
//
//
//
// Constraints:
//
//     2 <= n <= 100
//     0 <= roads.length <= n * (n - 1) / 2
//     roads[i].length == 2
//     0 <= ai, bi <= n-1
//     ai != bi
//     Each pair of cities has at most one road connecting them.

func maximalNetworkRank(n int, roads [][]int) int {
	graph := make([][]bool, n)
	degrees := make([]int, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	for _, road := range roads {
		graph[road[0]][road[1]] = true
		graph[road[1]][road[0]] = true
		degrees[road[0]]++
		degrees[road[1]]++
	}

	var maxConnection int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] {
				maxConnection = max(maxConnection, degrees[i]+degrees[j]-1)
			} else {
				maxConnection = max(maxConnection, degrees[i]+degrees[j])
			}
		}
	}

	return maxConnection
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	each pair of city has at most 1 road connected, total roads =
//		roads(city 1) + roads(city 2) - roads(city 1-2)

//	2.	roads connects to a city is fixed, in order to reduce computation, better
//		not to count it again

//	3.	care about # of roads, and for a given city, is it connect to me
