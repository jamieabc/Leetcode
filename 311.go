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

//	problems
//	1.	Wrong index of product matrix

//	2.	Optimization, store A's row & B's column that contains non-zero
//		numbers (sparse matrix), use that information to do calculation.

//		Be ware to use actual value in row & column, not index

//	3. reference from https://leetcode.com/problems/sparse-matrix-multiplication/discuss/76151/54ms-Detailed-Summary-of-Easiest-JAVA-solutions-Beating-99.9

//		check for first array non-zero value
//		convert second array into sequence of non-zero data
//		B = | 1 0 0 1 |
//		    | 0 0 0 0 |
//		    | 1 0 0 0 |

//		=> [ [ [0, 1], [3, 1] ], [[]], [ [ [0, 1] ] ] ]
//		CMU sparse matrix http://www.cs.cmu.edu/~scandal/cacm/node9.html

//	4.	another reference https://github.com/SCIN/Facebook-Interview-Coding-1/blob/master/Sparce%20Matrix%20Multiplication.java

//		this one is more details in the interview process, use list to save
//		space, sorted index, different length of matrix, etc. Helpful
