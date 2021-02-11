package main

// Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
//
//
//
// Example 1:
//
// Input: matrix =
// [
//   [0,1,1,1],
//   [1,1,1,1],
//   [0,1,1,1]
// ]
// Output: 15
// Explanation:
// There are 10 squares of side 1.
// There are 4 squares of side 2.
// There is  1 square of side 3.
// Total number of squares = 10 + 4 + 1 = 15.
//
// Example 2:
//
// Input: matrix =
// [
//   [1,0,1],
//   [1,1,0],
//   [1,1,0]
// ]
// Output: 7
// Explanation:
// There are 6 squares of side 1.
// There is 1 square of side 2.
// Total number of squares = 6 + 1 = 7.
//
//
//
// Constraints:
//
//     1 <= arr.length <= 300
//     1 <= arr[0].length <= 300
//     0 <= arr[i][j] <= 1

func countSquares(matrix [][]int) int {
	y := len(matrix)
	if y == 0 {
		return 0
	}
	x := len(matrix[0])

	dp := make([]int, x)
	var count, prev int

	for i := range matrix {
		prev = 0
		for j := range matrix[0] {
			if i == 0 || j == 0 {
				prev, dp[j] = dp[j], matrix[i][j]
			} else {
				if matrix[i][j] == 1 && prev > 0 && dp[j-1] > 0 && dp[j] > 0 {
					prev, dp[j] = dp[j], min(min(dp[j-1], dp[j]), prev)+1
				} else {
					prev, dp[j] = dp[j], matrix[i][j]
				}
			}

			if dp[j] > 0 {
				count += dp[j]
			}
		}
	}

	return count
}

func countSquares1(matrix [][]int) int {
	y := len(matrix)
	if y == 0 {
		return 0
	}
	x := len(matrix[0])

	dp := make([][]int, y)
	for i := range dp {
		dp[i] = make([]int, x)
	}

	var count int

	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 1 {
				if i > 0 && j > 0 {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				} else {
					dp[i][j] = 1
				}
				count += dp[i][j]
			}
		}
	}

	return count
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	wrong logic when calculate square

//	2.	i think it can be done in 1D dp, because dp calculation only relates
//		to dp[i-1][j], dp[i][j-1], dp[i-1][j-1]

//	3.	wrong place to update/reset prev

//	4.	reference from sample code and https://leetcode.com/problems/count-square-submatrices-with-all-ones/discuss/441306/JavaC%2B%2BPython-DP-solution

//		it can use original matrix to store dp, so not extra memory space is needed

//		description of dp is precise, dp[i][j] means the size of biggest square with A[i][j] as bottom-right corner.

//	5.	inspired from https://leetcode.com/problems/count-square-submatrices-with-all-ones/discuss/441620/DP-with-figure-explanation

//		author provides thinking process of dp, basically find transformation among each state