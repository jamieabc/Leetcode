package main

// This question is about implementing a basic elimination algorithm for Candy Crush.
//
// Given a 2D integer array board representing the grid of candy, different positive integers board[i][j] represent different types of candies. A value of board[i][j] = 0 represents that the cell at position (i, j) is empty. The given board represents the state of the game following the player's move. Now, you need to restore the board to a stable state by crushing candies according to the following rules:
//
//     If three or more candies of the same type are adjacent vertically or horizontally, "crush" them all at the same time - these positions become empty.
//     After crushing all candies simultaneously, if an empty space on the board has candies on top of itself, then these candies will drop until they hit a candy or bottom at the same time. (No new candies will drop outside the top boundary.)
//     After the above steps, there may exist more candies that can be crushed. If so, you need to repeat the above steps.
//     If there does not exist more candies that can be crushed (ie. the board is stable), then return the current board.
//
// You need to perform the above rules until the board becomes stable, then return the current board.
//
//
//
// Example:
//
// Input:
// board =
// [[110,5,112,113,114],[210,211,5,213,214],[310,311,3,313,314],[410,411,412,5,414],[5,1,512,3,3],[610,4,1,613,614],[710,1,2,713,714],[810,1,2,1,1],[1,1,2,2,2],[4,1,4,4,1014]]
//
// Output:
// [[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[110,0,0,0,114],[210,0,0,0,214],[310,0,0,113,314],[410,0,0,213,414],[610,211,112,313,614],[710,311,412,613,714],[810,411,512,713,1014]]
//
// Explanation:
//
//
//
// Note:
//
//     The length of board will be in the range [3, 50].
//     The length of board[i] will be in the range [3, 50].
//     Each board[i][j] will initially start as an integer in the range [1, 2000].

func candyCrush(board [][]int) [][]int {
	y := len(board)
	if y == 0 {
		return [][]int{}
	}

	for scan(board) {
		reduce(board)
	}

	return board
}

func scan(board [][]int) bool {
	r := scanRow(board)
	c := scanColumn(board)
	return r || c
}

func scanRow(board [][]int) bool {
	y := len(board)
	x := len(board[0])

	var consecutive, prev, j int
	var found bool

	for i := 0; i < y; i++ {
		for consecutive, prev, j = 1, board[i][0], 1; j < x; j++ {
			if board[i][j] == prev && board[i][j] != 0 {
				consecutive++
			} else {
				if consecutive >= 3 {
					for k := j - 1; consecutive > 0; k, consecutive = k-1, consecutive-1 {
						board[i][k] = -board[i][k]
					}
					found = true
				}

				consecutive = 1
			}
			prev = board[i][j]
		}

		if consecutive >= 3 {
			for k := j - 1; consecutive > 0; k, consecutive = k-1, consecutive-1 {
				board[i][k] = -board[i][k]
			}
			found = true
		}
	}

	return found
}

func scanColumn(board [][]int) bool {
	y := len(board)
	x := len(board[0])

	var prev, consecutive, i int
	var found bool
	for j := 0; j < x; j++ {
		for i, prev, consecutive = 1, abs(board[0][j]), 1; i < y; i++ {
			if abs(board[i][j]) == prev && board[i][j] != 0 {
				consecutive++
			} else {
				if consecutive >= 3 {
					for k := i - 1; consecutive > 0; k, consecutive = k-1, consecutive-1 {
						if board[k][j] > 0 {
							board[k][j] = -board[k][j]
						}
					}
					found = true
				}
				consecutive = 1
			}
			prev = abs(board[i][j])
		}

		if consecutive >= 3 {
			for k := i - 1; consecutive > 0; k, consecutive = k-1, consecutive-1 {
				if board[k][j] > 0 {
					board[k][j] = -board[k][j]
				}
			}
			found = true
		}
	}

	return found
}

func reduce(board [][]int) {
	var writtenRow int
	for j := range board[0] {
		writtenRow = len(board) - 1
		for i := len(board) - 1; i >= 0; i-- {
			if board[i][j] > 0 {
				board[writtenRow][j] = board[i][j]
				writtenRow--
			}
		}

		for ; writtenRow >= 0; writtenRow-- {
			board[writtenRow][j] = 0
		}
	}
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	problems
//	1.	based from discussion, it can be done by no extra memory space. The
//		trick is to mark numbers that are removed next time as negative.

//		reference: https://leetcode.com/problems/candy-crush/discuss/113914/15-ms-Short-Java-Solution-Mark-crush-with-negative-value
