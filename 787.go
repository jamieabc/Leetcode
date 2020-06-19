package main

import (
	"container/heap"
	"math"
)

// There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.
//
// Now given all the cities and flights, together with starting city src and the destination dst, your task is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.
//
// Example 1:
// Input:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 1
// Output: 200
// Explanation:
// The graph looks like this:
//
//
// The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
// Example 2:
// Input:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 0
// Output: 500
// Explanation:
// The graph looks like this:
//
//
// The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
//
//
// Constraints:
//
// The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
// The size of flights will be in range [0, n * (n - 1) / 2].
// The format of each flight will be (src, dst, price).
// The price of each flight will be in the range [1, 10000].
// k is in the range of [0, n - 1].
// There will not be any duplicated flights or self cycles.

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	neighbors := make(map[int]map[int]int)
	for _, flight := range flights {
		if _, ok := neighbors[flight[0]]; !ok {
			neighbors[flight[0]] = make(map[int]int)
		}
		neighbors[flight[0]][flight[1]] = flight[2]
	}
	visited := make(map[int]bool)
	minCost := math.MaxInt32

	dfs(src, dst, K+1, 0, &minCost, neighbors, visited)

	if minCost == math.MaxInt32 {
		return -1
	}
	return minCost
}

func dfs(src, dst, stops, cost int, minCost *int, neighbors map[int]map[int]int, visited map[int]bool) {
	// dst city reached
	if src == dst {
		*minCost = min(*minCost, cost)
		return
	}

	// reach stop limit
	if stops == 0 {
		return
	}

	for toCity, newCost := range neighbors[src] {
		if !visited[toCity] {
			// no need to try cost larger than minCost
			if cost > *minCost || cost+newCost > *minCost {
				continue
			}

			visited[toCity] = true

			dfs(toCity, dst, stops-1, cost+newCost, minCost, neighbors, visited)

			visited[toCity] = false
		}
	}
}

func findCheapestPrice2(n int, flights [][]int, src int, dst int, K int) int {
	// dp[i][j] means minimum cost for vertex src to vertex i using j stops
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, K+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// every flight from src city is ok
	for i := range dp[src] {
		dp[src][i] = 0
	}

	// for every flight, check by stops to decided next minimum cost
	for i := 1; i <= K+1; i++ {
		for _, flight := range flights {
			if dp[flight[0]][i-1] != math.MaxInt32 {
				dp[flight[1]][i] = min(dp[flight[1]][i], dp[flight[0]][i-1]+flight[2])
			}
		}
	}

	if dp[dst][K+1] == math.MaxInt32 {
		return -1
	}
	return dp[dst][K+1]
}

type flight struct {
	cost  int
	city  int
	stops int
}

type shortest []flight

func (s shortest) Len() int {
	return len(s)
}

func (s shortest) Less(i, j int) bool {
	return s[i].cost < s[j].cost
}

func (s shortest) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *shortest) Pop() interface{} {
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *shortest) Push(i interface{}) {
	*s = append(*s, i.(flight))
}

func findCheapestPrice1(n int, flights [][]int, src int, dst int, K int) int {
	h := shortest{}
	heap.Init(&h)

	// routes - store source city to destination city
	routes := make(map[int][]flight)
	for _, oneFlight := range flights {
		routes[oneFlight[0]] = append(routes[oneFlight[0]], flight{
			cost: oneFlight[2],
			city: oneFlight[1],
		})
	}

	heap.Push(&h, flight{
		cost:  0,
		city:  src,
		stops: K + 1,
	})

	for h.Len() > 0 {
		popped := heap.Pop(&h).(flight)

		// shortest path to dst found
		if popped.city == dst {
			return popped.cost
		}

		// only process those cities in K stops
		if popped.stops > 0 {
			neighbors := routes[popped.city]

			for _, neighbor := range neighbors {
				// not to process duplicate city
				visited := make(map[int]bool)
				if _, ok := visited[neighbor.city]; !ok {
					f := flight{
						cost:  popped.cost + neighbor.cost,
						city:  neighbor.city,
						stops: popped.stops - 1,
					}

					heap.Push(&h, f)
					visited[neighbor.city] = true
				}
			}
		}
	}

	return -1
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	when traversing possible flights, should store all possible flight, not
//		only the last one

//	2.	need to revisit city again, because it is from different stop

//	3.	too many flight to a city, running out of memory, use map instead

//	4.	K is a limit criteria, I need to know min cost for a city with # of
//		stops

//	5.	can only push if stops valid

//	6.	avoid circular route

//	7.	inspired from https://leetcode.com/problems/cheapest-flights-within-k-stops/discuss/686774/SUGGESTION-FOR-BEGINNERS-SIMPLE-STEPS-BFS-or-DIJKSHTRA-or-DP-DIAGRAM

//		author has a better naming of neighbor

//		also, when using dp to store minimum cost, K is an important factor,
//		more than K vertex visited should not allowed

//		one important thing is that every flight start from src should be 0
//		cost

//		originally I was using a queue to store next round cities, which is more
//		complex, author uses dp[i][j] as an indicator to know what city should
//		be processed, very clever

//	8.	the reason heap can be used here is because it's a increasing sequence,
//		every stop visited will increase cost.

//	9.	inspired from https://leetcode.com/problems/cheapest-flights-within-k-stops/discuss/128217/Three-C%2B%2B-solutions-BFS-DFS-and-BF

//		add dfs & bfs

//	10.	dfs time limit exceed, need to prune impossible paths

//	11.	inspired from https://leetcode.com/problems/cheapest-flights-within-k-stops/discuss/361711/Java-DFSBFSBellman-Ford-Dijkstra's

//		DFS tc: O(n*k), sc: O(n+k)
//		BFS tc: O(n*k), sc: O(n*k)
//		bellman-ford tc: O(n*k), sc: O(n)
//		Dijkstra tc: O(n*k*(log(n*k)) for each vertex will be put into pq at
//					 most k times, and operation for pq is maximum size of n*k
//				 sc: O(m+n*k)
