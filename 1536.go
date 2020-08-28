package main

//Given an n x n binary grid, in one step you can choose two adjacent rows of the grid and swap them.
//
//A grid is said to be valid if all the cells above the main diagonal are zeros.
//
//Return the minimum number of steps needed to make the grid valid, or -1 if the grid cannot be valid.
//
//The main diagonal of a grid is the diagonal that starts at cell (1, 1) and ends at cell (n, n).
//
//
//
//Example 1:
//
//Input: grid = [[0,0,1],[1,1,0],[1,0,0]]
//Output: 3
//
//Example 2:
//
//Input: grid = [[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]
//Output: -1
//Explanation: All rows are similar, swaps have no effect on the grid.
//
//Example 3:
//
//Input: grid = [[1,0,0],[1,1,0],[1,1,1]]
//Output: 0
//
//
//
//Constraints:
//
//    n == grid.length
//    n == grid[i].length
//    1 <= n <= 200
//    grid[i][j] is 0 or 1

func minSwaps(grid [][]int) int {
	oneIndex := make([]int, len(grid[0]))

	// find each row 1 index position
	for i := range grid {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				oneIndex[i] = j
				break
			}
		}
	}

	var count, j int

	// sort starts from smallest number, because if a row fits this
	// criteria, it also fits any other row afterwards
	for i := range oneIndex {
		for j = i; j < len(oneIndex); j++ {
			if oneIndex[j] <= i {
				count += j - i
				break
			}
		}

		// not found any suitable row
		if j == len(oneIndex) {
			return -1
		}

		// swap rows
		for ; j > i; j-- {
			oneIndex[j], oneIndex[j-1] = oneIndex[j-1], oneIndex[j]
		}
	}

	return count
}

//	Notes
//	1.	inspired from sample code, count can add first time, without following
//		row swap

//	2.	code can be cleaner, use only j w/o additional variable k
