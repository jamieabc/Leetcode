package main

// We have a two dimensional matrix A where each value is 0 or 1.
//
// A move consists of choosing any row or column, and toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s.
//
// After making any number of moves, every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.
//
// Return the highest possible score.
//
//
//
// Example 1:
//
// Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
// Output: 39
// Explanation:
// Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
// 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
//
//
//
// Note:
//
//     1 <= A.length <= 20
//     1 <= A[0].length <= 20
//     A[i][j] is 0 or 1.

func matrixScore(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	// make sure first digit of every row is 1
	for i := range A {
		if A[i][0] == 0 {
			flipRow(i, A)
		}
	}

	// for every other column, make sure it flips only if 0s more than 1s
	for i := 1; i < len(A[0]); i++ {
		if columnWithMore0(i, A) {
			flipColumn(i, A)
		}
	}

	var sum int
	for i := range A {
		sum += toNumber(A[i])
	}
	return sum
}

func toNumber(arr []int) int {
	var total int
	for _, i := range arr {
		total *= 2

		if i == 1 {
			total++
		}
	}
	return total
}

func columnWithMore0(index int, A [][]int) bool {
	var oneCount int
	for i := range A {
		if A[i][index] == 1 {
			oneCount++
		}
	}

	return len(A)-oneCount > oneCount
}

func flipRow(index int, A [][]int) {
	for i := range A[index] {
		if A[index][i] == 1 {
			A[index][i] = 0
		} else {
			A[index][i] = 1
		}
	}
}

func flipColumn(index int, A [][]int) {
	for i := range A {
		if A[i][index] == 1 {
			A[i][index] = 0
		} else {
			A[i][index] = 1
		}
	}
}
