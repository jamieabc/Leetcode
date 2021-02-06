package main

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
// Example:
//
// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

func minPathSum(grid [][]int) int {
	w, h := len(grid[0]), len(grid)

	// dp[i][j]: max point at [i, j]
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
	}

	dp[0][0] = grid[0][0]
	for j := 1; j < w; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < h; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[h-1][w-1]
}

func minPathSum2(grid [][]int) int {
	y := len(grid)
	if y == 0 {
		return 0
	}
	x := len(grid[0])
	dp := make([]int, x)
	dp[0] = grid[0][0]

	for i := 1; i < x; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < y; i++ {
		for j := range grid[0] {
			if j > 0 {
				dp[j] = grid[i][j] + min(dp[j-1], dp[j])
			} else {
				dp[j] = grid[i][j] + dp[j]
			}
		}
	}

	return dp[x-1]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func minPathSum1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	return traverse(grid, len(grid)-1, len(grid[0])-1)
}

func traverse(grid [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if i > 0 && j > 0 {
		return grid[i][j] + min(traverse(grid, i-1, j), traverse(grid, i, j-1))
	} else if i != 0 {
		return grid[i][j] + traverse(grid, i-1, j)
	} else if j != 0 {
		return grid[i][j] + traverse(grid, i, j-1)
	} else {
		return grid[0][0]
	}
}

//  problems

//  1.  wrong index

//	2.	array could not be square, dp size is column size

//	3.	inspired from https://leetcode.com/problems/minimum-path-sum/discuss/344980/Java.-Details-from-Recursion-to-DP.

//		author uses recursion
