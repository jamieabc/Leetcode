package main

// Write an efficient algorithm that searches for a target value in an m x n integer matrix. The matrix has the following properties:
//
// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.
//
//
//
// Example 1:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// Output: true
//
// Example 2:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// Output: false
//
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -109 <= matix[i][j] <= 109
// All the integers in each row are sorted in ascending order.
// All the integers in each column are sorted in ascending order.
// -109 <= target <= 109

// tc: O(m+n)
func searchMatrix(matrix [][]int, target int) bool {
	w, h := len(matrix[0]), len(matrix)

	if w == 0 || h == 0 {
		return false
	}

	x, y := w-1, 0

	for x >= 0 && y < h {
		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			x--
		} else {
			y++
		}
	}

	return false
}

// tc: O(n log(n)), assumes m ~ n
func searchMatrix1(matrix [][]int, target int) bool {
	w, h := len(matrix[0]), len(matrix)

	if w == 0 || h == 0 {
		return false
	}

	return binarySearch(matrix, target, 0, w-1, 0, h-1)
}

func binarySearch(matrix [][]int, target, left, right, top, bottom int) bool {
	if left > right || top > bottom || target < matrix[top][left] || target > matrix[bottom][right] {
		return false
	}

	x := left + (right-left)/2
	y := top

	// find y such that [y-1][x] < target < [y][x]
	for ; y <= bottom && matrix[y][x] <= target; y++ {
		if matrix[y][x] == target {
			return true
		}
	}

	return binarySearch(matrix, target, left, x-1, y, bottom) || binarySearch(matrix, target, x+1, right, top, y-1)
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/66160/AC-clean-Java-solution

//		author provides a very brilliant insight that start from top-right,
//		if target < current, go left; target > current go down

//		because last number of each line is largest, and last line from ascending
//		order, so it can guarantee that if a number is larger than current,
//		need to go next line

//	2.	inspired from solution, divide conquer can help to solve this problem

//		the point here is to find a range that [row-1][mid] < target < [row][mid]

//		|		|		|
//		|		|		|
//		|		|		|
//		|-------|-------| 	row
//		|		|		|
//		|		|		|
//		left	mid		right

//		since [row-1][mid] < target < [row][mid], no need to search top-left
//		and bottom-right

//	3.	inspired form https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/66147/*Java*-an-easy-to-understand-divide-and-conquer-method

//		similar to solution, but gets worse because it considers 3 areas
