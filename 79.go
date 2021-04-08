package main

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
//
//
// Example 1:
//
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true
//
// Example 2:
//
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// Output: true
//
// Example 3:
//
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// Output: false
//
//
//
// Constraints:
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.
//
//
//
// Follow up: Could you use search pruning to make your solution faster with a larger board?

// tc: O(n * 3^m), n: number of cells in board, m: length of word
func exist(board [][]byte, word string) bool {
	w, h := len(board[0]), len(board)
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				visited[i][j] = true

				if backtracking(board, word, visited, i, j, 1) {
					return true
				}

				visited[i][j] = false
			}
		}
	}

	return false
}

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func backtracking(board [][]byte, word string, visited [][]bool, y, x, idx int) bool {
	if idx == len(word) {
		return true
	}

	w, h := len(board[0]), len(board)
	var result bool

	for _, dir := range dirs {
		newY, newX := y+dir[0], x+dir[1]

		if newX >= 0 && newY >= 0 && newX < w && newY < h && !visited[newY][newX] && board[newY][newX] == word[idx] {
			visited[newY][newX] = true

			result = result || backtracking(board, word, visited, newY, newX, idx+1)

			visited[newY][newX] = false
		}
	}

	return result
}

//	Notes
//	1.	this is not bfs, bfs usually don't care about order, cares about which
//		destination can be reached or not, but word search cares about order

//		the technique used in bfs cannot be applies

//		e.g.
//		  a a a
//		b a a a

//		to find aaaaab, if start from [0][2]
//		if [0][2] starts first from left, with visited, it would be

//		  t a t
//		f f t f

//		so right side of a ([0][3]) never has the ability to go to left side

//		I stuck at this point for quite a long time

//	2.	inspired from solution, no need to allocate visited everytime, can
//		clean-up after each traversing
