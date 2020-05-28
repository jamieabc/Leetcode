package main

// Let's play the minesweeper game (Wikipedia, online game)!
//
// You are given a 2D char matrix representing the game board. 'M' represents an unrevealed mine, 'E' represents an unrevealed empty square, 'B' represents a revealed blank square that has no adjacent (above, below, left, right, and all 4 diagonals) mines, digit ('1' to '8') represents how many mines are adjacent to this revealed square, and finally 'X' represents a revealed mine.
//
// Now given the next click position (row and column indices) among all the unrevealed squares ('M' or 'E'), return the board after revealing this position according to the following rules:
//
// If a mine ('M') is revealed, then the game is over - change it to 'X'.
// If an empty square ('E') with no adjacent mines is revealed, then change it to revealed blank ('B') and all of its adjacent unrevealed squares should be revealed recursively.
// If an empty square ('E') with at least one adjacent mine is revealed, then change it to a digit ('1' to '8') representing the number of adjacent mines.
// Return the board when no more squares will be revealed.
//
//
// Example 1:
//
// Input:
//
// [['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'M', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E']]
//
// Click : [3,0]
//
// Output:
//
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
//
// Explanation:
//
// Example 2:
//
// Input:
//
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
//
// Click : [1,2]
//
// Output:
//
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'X', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
//
// Explanation:
//
//
//
// Note:
//
// The range of the input matrix's height and width is [1,50].
// The click position will only be an unrevealed square ('M' or 'E'), which also means the input board contains at least one clickable square.
// The input board won't be a stage when game is over (some mines have been revealed).
// For simplicity, not mentioned rules should be ignored in this problem. For example, you don't need to reveal all the unrevealed mines when the game is over, consider any cases that you will win the game or flag any squares.

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	update(board, click[0], click[1])
	for _, i := range adj(board, click[0], click[1]) {
		_ = updateBoard(board, []int{i[0], i[1]})
	}

	return board
}

func isValid(board [][]byte, y, x int) bool {
	if x < 0 || y < 0 || x == len(board[0]) || y == len(board) {
		return false
	}
	return true
}

func update(board [][]byte, y, x int) {
	if !isValid(board, y, x) {
		return
	}

	if board[y][x] == 'E' {
		b := adjBomb(board, y, x)
		if b == 0 {
			board[y][x] = 'B'
		} else {
			board[y][x] = byte(b + '0')
		}
	}
}

func adjBomb(board [][]byte, y, x int) int {
	var bomb int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}

			if isBomb(board, i, j) {
				bomb++
			}
		}
	}

	return bomb
}

func isBomb(board [][]byte, y, x int) bool {
	if !isValid(board, y, x) {
		return false
	}

	return board[y][x] == 'M'
}

func isRevealable(board [][]byte, y, x int) bool {
	if !isValid(board, y, x) {
		return false
	}

	if board[y][x] == 'E' {
		return true
	}

	return false
}

// return with y, x
func adj(board [][]byte, y, x int) [][]int {
	result := make([][]int, 0)

	// adj only move forward when it's blank (B)
	if board[y][x] != 'B' {
		return result
	}

	// bottom
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == x && j == y {
				continue
			}

			if isRevealable(board, i, j) {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

//	problems
//	1.	too slow, don't use stack for too many memory update
