package main

// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
//
// Example:
//
// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

var dirs = [][]int{
	{1, 0},  // east
	{0, 1},  // south
	{-1, 0}, // west
	{0, -1}, // north
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	x, y := -1, 0

	for i, dir, steps := 1, 0, n; i <= n*n; {
		for j := 0; j < steps; i, j = i+1, j+1 {
			x, y = x+dirs[dir][0], y+dirs[dir][1]

			ans[y][x] = i
		}

		dir = (dir + 1) % 4
		if dir == 1 || dir == 3 {
			steps--
		}
	}

	return ans
}

var steps = [][2]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func generateMatrix1(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i, j, num, dir := 0, 0, 1, 0; num <= n*n; num++ {
		result[i][j] = num

		newI, newJ := i+steps[dir][0], j+steps[dir][1]
		if newI < 0 || newI == n || newJ < 0 || newJ == n || result[newI][newJ] != 0 {
			dir = (dir + 1) % 4
			i += steps[dir][0]
			j += steps[dir][1]
		} else {
			i, j = newI, newJ
		}
	}

	return result
}

//	problems
//	1.	when reset direction up -> right, with a typo
