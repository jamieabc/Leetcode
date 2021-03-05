package main

import "math"

// Given a square array of integers A, we want the minimum sum of a falling path through A.
//
// A falling path starts at any element in the first row, and chooses one element from each row.  The next row's choice must be in a column that is different from the previous row's column by at most one.
//
//
//
// Example 1:
//
// Input: [[1,2,3],[4,5,6],[7,8,9]]
// Output: 12
// Explanation:
// The possible falling paths are:
//
//     [1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
//     [2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
//     [3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
//
// The falling path with the smallest sum is [1,4,7], so the answer is 12.
//
//
//
// Note:
//
//     1 <= A.length == A[0].length <= 100
//     -100 <= A[i][j] <= 100

func minFallingPathSum(matrix [][]int) int {
	w, h := len(matrix[0]), len(matrix)

	// dp[i][j]: minimum sum from [i, j] downward
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)

		for j := range dp[0] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := range matrix[0] {
		dp[h-1][i] = matrix[h-1][i]
	}

	for i := h - 2; i >= 0; i-- {
		for j := range dp[0] {
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1]+matrix[i][j])
			}

			dp[i][j] = min(dp[i][j], dp[i+1][j]+matrix[i][j])

			if j < w-1 {
				dp[i][j] = min(dp[i][j], dp[i+1][j+1]+matrix[i][j])
			}
		}
	}

	ans := math.MaxInt32
	for i := range dp[0] {
		ans = min(ans, dp[0][i])
	}

	return ans
}

func minFallingPathSum1(A [][]int) int {
	length := len(A)

	if length == 1 {
		return A[0][0]
	}

	for i := 1; i < length; i++ {
		for j := 0; j < length; j++ {
			if j == 0 {
				A[0][j] += min(A[i][0], A[i][1])
			} else if j == length-1 {
				A[0][j] += min(A[i][j], A[i][j-1])
			} else {
				A[0][j] += min(A[i][j], min(A[i][j-1], A[i][j+1]))
			}
		}
	}

	cost := 100 * length
	for i := range A {
		cost = min(cost, A[0][i])
	}

	return cost
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	this is wrong, if a number is chosen, left choices are only one
//		direction, e.g. if choose right number, then choices are right
//		right-1, right+1

//		if number can be duplicate, then selection range is even more
//		increased in next line, be decision cannot be made, it has to
//		calculate all possibilities
