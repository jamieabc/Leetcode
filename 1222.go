package main

// On an 8x8 chessboard, there can be multiple Black Queens and one White King.
//
// Given an array of integer coordinates queens that represents the positions of the Black Queens, and a pair of coordinates king that represent the position of the White King, return the coordinates of all the queens (in any order) that can attack the King.
//
//
//
// Example 1:
//
// Input: queens = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], king = [0,0]
// Output: [[0,1],[1,0],[3,3]]
// Explanation:
// The queen at [0,1] can attack the king cause they're in the same row.
// The queen at [1,0] can attack the king cause they're in the same column.
// The queen at [3,3] can attack the king cause they're in the same diagnal.
// The queen at [0,4] can't attack the king cause it's blocked by the queen at [0,1].
// The queen at [4,0] can't attack the king cause it's blocked by the queen at [1,0].
// The queen at [2,4] can't attack the king cause it's not in the same row/column/diagnal as the king.
//
// Example 2:
//
// Input: queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]
// Output: [[2,2],[3,4],[4,4]]
//
// Example 3:
//
// Input: queens = [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]], king = [3,4]
// Output: [[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]
//
//
//
// Constraints:
//
//     1 <= queens.length <= 63
//     queens[0].length == 2
//     0 <= queens[i][j] < 8
//     king.length == 2
//     0 <= king[0], king[1] < 8
//     At most one piece is allowed in a cell.

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	mapping := make(map[int]bool)
	size := 8
	result := make([][]int, 0)

	for _, q := range queens {
		mapping[q[0]*size+q[1]] = true
	}

	type direction int
	const (
		upLeft direction = iota
		up
		upRight
		left
		right
		downLeft
		down
		downRight
	)

	steps := [][]int{
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // down-left
		{1, 0},   // down
		{1, 1},   // down-right
	}

	dirs := make([][]int, size)

	for count, i := 0, 1; count < size && i <= max(max(size-king[0], king[0]), max(size-king[1], king[1])); i++ {
		for j := range steps {
			if len(dirs[j]) == 0 {
				y, x := king[0]+i*steps[j][0], king[1]+i*steps[j][1]
				if y < 0 || x < 0 || x >= size || y >= size {
					continue
				}

				if _, ok := mapping[y*size+x]; ok {
					dirs[j] = []int{y, x}
					count++
				}
			}
		}
	}

	for i := range dirs {
		if len(dirs[i]) > 0 {
			result = append(result, dirs[i])
		}
	}

	return result
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	inspired from https://leetcode.com/problems/queens-that-can-attack-the-king/discuss/403669/Python-Check-8-steps-in-8-Directions

//		transform (i, j) => i * N + j, which can be set in a map, then start
//		from king's position distance from 1 to 8 for all directions.
//		with this method, first found will be closest

//	2.	be careful about negative index

//	3.	be careful about boundaries, the index should be at most N-1

//	4.	inspired from https://leetcode.com/problems/queens-that-can-attack-the-king/discuss/403755/C%2B%2B-Tracing

//		I can set max distance for looping
