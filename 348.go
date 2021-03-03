package main

// Design a Tic-tac-toe game that is played between two players on a n x n grid.
//
// You may assume the following rules:
//
// A move is guaranteed to be valid and is placed on an empty block.
// Once a winning condition is reached, no more moves is allowed.
// A player who succeeds in placing n of their marks in a horizontal, vertical, or diagonal row wins the game.
// Example:
//
// Given n = 3, assume that player 1 is "X" and player 2 is "O" in the board.
//
// TicTacToe toe = new TicTacToe(3);
//
// toe.move(0, 0, 1); -> Returns 0 (no one wins)
// |X| | |
// | | | |    // Player 1 makes a move at (0, 0).
// | | | |
//
// toe.move(0, 2, 2); -> Returns 0 (no one wins)
// |X| |O|
// | | | |    // Player 2 makes a move at (0, 2).
// | | | |
//
// toe.move(2, 2, 1); -> Returns 0 (no one wins)
// |X| |O|
// | | | |    // Player 1 makes a move at (2, 2).
// | | |X|
//
// toe.move(1, 1, 2); -> Returns 0 (no one wins)
// |X| |O|
// | |O| |    // Player 2 makes a move at (1, 1).
// | | |X|
//
// toe.move(2, 0, 1); -> Returns 0 (no one wins)
// |X| |O|
// | |O| |    // Player 1 makes a move at (2, 0).
// |X| |X|
//
// toe.move(1, 0, 2); -> Returns 0 (no one wins)
// |X| |O|
// |O|O| |    // Player 2 makes a move at (1, 0).
// |X| |X|
//
// toe.move(2, 1, 1); -> Returns 1 (player 1 wins)
// |X| |O|
// |O|O| |    // Player 1 makes a move at (2, 1).
// |X|X|X|
// Follow up:
// Could you do better than O(n2) per move() operation?

type TicTacToe struct {
	Row, Col, Diagonal []int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{
		Row:      make([]int, n),
		Col:      make([]int, n),
		Diagonal: make([]int, 2),
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	n := len(this.Row)

	var step int
	if player == 1 {
		step = 1
	} else {
		step = -1
	}

	this.Row[row] += step
	this.Col[col] += step

	if row == col {
		this.Diagonal[0] += step
	}

	if col+row == n-1 {
		this.Diagonal[1] += step
	}

	if this.Row[row] == n || this.Row[row] == -n || this.Col[col] == n || this.Col[col] == -n || this.Diagonal[0] == n || this.Diagonal[0] == -n || this.Diagonal[1] == n || this.Diagonal[1] == -n {
		return player
	}

	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */

type TicTacToe1 struct {
	row      []int
	column   []int
	diagonal []int
	fail     int
}

/** Initialize your data structure here. */
func Constructor1(n int) TicTacToe1 {
	return TicTacToe1{
		row:      make([]int, n),
		column:   make([]int, n),
		diagonal: make([]int, 2),
		fail:     n + 1,
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe1) Move(row int, col int, player int) int {
	r := this.row[row]
	c := this.column[col]

	step := 1
	if player == 2 {
		step = -1
	}
	var positive bool
	if player == 1 {
		positive = true
	}

	if r != this.fail {
		if r == 0 || positive == (r > 0) {
			this.row[row] += step
		} else {
			this.row[row] = this.fail
		}
	}

	if c != this.fail {
		if c == 0 || positive == (c > 0) {
			this.column[col] += step
		} else {
			this.column[col] = this.fail
		}
	}

	d1 := this.diagonal[0]
	d2 := this.diagonal[1]

	if row == col {
		if d1 != this.fail {
			if d1 == 0 || positive == (d1 > 0) {
				this.diagonal[0] += step
			} else {
				this.diagonal[0] = this.fail
			}
		}
	}

	if col == len(this.row)-1-row {
		if d2 != this.fail {
			if d2 == 0 || positive == (d2 > 0) {
				this.diagonal[1] += step
			} else {
				this.diagonal[1] = this.fail
			}
		}
	}

	return this.check()
}

// 0 - no one win
// 1 - player 1 wins
// 2 -player 2 wins
func (this *TicTacToe1) check() int {
	for _, i := range this.row {
		if abs(i) == this.fail-1 {
			return winner(i)
		}
	}

	for _, i := range this.column {
		if abs(i) == this.fail-1 {
			return winner(i)
		}
	}

	for _, i := range this.diagonal {
		if abs(i) == this.fail-1 {
			return winner(i)
		}
	}

	return 0
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func winner(result int) int {
	if result > 0 {
		return 1
	}
	return 2
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */

//	problems
//	1.	when merging conditions for both players, it needs to separete statement
//		of ==0, >0, <0

//	2.	wrong logic, when player put at middle of board, it could be both
//		diagonal & reverse diagonal

//	3.	add reference https://leetcode.com/problems/design-tic-tac-toe/discuss/343824/Python-O(1)-time-O(n)-space.-Detailed-explanation

//		original I think tc is O(n), but it's actually O(1), because it only
//		relates to array accessing which is constant time, not relates to how
//		big board is

//		sc is O(n)
