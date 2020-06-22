package main

import (
	"container/heap"
	"math"
)

// There are N network nodes, labelled 1 to N.
//
// Given times, a list of travel times as directed edges times[i] = (u, v, w), where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.
//
// Now, we send a signal from a certain node K. How long will it take for all nodes to receive the signal? If it is impossible, return -1.
//
//
//
// Example 1:
//
// Input: times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
// Output: 2
//
//
//
// Note:
//
//     N will be in the range [1, 100].
//     K will be in the range [1, N].
//     The length of times will be in the range [1, 6000].
//     All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 0 <= w <= 100.

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make([][]int, N+1)
	for i := range graph {
		graph[i] = make([]int, N+1)
		for j := range graph {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = math.MaxInt32
			}
		}
	}

	for _, time := range times {
		graph[time[0]][time[1]] = time[2]
	}

	queue := []int{K}
	distance := make([]int, N+1)
	for i := range distance {
		if i == K {
			distance[i] = 0
		} else {
			distance[i] = math.MaxInt32
		}
	}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		// find all possible nodes
		for i := 1; i < N+1; i++ {
			if i != K && graph[popped][i] != math.MaxInt32 {
				newDist := distance[popped] + graph[popped][i]

				if distance[i] > newDist {
					distance[i] = newDist
					queue = append(queue, i)
				}
			}
		}
	}

	maxCost := math.MinInt32
	for i := 1; i < N+1; i++ {
		if distance[i] == math.MaxInt32 {
			return -1
		} else {
			maxCost = max(maxCost, distance[i])
		}
	}
	return maxCost
}

type route struct {
	cost int
	node int
}

type PriorityQueue []route

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(route))
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	popped := old[len(old)-1]
	*p = old[:len(old)-1]
	return popped
}

func networkDelayTime1(times [][]int, N int, K int) int {
	// src -> dst with cost
	routes := make(map[int]map[int]int)
	for _, time := range times {
		if _, ok := routes[time[0]]; !ok {
			routes[time[0]] = make(map[int]int)
		}
		routes[time[0]][time[1]] = time[2]
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	visited := make(map[int]int)
	heap.Push(pq, route{
		cost: 0,
		node: K,
	})

	for pq.Len() > 0 {
		popped := heap.Pop(pq).(route)

		if _, ok := visited[popped.node]; ok {
			continue
		} else {
			visited[popped.node] = popped.cost
		}

		for n, cost := range routes[popped.node] {
			heap.Push(pq, route{
				cost: popped.cost + cost,
				node: n,
			})
		}
	}

	totalCost := math.MinInt32
	for i := 1; i <= N; i++ {
		if cost, ok := visited[i]; !ok {
			return -1
		} else {
			totalCost = max(totalCost, cost)
		}
	}
	return totalCost
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	could exist acyclic, use dijkstra

//	2.	inspired from sample code a node is put into queue as long as cost
//		is minimum

//	3.	cost could be 0

//	4.	inspired from https://leetcode.com/problems/network-delay-time/discuss/109982/C%2B%2B-Bellman-Ford

//		some comparison:
//			tc for dijkstra: O(e + v log v)
//			tc for bellman-ford: O(ve)
//			v: vertex, e: edge

//		dijkstra may be wrong when weight is negative

//		the reason I cannot write out a solution in first attempt is because
//		I stuck at the "visited" nodes. a node to traverse again doesn't
//		decided by visited or not, it's by its cost. If cost is smaller,
//		then all nodes afterwards should be updated.
