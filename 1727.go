package main

import (
	"math"
	"sort"
)

// You are given a binary matrix matrix of size m x n, and you are allowed to rearrange the columns of the matrix in any order.
//
// Return the area of the largest submatrix within matrix where every element of the submatrix is 1 after reordering the columns optimally.
//
//
//
// Example 1:
//
// Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]
// Output: 4
// Explanation: You can rearrange the columns as shown above.
// The largest submatrix of 1s, in bold, has an area of 4.
//
// Example 2:
//
// Input: matrix = [[1,0,1,0,1]]
// Output: 3
// Explanation: You can rearrange the columns as shown above.
// The largest submatrix of 1s, in bold, has an area of 3.
//
// Example 3:
//
// Input: matrix = [[1,1,0],[1,0,1]]
// Output: 2
// Explanation: Notice that you must rearrange entire columns, and there is no way to make a submatrix of 1s larger than an area of 2.
//
// Example 4:
//
// Input: matrix = [[0,0],[0,0]]
// Output: 0
// Explanation: As there are no 1s, no submatrix of 1s can be formed and the area is 0.
//
//
//
// Constraints:
//
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m * n <= 105
//     matrix[i][j] is 0 or 1.

// O(mn log(n))
func largestSubmatrix(matrix [][]int) int {
	w, h := len(matrix[0]), len(matrix)
	// dp[i][j]: # of consecutive 1 up from specific row counting upward
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
	}

	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 1 {
				if i > 0 {
					dp[i][j] = dp[i-1][j] + 1
				} else {
					dp[i][j] = 1
				}
			}
		}
	}

	var maxArea int

	for i := h - 1; i >= 0; i-- {
		sort.Ints(dp[i])

		// numbers are sorted ascending, so start from right-most number means
		// higher height
		for j := w - 1; j >= 0; j-- {
			maxArea = max(maxArea, dp[i][j]*(w-j))
		}
	}

	return maxArea
}

// tc: O(mn log(n))
func largestSubmatrix1(matrix [][]int) int {
	// dp[i]: # of consecutive 1 start from specific row counting downward
	w, h := len(matrix[0]), len(matrix)
	dp := make([]int, w)

	for j := range matrix[0] {
		for i := 0; i < h; i++ {
			if matrix[i][j] == 0 {
				break
			} else {
				dp[j]++
			}
		}
	}
	maxArea := findMaxArea(dp)

	for i := 1; i < h; i++ {
		for j := range matrix[0] {
			if matrix[i-1][j] == 1 {
				dp[j]--
			} else {
				dp[j] = 0
				for k := i; k < h; k++ {
					if matrix[k][j] == 0 {
						break
					} else {
						dp[j]++
					}
				}
			}
		}
		maxArea = max(maxArea, findMaxArea(dp))
	}

	return maxArea
}

// O(n log(n))
func findMaxArea(dp []int) int {
	tmp := make([]int, len(dp))
	copy(tmp, dp)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})

	var area, w int
	height := math.MaxInt32

	for i := range tmp {
		if tmp[i] == 0 {
			break
		}
		w++

		height = min(height, tmp[i])
		area = max(area, w*height)
	}

	return area
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	the concept of dp is correct, but there's better way of storing data:
//		convert matrix into consecutive 1s downward

//		e.g. [0, 0, 1]		[0, 0, 3]
//			 [1, 1, 1] 	=>	[2, 1, 2]
//			 [1, 0, 1]		[1, 0, 1]

//		when counting this, scan order will be row i, find column j is 1, keep
//		scanning row i+1 same column j to check if it's 1

//		this method works, but can be further optimized by storing consecutive
//		1s upward, thus dp can be utilized and no

//		e.g. [0, 0, 1]		[0, 0, 1]
//			 [1, 1, 1] 	=>	[1, 1, 2]
//			 [1, 0, 1]		[2, 0, 3]
