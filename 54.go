package main

// Given an m x n matrix, return all elements of the matrix in spiral order.
//
//
//
// Example 1:
//
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
//
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	spiral := make([]int, 0)
	var dir int
	w, h := len(matrix[0]), len(matrix)
	x, y := -1, 0
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	for len(spiral) < w*h {
		// change direction
		newY, newX := y+dirs[dir][0], x+dirs[dir][1]

		if newX < 0 || newX == w || newY < 0 || newY == h || visited[newY][newX] {
			dir = (dir + 1) % 4
			continue
		}

		x, y = newX, newY
		visited[y][x] = true
		spiral = append(spiral, matrix[y][x])
	}

	return spiral
}
