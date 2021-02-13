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
	if board[click[0]][click[1]] == byte('M') {
		board[click[0]][click[1]] = byte('X')
		return board
	}

	queue := [][]int{click}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		// already processed
		if board[p[0]][p[1]] == byte('B') {
			continue
		}

		mine := adjMines(board, p)
		if mine == 0 {
			board[p[0]][p[1]] = byte('B')

			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					if i == 0 && j == 0 {
						continue
					}

					newY, newX := p[0]+i, p[1]+j

					if validPoint(board, newY, newX) && board[newY][newX] == byte('E') {
						queue = append(queue, []int{newY, newX})
					}
				}
			}
		} else {
			board[p[0]][p[1]] = byte('0'+mine)
		}
	}

	return board
}

func adjMines(board [][]byte, pos []int) int {
	var mine int

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}

			newY, newX := pos[0]+i, pos[1]+j

			if validPoint(board, newY, newX) && board[newY][newX] == byte('M') {
				mine++
			}
		}
	}

	return mine
}

func validPoint(board [][]byte, y, x int) bool {
	w, h := len(board[0]), len(board)

	return y >= 0 && x >= 0 && y < h && x < w
}

//	Notes
//	1.	too slow, don't use stack for too many memory update

//	2.	this is BFS from click position, only thing to be aware is that all 8 directions
//		are considered as neighbor, will be put into queue

//	3.	visited might not be necessary, but becareful causes infinite loop
//		because there are 8 directions, same position might be added multiple time, so
//		need to check if any point is already processed

//	4.	inspired from https://leetcode.com/problems/minesweeper/discuss/99841/Straight-forward-Java-solution

//		could also use to denote
//		int[] dx = {-1, 0, 1, -1, 1, 0, 1, -1};
//    	int[] dy = {-1, 1, 1, 0, -1, -1, 0, 1};