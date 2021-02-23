package main

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//
//     Integers in each row are sorted from left to right.
//     The first integer of each row is greater than the last integer of the previous row.
//
//
//
// Example 1:
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
//
// Example 2:
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false
//
//
//
// Constraints:
//
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 100
//     -104 <= matrix[i][j], target <= 104

func searchMatrix(matrix [][]int, target int) bool {
	w, h := len(matrix[0]), len(matrix)

	for low, high := 0, w*h-1; low <= high; {
		mid := low + (high-low)/2

		y, x := mid/w, mid%w

		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
