package main

import (
	"container/heap"
	"math"
	"sort"
)

//You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].
//
//The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value of val.
//
//Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.
//
//
//
//Example 1:
//
//Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
//Output: 20
//Explanation:
//
//We can connect the points as shown above to get the minimum cost of 20.
//Notice that there is a unique path between every pair of points.
//
//Example 2:
//
//Input: points = [[3,12],[-2,5],[-4,1]]
//Output: 18
//
//Example 3:
//
//Input: points = [[0,0],[1,1],[1,0],[-1,1]]
//Output: 4
//
//Example 4:
//
//Input: points = [[-1000000,-1000000],[1000000,1000000]]
//Output: 4000000
//
//Example 5:
//
//Input: points = [[0,0]]
//Output: 0
//
//
//
//Constraints:
//
//    1 <= points.length <= 1000
//    -106 <= xi, yi <= 106
//    All pairs (xi, yi) are distinct.

// tc: O(n^2)
func minCostConnectPoints(points [][]int) int {
	visited := make(map[int]bool)

	weights := make([]int, len(points))
	for i := 1; i < len(points); i++ {
		weights[i] = math.MaxInt32
	}

	var cur, total, curCost int

	for count := 0; count < len(points)-1; count++ {
		visited[cur] = true
		curCost = math.MaxInt32
		// update weights for newly visited vertex
		for i := range points {
			if visited[i] {
				continue
			}

			weights[i] = min(weights[i], cost(points[cur], points[i]))

			if weights[i] < curCost {
				curCost = weights[i]
			}
		}

		total += curCost

		// pick next vertex with minimum distance
		curCost = math.MaxInt32
		for i := range weights {
			if visited[i] {
				continue
			}

			if weights[i] < curCost {
				curCost = weights[i]
				cur = i
			}
		}
	}

	return total
}

func dist(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func cost(p1, p2 []int) int {
	return dist(p1[0], p2[0]) + dist(p1[1], p2[1])
}

type MinHeap [][]int // distance, point

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

// tc: O(n^2)
func minCostConnectPoints3(points [][]int) int {
	size := len(points)
	visited := make([]bool, size)

	var cost int
	minHeap := &MinHeap{}

	// start from points[0], find all reachable point weights
	for i := 1; i < size; i++ {
		heap.Push(minHeap, []int{dist(points[0], points[i]), i})
	}
	visited[0] = true
	var popped []int

	for remain := size - 1; remain > 0; remain-- {
		for minHeap.Len() > 0 {
			popped = heap.Pop(minHeap).([]int)
			if !visited[popped[1]] {
				break
			}
		}
		visited[popped[1]] = true
		cost += popped[0]

		// add neighbors reachable from that point
		for i := range points {
			if visited[i] {
				continue
			}

			heap.Push(minHeap, []int{dist(points[popped[1]], points[i]), i})
		}
	}

	return cost
}

// still slow, 832ms
func minCostConnectPoints2(points [][]int) int {
	size := len(points)
	edges := make([][]int, 0)

	for i := range points {
		for j := i + 1; j < size; j++ {
			edges = append(edges, []int{dist(points[i], points[j]), i, j})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	var cost int
	group := make([]int, size)
	rank := make([]int, size)
	for i := range group {
		group[i] = i
		rank[i] = 1
	}

	for _, edge := range edges {
		p1, p2 := find(group, edge[1]), find(group, edge[2])

		// already belongs to same group (connected)
		if p1 == p2 {
			continue
		}

		if rank[p1] >= rank[p2] {
			group[p2] = p1
			rank[p1] += rank[p2]
		} else {
			group[p1] = p2
			rank[p2] += rank[p1]
		}
		cost += edge[0]
	}

	return cost
}

type MinHeap [][]int

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

// really slow 1560ms
func minCostConnectPoints1(points [][]int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	size := len(points)

	for i := range points {
		for j := i + 1; j < size; j++ {
			heap.Push(minHeap, []int{dist(points[i], points[j]), i, j})
		}
	}

	var cost int
	group := make([]int, size)
	for i := range group {
		group[i] = i
	}

	for minHeap.Len() > 0 {
		p := heap.Pop(minHeap).([]int)

		// already visited
		p1, p2 := find(group, p[1]), find(group, p[2])

		// already connected
		if p1 == p2 {
			continue
		}
		group[p2] = p1
		cost += p[0]
	}

	return cost
}

func dist(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func find(group []int, idx int) int {
	if group[idx] != idx {
		group[idx] = find(group, group[idx])
	}

	return group[idx]
}

//	Notes
//	1.	at first I try to use greedy algorithm, but after some failures, I
//		didn't think of a good greedy algorithm to guarantee optimal solution,
//		simply because I didn't consider later intervals

//	2.	from discussion, seems like an algorithm to find minimum spanning tree
//		a video tutorial: https://www.youtube.com/watch?v=5xosHRdxqHA &
//		https://www.youtube.com/watch?v=qOv8K-AJ7o0

//		the problems provides some vertexes & edges, minimum spanning tree is
//		to find a path that connects all points and with minimum cost

//		kruskal algorithm:
//		- separate every vertexes into disjoint sets
//		- sort edges by it's cost (manhattan distance)
//		- start from minimum cost edge, check if two points are not connected
//		  (belongs to different set group)
//		- merges two groups contains two points into one group, record edge

//		to faster determine two points with same group, use union-find

//	3.	another greedy algorithm from https://www.youtube.com/watch?v=K_1urzWrzLs

//		prim's algorithm:
//		- random pick one vertex a
//		- put weight of a to other unvisited vertex into priority queue
//		- select next unvisited vertex with minimum cost edge
//		- record new vertex
//		- pick next vertex w/ minimum cost among all reachable vertexes
//		  (including all former visited vertex, they are all reachable)

//	4.	inspired from https://leetcode.com/problems/min-cost-to-connect-all-points/discuss/843940/C%2B%2B-MST%3A-Kruskal-%2B-Prim's-%2B-Complete-Graph

//		since this problem edges forms a complete graph, can use array to
//		store minimum distance
