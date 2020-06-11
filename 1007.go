package main

import "math"

//In a row of dominoes, A[i] and B[i] represent the top and bottom halves of the i-th domino.  (A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)
//
//We may rotate the i-th domino, so that A[i] and B[i] swap values.
//
//Return the minimum number of rotations so that all the values in A are the same, or all the values in B are the same.
//
//If it cannot be done, return -1.
//
//
//
//Example 1:
//
//Input: A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
//Output: 2
//Explanation:
//The first figure represents the dominoes as given by A and B: before we do any rotations.
//If we rotate the second and fourth dominoes, we can make every value in the top row equal to 2, as indicated by the second figure.
//
//Example 2:
//
//Input: A = [3,5,1,2,3], B = [3,6,3,3,4]
//Output: -1
//Explanation:
//In this case, it is not possible to rotate the dominoes to make one row of values equal.
//
//
//
//Note:
//
//    1 <= A[i], B[i] <= 6
//    2 <= A.length == B.length <= 20000

func minDominoRotations(A []int, B []int) int {
	// if there's a solution, choose A[0] or B[0] are identical
	a := rotate(A, B, A[0])
	if a != -1 {
		return a
	}

	// if only condition for B[0]
	return rotate(A, B, B[0])
}

func rotate(A, B []int, target int) int {
	var a, b int
	for i := range A {
		if A[i] != target && B[i] != target {
			return -1
		} else if A[i] != target && B[i] == target {
			a++
		} else if B[i] != target && A[i] == target {
			b++
		}
	}

	return min(a, b)
}

func minDominoRotations2(A []int, B []int) int {
	a, b := A[0], B[0]
	dp := make([]int, 4) // a at A, a at B, b at A, b at B

	for i := range A {
		if A[i] != a && B[i] != a {
			dp[0] = -1
			dp[1] = -1
		} else if A[i] != a {
			if B[i] == a && dp[0] >= 0 {
				dp[0]++
			} else {
				dp[0] = -1
			}
		} else if B[i] != a {
			if A[i] == a && dp[1] >= 0 {
				dp[1]++
			} else {
				dp[1] = -1
			}
		}

		if A[i] != b && B[i] != b {
			dp[2] = -1
			dp[3] = -1
		} else if A[i] != b {
			if B[i] == b && dp[2] >= 0 {
				dp[2]++
			} else {
				dp[2] = -1
			}
		} else if B[i] != b {
			if A[i] == b && dp[3] >= 0 {
				dp[3]++
			} else {
				dp[3] = -1
			}
		}
	}

	m := math.MaxInt32
	for i := range dp {
		if dp[i] >= 0 {
			m = min(m, dp[i])
		}
	}

	if m == math.MaxInt32 {
		return -1
	}

	return m
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func minDominoRotations1(A []int, B []int) int {
	lengthA := len(A)

	if lengthA <= 1 {
		return 0
	}

	minCount := -1                // default find nothing
	toCheck := [2]int{A[0], B[0]} // number is aligned with either A[0] or B[0] because swap of first element comes from these two values

	for _, num := range toCheck {
		countA := findRotation(A, B, num)
		countB := findRotation(B, A, num)

		tmp := positiveMin(countA, countB)
		if minCount == -1 {
			minCount = tmp
		} else {
			if tmp > -1 && tmp < minCount {
				minCount = tmp
			}
		}
	}
	return minCount
}

func positiveMin(i, j int) int {
	// both not found
	if i == -1 && j == -1 {
		return -1
	}

	// fll found, find min
	if i > 0 && j > 0 {
		if i <= j {
			return i
		}
		return j
	}

	// one found, find larger one
	if i >= j {
		return i
	}
	return j
}

func findRotation(target, backup []int, num int) int {
	count := 0
	length := len(target)

	var j int
	for j = 0; j < length; j++ {
		if target[j] != num {
			if backup[j] == num {
				count++
			} else {
				break
			}
		}
	}

	// find a solution
	if j == length {
		return count
	}

	return -1
}

// problems
// 1. either A or B can do, no as what I think just A
// 2. when length is 1, no need to compare
// 3. optimization, since it's either all A or B are same, possible number is either A[0] or B[0], this can reduce complexity
// 4. length of A & B are same, because each means half side value of dominos

//	5.	if every number are same, no need to to anything

//	6.	inspired from solution, if 2 rows can both fits condition, then
//		flipping numbers for 2 conditions (A[0] or B[0]) are same inverted.

//		e.g. A: 2, 5, 5, 5, 2  2: 3, 5: 2
//			 B: 5, 2, 2, 2, 5  5: 2, 2: 3

//		however, there still exists condition that A cannot be same with
//		A[0], e.g.

//		A: 2, 5, 5, 2, 2
//		B: 5, 1, 2, 5, 5

//	7.	inspired from https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/discuss/252242/JavaC%2B%2BPython-Different-Ideas

//		lee with 3 solutions....how come so many!
