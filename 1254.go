package main

// Given a 2D grid consists of 0s (land) and 1s (water).  An island is a maximal 4-directionally connected group of 0s and a closed island is an island totally (all left, top, right, bottom) surrounded by 1s.
//
// Return the number of closed islands.
//
//
//
// Example 1:
//
// Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
// Output: 2
// Explanation:
// Islands in gray are closed because they are completely surrounded by water (group of 1s).
//
// Example 2:
//
// Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
// Output: 1
//
// Example 3:
//
// Input: grid = [[1,1,1,1,1,1,1],
// [1,0,0,0,0,0,1],
// [1,0,1,1,1,0,1],
// [1,0,1,0,1,0,1],
// [1,0,1,1,1,0,1],
// [1,0,0,0,0,0,1],
// [1,1,1,1,1,1,1]]
// Output: 2
//
//
//
// Constraints:
//
// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1

func closedIsland(grid [][]int) int {
	w, h := len(grid[0]), len(grid)

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var closedIsland int

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 && !visited[i][j] {
				if dfs(grid, visited, i, j) {
					closedIsland++
				}
			}
		}
	}

	return closedIsland
}

func dfs(grid [][]int, visited [][]bool, y, x int) bool {
	w, h := len(grid[0]), len(grid)

	if !(x >= 0 && y >= 0 && x < w && y < h) {
		return false
	}

	if grid[y][x] == 1 || visited[y][x] {
		return true
	}
	visited[y][x] = true

	inside := true

	if x == 0 || y == 0 || x == w-1 || y == h-1 {
		inside = false
	}

	inside = dfs(grid, visited, y+1, x) && inside
	inside = dfs(grid, visited, y, x+1) && inside
	inside = dfs(grid, visited, y-1, x) && inside
	inside = dfs(grid, visited, y, x-1) && inside

	return inside
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/number-of-closed-islands/discuss/425150/JavaC%2B%2B-with-picture-Number-of-Enclaves

//		another way to solve it is to first fill surrounded area of boundary, then count

//		also, author uses + to denote any invalid encounter

//		the core of this is to first remove cells should not be calculated, then do normal
//		dfs

//	2.	inspired from https://leetcode.com/problems/number-of-closed-islands/discuss/426294/JavaPython-3-DFS-BFS-and-Union-Find-codes-w-brief-explanation-and-analysis.

//		author provides another way of solving the problem: union-find
