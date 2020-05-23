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

var steps = [][2]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func generateMatrix(n int) [][]int {
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
