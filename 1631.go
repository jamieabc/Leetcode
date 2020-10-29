package main

import (
	"math"
	"sort"
)

// You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.
//
// A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.
//
// Return the minimum effort required to travel from the top-left cell to the bottom-right cell.
//
//
//
// Example 1:
//
// Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
// Output: 2
// Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
// This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
//
// Example 2:
//
// Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
// Output: 1
// Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].
//
// Example 3:
//
// Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
// Output: 0
// Explanation: This route does not require any effort.
//
//
//
// Constraints:
//
//     rows == heights.length
//     columns == heights[i].length
//     1 <= rows, columns <= 100
//     1 <= heights[i][j] <= 106

// bellman-ford, relaxation by edges
// tc: O((mn)^2)
func minimumEffortPath(heights [][]int) int {
	w, h := len(heights[0]), len(heights)

	// dp[i][j] means shortest path from (0, 0) to (i, j)
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	queue := [][]int{{0, 0, 0}}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		for _, d := range dir {
			x, y := q[0]+d[0], q[1]+d[1]

			if validPoint(heights, x, y) {
				if cost := max(q[2], abs(heights[y][x]-heights[q[1]][q[0]])); dp[y][x] > cost {
					dp[y][x] = cost
					queue = append(queue, []int{x, y, cost})
				}
			}
		}
	}

	return dp[h-1][w-1]
}

func minimumEffortPath4(heights [][]int) int {
	w, h := len(heights[0]), len(heights)
	var low, high, ans int
	for i := range heights {
		for j := range heights[0] {
			high = max(high, heights[i][j])
		}
	}

	for low <= high {
		mid := low + (high-low)/2
		visited := make([][]bool, h)
		for i := range visited {
			visited[i] = make([]bool, w)
		}

		if reachableDFS(visited, heights, mid, 0, 0) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

// DFS
func reachableDFS(visited [][]bool, heights [][]int, criteria, x, y int) bool {
	w, h := len(heights[0]), len(heights)
	if x == w-1 && y == h-1 {
		return true
	}

	if visited[y][x] {
		return false
	}
	visited[y][x] = true

	var reached bool

	for _, d := range dir {
		newX, newY := x+d[0], y+d[1]
		if validPoint(heights, newX,
			newY) && abs(heights[y][x]-heights[newY][newX]) <= criteria {
			reached = reached || reachableDFS(visited, heights, criteria, newX, newY)

			if reached {
				return true
			}
		}
	}

	return false
}

// tc: O(log(10^6) + mn)
// binary search from 0 - 10^6
// BFS tc: O(v+e), vertex: m*n, edges: m*n
func minimumEffortPath3(heights [][]int) int {
	var low, high, ans int
	for i := range heights {
		for j := range heights[0] {
			high = max(high, heights[i][j])
		}
	}

	for low <= high {
		mid := low + (high-low)/2

		if reachable(heights, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

// BFS
func reachable(heights [][]int, criteria int) bool {
	w, h := len(heights[0]), len(heights)
	queue := [][]int{{0, 0}}

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if q[0] == w-1 && q[1] == h-1 {
			return true
		}

		if visited[q[1]][q[0]] {
			continue
		}
		visited[q[1]][q[0]] = true

		for _, d := range dir {
			x, y := q[0]+d[0], q[1]+d[1]

			if validPoint(heights, x, y) && !visited[y][x] && abs(heights[y][x]-heights[q[1]][q[0]]) <= criteria {
				queue = append(queue, []int{
					x, y,
				})
			}
		}
	}

	return false
}

// tc: O(mn log(mn))
func minimumEffortPath2(heights [][]int) int {
	w, h := len(heights[0]), len(heights)
	target := w*h - 1
	edges := buildGraph(heights)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	parents := make([]int, w*h)
	for i := range parents {
		parents[i] = i
	}

	rank := make([]int, w*h) // merge shorter path to longer path

	for _, edge := range edges {
		p1, p2 := find(parents, edge[0]), find(parents, edge[1])

		if p1 != p2 {
			if rank[p1] >= rank[p2] {
				parents[p2] = p1
				rank[p1]++
			} else if rank[p1] < rank[p2] {
				parents[p1] = p2
				rank[p2]++
			}
		}

		if find(parents, 0) == find(parents, target) {
			return edge[2]
		}
	}

	return 0
}

func find(parents []int, idx int) int {
	if parents[idx] != idx {
		parents[idx] = find(parents, parents[idx])
	}

	return parents[idx]
}

func buildGraph(heights [][]int) [][]int {
	w := len(heights[0])
	edges := make([][]int, 0)

	for i := range heights {
		for j := range heights[0] {
			idx := i*w + j

			for _, d := range dir {
				x, y := j+d[0], i+d[1]
				nextID := y*w + x

				if validPoint(heights, x, y) {
					edges = append(edges, []int{
						idx,
						nextID,
						abs(heights[y][x] - heights[i][j]),
					})
				}
			}
		}
	}

	return edges
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
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

// tc: O(wh log(wh)), w & h: width & height of points
func minimumEffortPath1(heights [][]int) int {
	w, h := len(heights[0])-1, len(heights)-1
	mh := &MinHeap{}
	heap.Init(mh)

	visited := make([][]bool, len(heights))
	for i := range visited {
		visited[i] = make([]bool, len(heights[0]))
	}
	heap.Push(mh, []int{0, 0, 0})

	for mh.Len() > 0 {
		top := heap.Pop(mh).([]int)
		if visited[top[0]][top[1]] {
			continue
		}
		visited[top[0]][top[1]] = true

		if top[0] == h && top[1] == w {
			return top[2]
		}

		for _, d := range dir {
			newY, newX := top[0]+d[0], top[1]+d[1]

			if validPoint(heights, newX, newY) && !visited[newY][newX] {
				heap.Push(mh, []int{
					newY, newX, max(top[2],
						abs(heights[newY][newX]-heights[top[0]][top[1]])),
				})
			}
		}
	}

	return 0
}

var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func validPoint(heights [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(heights[0]) && y < len(heights)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	for all paths, find minimum effort, tc: O(3^(mn)), for every point, there
//		are 3 ways to go (discard where comes from)

//	2.	every edge >= 0, find min cost from one point to another, Dijkstra
//		can be used, tc: O(mn log(mn)), m: row count, n: column count.
//		it's actually O(v log(v)), v: vertex count

//	3.	union by rank: while doing union, put shorter paths into longer paths to
//		reduce computation

//	4.	dijkstra cannot apply to any edge with negative weight, and bellman-ford
//		cannot apply to loop with negative weight

//	5.	for binary search, iterate through whole array to find reasonable range,
//		reduce unwanted computations
