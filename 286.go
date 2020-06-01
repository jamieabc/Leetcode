package main

import "math"

// You are given a m x n 2D grid initialized with these three possible values.
//
//     -1 - A wall or an obstacle.
//     0 - A gate.
//     INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
//
// Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.
//
// Example:
//
// Given the 2D grid:
//
// INF  -1  0  INF
// INF INF INF  -1
// INF  -1 INF  -1
//   0  -1 INF INF
//
// After running your function, the 2D grid should be:
//
//   3  -1   0   1
//   2   2   1  -1
//   1  -1   2  -1
//   0  -1   3   4

func wallsAndGates(rooms [][]int) {
	stack := make([][]int, 0)

	if len(rooms) == 0 {
		return
	}

	for i := range rooms {
		for j := range rooms[0] {
			if rooms[i][j] == 0 {
				stack = append(stack, []int{i, j})
			}
		}
	}

	dist := 1
	for len(stack) > 0 {
		stop := len(stack) - 1
		for i := 0; i <= stop; i++ {
			s := stack[i]
			for _, j := range adj(rooms, s[0], s[1]) {
				rooms[j[0]][j[1]] = dist
				stack = append(stack, []int{j[0], j[1]})
			}
		}
		dist++
		stack = stack[stop+1:]
	}
}

var steps = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{-1, 0}, // up
	{1, 0},  // down
}

func adj(rooms [][]int, i, j int) [][]int {
	result := make([][]int, 0)
	for _, s := range steps {
		y, x := i+s[0], j+s[1]
		if (x >= 0 && x < len(rooms[0])) && (y >= 0 && y < len(rooms)) {
			if rooms[y][x] == math.MaxInt32 {
				result = append(result, []int{y, x})
			}
		}
	}

	return result
}
