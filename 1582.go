package main

//Given a rows x cols matrix mat, where mat[i][j] is either 0 or 1, return the number of special positions in mat.
//
//A position (i,j) is called special if mat[i][j] == 1 and all other elements in row i and column j are 0 (rows and columns are 0-indexed).
//
//
//
//Example 1:
//
//Input: mat = [[1,0,0],
//              [0,0,1],
//              [1,0,0]]
//Output: 1
//Explanation: (1,2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
//
//Example 2:
//
//Input: mat = [[1,0,0],
//              [0,1,0],
//              [0,0,1]]
//Output: 3
//Explanation: (0,0), (1,1) and (2,2) are special positions.
//
//Example 3:
//
//Input: mat = [[0,0,0,1],
//              [1,0,0,0],
//              [0,1,1,0],
//              [0,0,0,0]]
//Output: 2
//
//Example 4:
//
//Input: mat = [[0,0,0,0,0],
//              [1,0,0,0,0],
//              [0,1,0,0,0],
//              [0,0,1,0,0],
//              [0,0,0,1,1]]
//Output: 3
//
//
//
//Constraints:
//
//    rows == mat.length
//    cols == mat[i].length
//    1 <= rows, cols <= 100
//    mat[i][j] is 0 or 1.

func numSpecial(mat [][]int) int {
	row, col := make([]int, len(mat)), make([]int, len(mat[0]))

	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}

	var special int
	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				special++
			}
		}
	}
	return special
}

//	Notes
//	1.	at first I think it's dp, but when I use row & col to store 1s count and
//		find it's not right. simply check row[i] == 1 & col[i] == 1 is not
//		correct, since row i has only one 1 doesn't mean col i also needs to
//		have only one 1.

//	2.	inspired from https://leetcode.com/problems/special-positions-in-a-binary-matrix/discuss/843949/C%2B%2B-2-passes

//		iterate array twice to get the count, first iteration count row & col
//		based 1s number, second pass to check for every 1 meets criteria
