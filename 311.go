package main

// Given two sparse matrices A and B, return the result of AB.
//
// You may assume that A's column number is equal to B's row number.
//
// Example:
//
// Input:
//
// A = [
//   [ 1, 0, 0],
//   [-1, 0, 3]
// ]
//
// B = [
//   [ 7, 0, 0 ],
//   [ 0, 0, 0 ],
//   [ 0, 0, 1 ]
// ]
//
// Output:
//
//      |  1 0 0 |   | 7 0 0 |   |  7 0 0 |
// AB = | -1 0 3 | x | 0 0 0 | = | -7 0 3 |
//                   | 0 0 1 |

func multiply(A [][]int, B [][]int) [][]int {
	y := len(A)
	if y == 0 {
		return [][]int{}
	}
	x := len(B[0])

	result := make([][]int, y)
	for i := range result {
		result[i] = make([]int, x)
	}

	for i := range A {
		for j := range A[0] {
			if A[i][j] != 0 {
				for k := range B[0] {
					if B[j][k] != 0 {
						result[i][k] += A[i][j] * B[j][k]
					}
				}
			}
		}
	}

	return result
}

//	Notes
//	1.	Wrong index of product matrix

//	2.	Optimization, store A's row & B's column that contains non-zero
//		numbers (sparse matrix), use that information to do calculation.

//		Be ware to use actual value in row & column, not index

//	3. reference from https://leetcode.com/problems/sparse-matrix-multiplication/discuss/76151/54ms-Detailed-Summary-of-Easiest-JAVA-solutions-Beating-99.9

//		check for first array non-zero value

//		space compression
//		convert second array into sequence of non-zero data (index, value)
//		B = | 1 0 0 1 |     => [ [ (0, 1), (3, 1)]  value 1 at index 0, 3 at index 1
//		    | 0 0 0 0 |			 []
//		    | 1 0 0 0 |			 [ (0, 1) ] ]       value 1 at index 0

//		CMU sparse matrix http://www.cs.cmu.edu/~scandal/cacm/node9.html

//	4.	another reference https://github.com/SCIN/Facebook-Interview-Coding-1/blob/master/Sparce%20Matrix%20Multiplication.java

//		this one is more details in the interview process, use list to save
//		space, sorted index, different length of matrix, etc. Helpful

//	5.	inspired from https://leetcode.com/problems/sparse-matrix-multiplication/discuss/419538/What-the-interviewer-is-expecting-when-this-problem-is-asked-in-an-interview...

//		use 2 pointers to compute, because cik = aij * bjk,
//		data structure: idx1: [0, 1, 3]
//					 	val1: [4, 10, 20]

//		index 0 w/ val 4, index 1 w/ val 10, index 3 w/ val 20

//		the other follow up might be one array is much longer than the other one
//		can use shorter one as base, binary search longer one, use use searched
//		index as next start point
