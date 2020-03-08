package main

//Given an 2D board, count how many battleships are in it. The battleships are represented with 'X's, empty slots are represented with '.'s. You may assume the following rules:
//You receive a valid board, made of only battleships or empty slots.
//Battleships can only be placed horizontally or vertically. In other words, they can only be made of the shape 1xN (1 row, N columns) or Nx1 (N rows, 1 column), where N can be of any size.
//At least one horizontal or vertical cell separates between two battleships - there are no adjacent battleships.
//Example:
//
//X..X
//...X
//...X
//In the above board there are 2 battleships.
//Invalid Example:
//
//...X
//XXXX
//...X
//This is an invalid board that you will not receive - as battleships will always have a cell separating between them.
//Follow up:
//Could you do it in one-pass, using only O(1) extra memory and without modifying the value of the board?

func countBattleships(board [][]byte) int {
	y := len(board)
	if y == 0 {
		return 0
	}
	x := len(board[0])
	if x == 0 {
		return 0
	}

	count := 0
	var i, j int
	for i = 0; i < y; i++ {
		for j = 0; j < x; j++ {
			if board[i][j] == '.' {
				continue
			}

			if j > 0 && board[i][j-1] == 'X' {
				continue
			}

			if i > 0 && board[i-1][j] == 'X' {
				continue
			}

			count++
		}
	}

	return count
}

// problems
//	1.	when checking for row, need to be continuous checking, e.g. j cannot
//		directory to k+1 because additional checking might need on k
//		for example ...x
//					.xx.
//	2.	optimization, the dp actually does nothing, without it, if a X appears
//		on previous line same index, then it must be a vertical one
//	3.	optimization, only count the start of ship, which means in two
//		conditions:
//		- vertical: up is not X
//		- horizontal: left is not X
