package main

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
//
// Given an integer n, return all distinct solutions to the n-queens puzzle.
//
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
//
//
// Example 1:
//
// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
//
// Example 2:
//
// Input: n = 1
// Output: [["Q"]]
//
//
//
// Constraints:
//
//     1 <= n <= 9

func solveNQueens(n int) [][]string {
	var col uint16
	ans := make([][]string, 0)
	// offset, x - y ranges from -n ~ n, convert it to 0 ~ 2n
	primary, secondary := make([]bool, 2*n), make([]bool, 2*n)

	backtracking(n, 0, []string{}, col, primary, secondary, &ans)

	return ans
}

func backtracking(n, row int, cur []string, col uint16, primary, secondary []bool, ans *[][]string) {
	if len(cur) == n {
		*ans = append(*ans, cur)
		return
	}

	for i := 0; i < n; i++ {
		// already used column, becareful, i == 0 means left most position
		offset := n - 1 - i
		if col&(1<<offset) > 0 {
			continue
		}

		// check diagonal
		// primary (\): x - y = constant
		// secondary (/): x + y = constant
		if primary[i-row+n] || secondary[i+row] {
			continue
		}

		primary[i-row+n] = true
		secondary[i+row] = true
		col |= 1 << offset

		str := make([]byte, n)
		for j := range str {
			str[j] = byte('.')
		}
		str[i] = byte('Q')

		tmp := make([]string, len(cur)+1)
		copy(tmp, cur)
		tmp[len(cur)] = string(str)

		backtracking(n, row+1, tmp, col, primary, secondary, ans)

		col = col ^ (1 << offset)
		primary[i-row+n] = false
		secondary[i+row] = false
	}
}

// each row contains only one queen, so not need to start from 0, just go to
// next row to find any possible solution
func recursive(n, y int, board [][]bool, row, col []bool, ans *[][]string) {
	size := len(board)

	if n == 0 {
		tmp := make([]string, size)

		for i := range board {
			str := make([]byte, size)

			for j := range str {
				if !board[i][j] {
					str[j] = byte('.')
				} else {
					str[j] = byte('Q')
				}
			}

			tmp[i] = string(str)
		}

		*ans = append(*ans, tmp)
		return
	}

	for i := y; i < size; i++ {
		if row[i] {
			continue
		}
		row[i] = true

		for j := 0; j < size; j++ {
			if col[j] {
				continue
			}

			// check diagonal
			diagonal := true

			tmp := min(i, j)
			for k, l := i-tmp, j-tmp; k < size && l < size; k, l = k+1, l+1 {
				if board[k][l] {
					diagonal = false
					break
				}
			}

			if !diagonal {
				continue
			}

			tmp = min(size-1-i, j)
			for k, l := i+tmp, j-tmp; k >= 0 && l < size; k, l = k-1, l+1 {
				if board[k][l] {
					diagonal = false
					break
				}
			}

			if !diagonal {
				continue
			}

			col[j] = true
			board[i][j] = true

			recursive(n-1, i+1, board, row, col, ans)

			board[i][j] = false
			col[j] = false
		}

		row[i] = false
	}
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	takes long time to debug, because I decrease n by 1 in each recursive
//		call, so it's no longer possible to use n as a checking criteria

//	2.	inspired from solution, since at most 9 rows/columns, use a single
//		unsigned integer to denote if a row/column is already used

//	3.	inspired form solution, \ (primary) & / (secondary) diagonals with
//		some property:

//		\ (primary):  x - y = constant
//		/ (secondary): x + y = constant

//	4.	since find queen's position row by row, it's no need to check rows
//		afterwards
