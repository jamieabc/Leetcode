package main

// You are given an n x n 2D matrix representing an image.
//
// Rotate the image by 90 degrees (clockwise).
//
// Note:
//
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
// Example 1:
//
// Given input matrix =
// [
// [1,2,3],
// [4,5,6],
// [7,8,9]
// ],
//
// rotate the input matrix in-place such that it becomes:
// [
// [7,4,1],
// [8,5,2],
// [9,6,3]
// ]
//
// Example 2:
//
// Given input matrix =
// [
// [ 5, 1, 9,11],
// [ 2, 4, 8,10],
// [13, 3, 6, 7],
// [15,14,12,16]
// ],
//
// rotate the input matrix in-place such that it becomes:
// [
// [15,13, 2, 5],
// [14, 3, 4, 1],
// [12, 6, 8, 9],
// [16, 7,10,11]
// ]

func rotate(matrix [][]int) {
	n := len(matrix)

	// flip top & down
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}

	// matrix transpose
	for i := range matrix {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate1(matrix [][]int) {
	size := len(matrix)
	var next int

	for rounds := 0; rounds < size/2; rounds++ {
		for start, end := rounds, size-rounds-1; start < end; start++ {
			matrix[start][size-1-rounds], next = matrix[rounds][start], matrix[start][size-1-rounds]
			matrix[size-1-rounds][size-1-start], next = next, matrix[size-1-rounds][size-1-start]
			matrix[size-1-start][rounds], next = next, matrix[size-1-start][rounds]
			matrix[rounds][start] = next
		}

	}
}

//  Notes
//  1.  reference from https://leetcode.com/problems/rotate-image/discuss/18872/A-common-method-to-rotate-the-image

//      clockwise rotation can be done by flipping tops & downs then do
//      matrix transpose

//      10 11 12 13        22 23 24 25        22 18 14 10
//      14 15 16 17   =>   18 19 20 21   =>   23 19 15 11
//      18 19 20 21        14 15 16 17        24 15 16 12
//      22 23 24 25        10 11 12 13        25 11 17 13

//      counter-clockwise rotation can be done by swap first column and do
//      matrix transpose

//      10 11 12 13        13 11 12 10        13 17 21 25
//      14 15 16 17   =>   17 15 16 14   =>   11 15 19 23
//      18 19 20 21        21 19 20 18        12 16 20 24
//      22 23 24 25        25 23 24 22        10 14 18 22
