package main

// You are given an m x n binary matrix matrix.
//
// You can choose any number of columns in the matrix and flip every cell in that column (i.e., Change the value of the cell from 0 to 1 or vice versa).
//
// Return the maximum number of rows that have all values equal after some number of flips.
//
//
//
// Example 1:
//
// Input: matrix = [[0,1],[1,1]]
// Output: 1
// Explanation: After flipping no values, 1 row has all values equal.
//
// Example 2:
//
// Input: matrix = [[0,1],[1,0]]
// Output: 2
// Explanation: After flipping values in the first column, both rows have equal values.
//
// Example 3:
//
// Input: matrix = [[0,0,0],[0,0,1],[1,1,0]]
// Output: 2
// Explanation: After flipping values in the first two columns, the last two rows have equal values.
//
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] is either 0 or 1.

func maxEqualRowsAfterFlips(matrix [][]int) int {
	table := make(map[string][]int)
	w := len(matrix[0])
	var largest int

	for i := range matrix {
		line := make([]byte, w)

		for j := range matrix[i] {
			if matrix[i][j] == 1 {
				line[j] = '1'
			} else {
				line[j] = '0'
			}
		}

		table[string(line)] = append(table[string(line)], i)
		flipped := flip(line)
		table[string(flipped)] = append(table[string(flipped)], i)
		largest = max(largest, len(table[string(line)]))
	}

	return largest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func flip(line []byte) []byte {
	flipped := make([]byte, len(line))

	for i := range line {
		if line[i] == '0' {
			flipped[i] = '1'
		} else {
			flipped[i] = '0'
		}
	}

	return flipped
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/discuss/303897/Java-easy-solution-%2B-explanation

//		can do max when traversing matrix, no need to scan again
