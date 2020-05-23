package main

// Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
// Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
//
// Example 1:
//
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
//
// Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
//
// Example 2:
//
// [[0,0,0,0,0,0,0,0]]
//
// Given the above grid, return 0.
//
// Note: The length of each dimension in the given grid does not exceed 50.

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, traverse(grid, i, j))
			}
		}
	}

	return maxArea
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func traverse(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = -1

	return 1 + traverse(grid, i-1, j) + traverse(grid, i+1, j) + traverse(grid, i, j-1) + traverse(grid, i, j+1)
}
