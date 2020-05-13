package main

// Given a m * n matrix mat and an integer K, return a matrix answer where each answer[i][j] is the sum of all elements mat[r][c] for i - K <= r <= i + K, j - K <= c <= j + K, and (r, c) is a valid position in the matrix.
//
//
//
// Example 1:
//
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 1
// Output: [[12,21,16],[27,45,33],[24,39,28]]
//
// Example 2:
//
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 2
// Output: [[45,45,45],[45,45,45],[45,45,45]]
//
//
//
// Constraints:
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n, K <= 100
// 1 <= mat[i][j] <= 100

func matrixBlockSum(mat [][]int, K int) [][]int {
	y := len(mat)
	x := len(mat[0])
	result := make([][]int, y)
	for i := range result {
		result[i] = make([]int, x)
	}
	dp := make([]int, x)

	// setup dp
	for i := 0; i <= K; i++ {
		for tmp, j := 0, 0; j < x; j++ {
			tmp += mat[i][j]
			dp[j] += tmp
		}
	}

	var left, sum int
	for i := range mat {
		for j := range mat[0] {
			sum = dp[min(j+K, x-1)]

			if j-K-1 >= 0 {
				left = dp[j-K-1]
			} else {
				left = 0
			}

			result[i][j] = sum - left
		}

		// out of range row from next row
		if i-K >= 0 {
			sum = 0
			for j := range mat[0] {
				sum += mat[i-K][j]
				dp[j] -= sum
			}
		}

		// newly added row from next row
		if i+1+K < y {
			sum = 0
			for j := range mat[0] {
				sum += mat[i+1+K][j]
				dp[j] += sum
			}
		}
	}

	return result
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	too slow, the complexity is O(mnk)

//	2.	from hint 2, use cumulative array to solve this

//	3.	when calculating cumulative, be ware that range should be +K, this
//		problem is about changing cumulative base, I spend 1 hour to find
//		the bug

//	4.	from sample code, those if can be replaced by min/max

//	5.	from reference https://leetcode.com/problems/matrix-block-sum/discuss/477041/Java-Prefix-sum-with-Picture-explain-Clean-code-O(m*n)

//		it's a technique to use size + 1 array to avoid a lot of boundary
//		limit checking, e.g. i >= 1, j >= 1, etc.
