package main

import (
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

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
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
//		- random pick one vertex
//		- select next vertex with minimum cost edge
//		- record new vertex
//		- pick next vertex w/ minimum cost among all reachable vertexes
//		  (including all former visited vertex, they are all reachable)
