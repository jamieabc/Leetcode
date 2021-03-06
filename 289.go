package main

// According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."
//
// The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):
//
//     Any live cell with fewer than two live neighbors dies as if caused by under-population.
//     Any live cell with two or three live neighbors lives on to the next generation.
//     Any live cell with more than three live neighbors dies, as if by over-population.
//     Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
//
// The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.
//
//
//
// Example 1:
//
// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
//
// Example 2:
//
// Input: board = [[1,1],[1,0]]
// Output: [[1,1],[1,1]]
//
//
//
// Constraints:
//
//     m == board.length
//     n == board[i].length
//     1 <= m, n <= 25
//     board[i][j] is 0 or 1.
//
//
//
// Follow up:
//
//     Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
//     In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func gameOfLife(board [][]int) {
	w, h := len(board[0]), len(board)

	for i := range board {
		for j := range board[0] {
			var dead, live int

			for _, dir := range dirs {
				newY, newX := i+dir[0], j+dir[1]

				if newX >= 0 && newY >= 0 && newX < w && newY < h {
					if isDead(board[newY][newX]) {
						dead++
					} else {
						live++
					}
				}
			}

			board[i][j] = nextState(board[i][j], live, dead)
		}
	}

	// reset to original state
	for i := range board {
		for j := range board[0] {
			if board[i][j] > 1 {
				board[i][j] = (board[i][j] - 1) & 1
			}
		}
	}
}

// 0: dead -> dead
// 1: live -> live
// 2: dead -> live
// 3: live -> dead
func isDead(val int) bool {
	return val&1 == 0
}

func nextState(current, live, dead int) int {
	if current == 0 {
		if live == 3 {
			return 2
		}
		return 0
	}

	if live == 2 || live == 3 {
		return 1
	}

	return 3
}

func gameOfLife2(board [][]int) {
	w, h := len(board[0]), len(board)
	row := make([]int, w)

	var lives, topLeft int

	for i := range board {
		for j := range board[0] {
			lives = 0

			if row[j] == 1 {
				lives++
			}

			if j > 0 {
				if topLeft == 1 {
					lives++
				}

				if row[j-1] == 1 {
					lives++
				}

				if i < h-1 && board[i+1][j-1] == 1 {
					lives++
				}
			}

			topLeft = row[j]
			row[j] = board[i][j]

			if j < w-1 {
				if row[j+1] == 1 {
					lives++
				}

				if board[i][j+1] == 1 {
					lives++
				}

				if i < h-1 && board[i+1][j+1] == 1 {
					lives++
				}
			}

			if i < h-1 && board[i+1][j] == 1 {
				lives++
			}

			row[j] = board[i][j]

			if board[i][j] == 0 {
				if lives == 3 {
					board[i][j] = 1
				}
			} else {
				if lives < 2 {
					board[i][j] = 0
				} else if lives > 3 {
					board[i][j] = 0
				}
			}
		}
	}
}

func gameOfLife1(board [][]int) {
	w, h := len(board[0]), len(board)
	row := make([]int, w)

	var lives, topLeft int

	for i := range board {
		for j := range board[0] {
			lives = 0

			if row[j] == 1 {
				lives++
			}

			if j > 0 {
				if topLeft == 1 {
					lives++
				}

				if row[j-1] == 1 {
					lives++
				}

				if i < h-1 && board[i+1][j-1] == 1 {
					lives++
				}
			}

			topLeft = row[j]
			row[j] = board[i][j]

			if j < w-1 {
				if row[j+1] == 1 {
					lives++
				}

				if board[i][j+1] == 1 {
					lives++
				}

				if i < h-1 && board[i+1][j+1] == 1 {
					lives++
				}
			}

			if i < h-1 && board[i+1][j] == 1 {
				lives++
			}

			row[j] = board[i][j]

			if board[i][j] == 0 {
				if lives == 3 {
					board[i][j] = 1
				}
			} else {
				if lives < 2 {
					board[i][j] = 0
				} else if lives > 3 {
					board[i][j] = 0
				}
			}
		}
	}
}

//	Notes
//	1.	take some time to figure out what variables need to store, it takes
//		additional O(n) space to store previous row

//	2.	inspired by solution, encode state by other values, demonstrate
//		live -> live/dead, dead -> live/dead, thus, it takes O(1) space

//	3.	becareful about checking, use previous state to check

//	4.	inspired from solution, it has a great model that if a it's infinite
//		board, assume each row stores in a file in each line
