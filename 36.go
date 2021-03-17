package main

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//     Each row must contain the digits 1-9 without repetition.
//     Each column must contain the digits 1-9 without repetition.
//     Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// Note:
//
//     A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//     Only the filled cells need to be validated according to the mentioned rules.
//
//
//
// Example 1:
//
// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
//
// Example 2:
//
// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
//
//
//
// Constraints:
//
//     board.length == 9
//     board[i].length == 9
//     board[i][j] is a digit or '.'.

func isValidSudoku(board [][]byte) bool {
	w, h := len(board[0]), len(board)

	rows, cols := make([][]bool, h), make([][]bool, w)
	for i := range rows {
		rows[i] = make([]bool, 10)
	}

	for i := range cols {
		cols[i] = make([]bool, 10)
	}

	rec := make([][]bool, w/3*h/3)
	for i := range rec {
		rec[i] = make([]bool, 10)
	}

	for i := range board {
		for j := range board[0] {
			if num := board[i][j]; num != '.' {
				idx := (i/3)*3 + j/3
				n := int(num - '0')

				if rows[i][n] || cols[j][n] || rec[idx][n] {
					return false
				}
				rows[i][n] = true
				cols[j][n] = true
				rec[idx][n] = true
			}
		}
	}

	return true
}

//	Notes
//	1.	3*3 grid number need to be unique
