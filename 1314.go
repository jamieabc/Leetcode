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
	width, height := len(mat[0]), len(mat)

	prefixSum := make([][]int, height)
	for i := range prefixSum {
		prefixSum[i] = make([]int, width)
	}

	for i := range mat {
		tmp := 0
		for j := range mat[0] {
			tmp += mat[i][j]

			if i == 0 {
				prefixSum[i][j] = tmp
			} else {
				prefixSum[i][j] += prefixSum[i-1][j] + tmp
			}
		}
	}

	ans := make([][]int, height)
	for i := range ans {
		ans[i] = make([]int, width)
	}

	for i := range ans {
		for j := range ans[0] {
			ans[i][j] += prefixSum[min(i+K, height-1)][min(j+K, width-1)]

			if i > K {
				ans[i][j] -= prefixSum[i-K-1][min(j+K, width-1)]
			}

			if j > K {
				ans[i][j] -= prefixSum[min(i+K, height-1)][j-K-1]
			}

			if i > K && j > K {
				ans[i][j] += prefixSum[i-K-1][j-K-1]
			}
		}
	}

	return ans
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
