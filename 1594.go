package main

//You are given a rows x cols matrix grid. Initially, you are located at the top-left corner (0, 0), and in each step, you can only move right or down in the matrix.
//
//Among all possible paths starting from the top-left corner (0, 0) and ending in the bottom-right corner (rows - 1, cols - 1), find the path with the maximum non-negative product. The product of a path is the product of all integers in the grid cells visited along the path.
//
//Return the maximum non-negative product modulo 109 + 7. If the maximum product is negative return -1.
//
//Notice that the modulo is performed after getting the maximum product.
//
//
//
//Example 1:
//
//Input: grid = [[-1,-2,-3],
//               [-2,-3,-3],
//               [-3,-3,-2]]
//Output: -1
//Explanation: It's not possible to get non-negative product in the path from (0, 0) to (2, 2), so return -1.
//
//Example 2:
//
//Input: grid = [[1,-2,1],
//               [1,-2,1],
//               [3,-4,1]]
//Output: 8
//Explanation: Maximum non-negative product is in bold (1 * 1 * -2 * -4 * 1 = 8).
//
//Example 3:
//
//Input: grid = [[1, 3],
//               [0,-4]]
//Output: 0
//Explanation: Maximum non-negative product is in bold (1 * 0 * -4 = 0).
//
//Example 4:
//
//Input: grid = [[ 1, 4,4,0],
//               [-2, 0,0,1],
//               [ 1,-1,1,1]]
//Output: 2
//Explanation: Maximum non-negative product is in bold (1 * -2 * 1 * -1 * 1 * 1 = 2).
//
//
//
//Constraints:
//
//    1 <= rows, cols <= 15
//    -4 <= grid[i][j] <= 4

func maxProductPath(grid [][]int) int {
	w, h := len(grid[0]), len(grid)

	dp := make([][][2]int, h)
	for i := range dp {
		dp[i] = make([][2]int, w)
	}

	dp[0][0][0] = grid[0][0]
	dp[0][0][1] = grid[0][0]

	// traverse right
	for j := 1; j < w; j++ {
		dp[0][j][0] = dp[0][j-1][0] * grid[0][j]
		dp[0][j][1] = dp[0][j][0]
	}

	// traverse down
	for i := 1; i < h; i++ {
		dp[i][0][0] = dp[i-1][0][0] * grid[i][0]
		dp[i][0][1] = dp[i][0][0]
	}

	// traverse
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if grid[i][j] > 0 {
				dp[i][j][0] = grid[i][j] * max(dp[i-1][j][0], dp[i][j-1][0])
				dp[i][j][1] = grid[i][j] * min(dp[i-1][j][1], dp[i][j-1][1])
			} else {
				dp[i][j][0] = grid[i][j] * min(dp[i-1][j][1], dp[i][j-1][1])
				dp[i][j][1] = grid[i][j] * max(dp[i-1][j][0], dp[i][j-1][0])
			}
		}
	}

	if dp[h-1][w-1][0] < 0 {
		return -1
	}

	return dp[h-1][w-1][0] % int(1e9+7)
}

type Val struct {
	Max, Min int64
}

func maxProductPath1(grid [][]int) int {
	mod := int64(1e9 + 7)
	dp := make([][]Val, len(grid))

	for i := range dp {
		dp[i] = make([]Val, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = Val{
				Max: 1,
				Min: 1,
			}
		}
	}

	// setup initial value on first column
	for i := range dp {
		if i == 0 {
			dp[0][0].Max = int64(grid[0][0])
			dp[0][0].Min = int64(grid[0][0])
		} else {
			dp[i][0].Max = dp[i-1][0].Max * int64(grid[i][0])
			dp[i][0].Min = dp[i][0].Max
		}
	}

	// setup initial value on first row
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i].Max = dp[0][i-1].Max * int64(grid[0][i])
		dp[0][i].Min = dp[0][i].Max
	}

	// calculate maximum product from previous
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j].Max, dp[i][j].Min = update(dp, grid, i, j)
		}
	}

	// cannot get positive product after traversal
	if dp[len(dp)-1][len(dp[0])-1].Max < 0 {
		return -1
	}

	return int(dp[len(dp)-1][len(dp[0])-1].Max % mod)
}

func update(dp [][]Val, grid [][]int, y, x int) (int64, int64) {
	var m, n int64
	if grid[y][x] >= 0 {
		m = max(dp[y][x-1].Max, dp[y-1][x].Max) * int64(grid[y][x])
		n = min(dp[y][x-1].Min, dp[y-1][x].Min) * int64(grid[y][x])
	} else {
		m = min(dp[y][x-1].Min, dp[y-1][x].Min) * int64(grid[y][x])
		n = max(dp[y][x-1].Max, dp[y-1][x].Max) * int64(grid[y][x])
	}

	return m, n
}

func max(i, j int64) int64 {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int64) int64 {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	only care about left/up max/min
