package main

// Given a matrix and a target, return the number of non-empty submatrices that sum to target.
//
// A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
//
// Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.
//
//
//
// Example 1:
//
// Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
// Output: 4
// Explanation: The four 1x1 submatrices that only contain 0.
//
// Example 2:
//
// Input: matrix = [[1,-1],[-1,1]], target = 0
// Output: 5
// Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.
//
// Example 3:
//
// Input: matrix = [[904]], target = 0
// Output: 0
//
//
//
// Constraints:
//
// 1 <= matrix.length <= 100
// 1 <= matrix[0].length <= 100
// -1000 <= matrix[i] <= 1000
// -10^8 <= target <= 10^8

// tc: O(n*m^2), n: height, m: width
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	w, h := len(matrix[0]), len(matrix)

	prefixSums := make([][]int, h)
	for i := range prefixSums {
		prefixSums[i] = make([]int, w)
	}

	for i := range matrix {
		for j := range matrix[0] {
			if j > 0 {
				prefixSums[i][j] = prefixSums[i][j-1] + matrix[i][j]
			} else {
				prefixSums[i][0] = matrix[i][0]
			}
		}
	}

	var count int

	for x := 0; x < w; x++ {
		for j := x; j < w; j++ {
			table := make(map[int]int)
			table[0] = 1
			sum := 0

			for y := range prefixSums {
				if x > 0 {
					sum += prefixSums[y][j] - prefixSums[y][x-1]
				} else {
					sum += prefixSums[y][j]
				}

				if val, ok := table[sum-target]; ok {
					count += val
				}

				table[sum]++
			}
		}
	}

	return count
}

//	Notes
//	1.	inspired from https://blog.csdn.net/weixin_45588823/article/details/100573071

//		it seems reasonable to solve by dp, so I first go back to solve the
//		prior problem: subarray sums to k again

//		the analogy here is to view each row as the subarray sums.

//		the problem states submatrices, convert it to real situation, it means
//		2D area with all possibilities:

//		e.g 2x2 array
//		x: 0, y: 0, 1, 0~1
//		x: 1, y: 0, 1, 0~1
//		x: 0~1, y: 0, 1, 0~1

//		to use purely prefix sumes, tc will be O(m^2 * n^2)

//		take the subarray sums to k as an analogy, it can view in x or y
//		direction, take x as an example

//		view all row with same x range as a number in subarray sum,
//		e.g. 0 1 0		0 1 1
//		   	 1 1 1  => 	1 2 3	prefix sums
//			 0 1 0		0 1 1

//		view vertically
//		start from x = 0 ~ 0, the 1D subarray becomes [0, 1, 1]
//				   x = 0 ~ 1, the 1D subarray becomes [1, 3, 4]
//				   x = 0 ~ 2, the 1D subarray becomes [1, 4, 5]

//		thus, subarray sum to k can be used

//		but, x doesn't only start from 0, it could be 1~2, 1~2, 2~2, so
//		additional loop for x need to be considered

//	2.	inspired from https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/discuss/803353/Java-Solution-with-Detailed-Explanation

//		author has a very detials graph to demonstrate the concept
