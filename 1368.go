package main

import "container/heap"

// Given a m x n grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell. The sign of grid[i][j] can be:
//
// 1 which means go to the cell to the right. (i.e go from grid[i][j] to grid[i][j + 1])
// 2 which means go to the cell to the left. (i.e go from grid[i][j] to grid[i][j - 1])
// 3 which means go to the lower cell. (i.e go from grid[i][j] to grid[i + 1][j])
// 4 which means go to the upper cell. (i.e go from grid[i][j] to grid[i - 1][j])
//
// Notice that there could be some invalid signs on the cells of the grid which points outside the grid.
//
// You will initially start at the upper left cell (0,0). A valid path in the grid is a path which starts from the upper left cell (0,0) and ends at the bottom-right cell (m - 1, n - 1) following the signs on the grid. The valid path doesn't have to be the shortest.
//
// You can modify the sign on a cell with cost = 1. You can modify the sign on a cell one time only.
//
// Return the minimum cost to make the grid have at least one valid path.
//
//
//
// Example 1:
//
// Input: grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
// Output: 3
// Explanation: You will start at point (0, 0).
// The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)
// The total cost = 3.
//
// Example 2:
//
// Input: grid = [[1,1,3],[3,2,2],[1,1,4]]
// Output: 0
// Explanation: You can follow the path from (0, 0) to (2, 2).
//
// Example 3:
//
// Input: grid = [[1,2],[4,3]]
// Output: 1
//
// Example 4:
//
// Input: grid = [[2,2,2],[2,2,2]]
// Output: 3
//
// Example 5:
//
// Input: grid = [[4]]
// Output: 0
//
//
//
// Constraints:
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100

var dirs = [][2]int{
	{0, 0},
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func minCost(grid [][]int) int {
	w, h := len(grid[0]), len(grid)
	cost := make([][]int, h)
	for i := range cost {
		cost[i] = make([]int, w)
		for j := range cost[i] {
			cost[i][j] = -1
		}
	}
	cost[0][0] = 0
	stack := [][2]int{{0, 0}} // y, x
	var dist int

	for cost[h-1][w-1] == -1 {
		sameCost := dfs(grid, cost, stack, dist)

		dist++
		nextCost := bfs(grid, cost, sameCost, dist)

		stack = nextCost
	}

	return cost[h-1][w-1]
}

func dfs(grid, cost [][]int, stack [][2]int, dist int) [][2]int {
	w, h := len(grid[0]), len(grid)
	next := make([][2]int, 0)

	for len(stack) > 0 {
		size := len(stack)

		for i := 0; i < size; i++ {
			p := stack[i]
			next = append(next, p)

			newY, newX := p[0]+dirs[grid[p[0]][p[1]]][0], p[1]+dirs[grid[p[0]][p[1]]][1]

			if newY >= 0 && newX >= 0 && newY < h && newX < w && cost[newY][newX] == -1 {
				next = append(next, [2]int{newY, newX})
				stack = append(stack, [2]int{newY, newX})
				cost[newY][newX] = dist
			}
		}

		stack = stack[size:]
	}

	return next
}

func bfs(grid, cost [][]int, stack [][2]int, dist int) [][2]int {
	w, h := len(grid[0]), len(grid)
	size := len(stack)

	for i := 0; i < size; i++ {
		p := stack[i]

		for j := 1; j < len(dirs); j++ {
			dir := dirs[j]
			newY, newX := p[0]+dir[0], p[1]+dir[1]

			if newY >= 0 && newX >= 0 && newY < h && newX < w && cost[newY][newX] == -1 {
				cost[newY][newX] = dist
				stack = append(stack, [2]int{newY, newX})
			}
		}
	}

	return stack[size:]
}

type MinHeap [][]int // cost, y, x

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *MinHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return popped
}

var dirs = [][]int{
	{0, 0},
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func minCost1(grid [][]int) int {
	w, h := len(grid[0]), len(grid)
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, []int{0, 0, 0})

	for minHeap.Len() > 0 {
		p := heap.Pop(minHeap).([]int)

		if visited[p[1]][p[2]] {
			continue
		}
		visited[p[1]][p[2]] = true

		if p[1] == h-1 && p[2] == w-1 {
			return p[0]
		}

		for i := 1; i <= 4; i++ {
			dir := dirs[i]

			newY, newX := p[1]+dir[0], p[2]+dir[1]

			if newX >= 0 && newY >= 0 && newX < w && newY < h && !visited[newY][newX] {
				if i == grid[p[1]][p[2]] {
					heap.Push(minHeap, []int{p[0], newY, newX})
				} else {
					heap.Push(minHeap, []int{p[0] + 1, newY, newX})
				}
			}
		}
	}

	return 0
}

//	Notes
//	1.	encounter this problem last year, not able to answer it

//	2.	use minHeap to try to find out minimum spam (similar to BFS), but this
//		is not the most efficient way

//	3.	inspired from https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/discuss/524886/JavaC%2B%2BPython-BFS-and-DFS

//		lee uses BFS & DFS to solve the problem, very brilliant
