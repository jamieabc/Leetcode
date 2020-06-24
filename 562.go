package main

// Given a 01 matrix M, find the longest line of consecutive one in the matrix. The line could be horizontal, vertical, diagonal or anti-diagonal.
// Example:
//
// Input:
// [[0,1,1,0],
//  [0,1,1,0],
//  [0,0,0,1]]
// Output: 3
// Hint: The number of elements in the given matrix will not exceed 10,000.

func longestLine(M [][]int) int {
	if len(M) == 0 {
		return 0
	}

	vertical := make([]int, len(M[0]))
	diagonal := make([]int, len(M[0]))
	antiDiagonal := make([]int, len(M[0]))

	var maxConsecutiveOne, horizontalOne, prevDiagonal, tmpDiagonal int

	for i := range M {
		horizontalOne, prevDiagonal = 0, 0

		for j := range M[0] {
			tmpDiagonal = diagonal[j]

			if M[i][j] == 1 {
				// vertical
				vertical[j]++
				maxConsecutiveOne = max(maxConsecutiveOne, vertical[j])

				// diagonal
				if j == 0 || i == 0 {
					diagonal[j] = 1
				} else {
					diagonal[j] = prevDiagonal + 1
				}
				maxConsecutiveOne = max(maxConsecutiveOne, diagonal[j])

				// anti diagonal
				if j == len(M[0])-1 || i == 0 {
					antiDiagonal[j] = 1
				} else {
					antiDiagonal[j] = antiDiagonal[j+1] + 1
				}
				maxConsecutiveOne = max(maxConsecutiveOne, antiDiagonal[j])

				// horizontal
				horizontalOne++
				maxConsecutiveOne = max(maxConsecutiveOne, horizontalOne)
			} else {
				vertical[j] = 0
				horizontalOne = 0
				diagonal[j] = 0
				antiDiagonal[j] = 0
			}

			prevDiagonal = tmpDiagonal
		}
	}

	return maxConsecutiveOne
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
