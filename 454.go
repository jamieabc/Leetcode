package main

// Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
//
// To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.
//
// Example:
//
// Input:
// A = [ 1, 2]
// B = [-2,-1]
// C = [-1, 2]
// D = [ 0, 2]
//
// Output:
// 2
//
// Explanation:
// The two tuples are:
// 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
// 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

// tc: O(n^2), sc: O(n^2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	table := make(map[int]int)
	for i := range A {
		for j := range B {
			table[A[i]+B[j]]++
		}
	}

	var count int
	for i := range C {
		for j := range D {
			count += table[-C[i]-D[j]]
		}
	}

	return count
}

// tc: O(n^2), sc: O(n^2)
func fourSumCount1(A []int, B []int, C []int, D []int) int {
	table1, table2 := make(map[int]int), make(map[int]int)

	for i := range A {
		for j := range B {
			table1[A[i]+B[j]]++
		}
	}

	for i := range C {
		for j := range D {
			table2[C[i]+D[j]]++
		}
	}

	var ans int

	for sum1, occurrence := range table1 {
		if val, ok := table2[-sum1]; ok {
			ans += occurrence * val
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/4sum-ii/discuss/93917/Easy-2-lines-O(N2)-Python

//		no need to use another table to store combination sums

//	2.	if there are duplicates, then cannot sort C & D to find matched sum, because when
//		iterating through array, if C[i] = 2 && D[j] = 2, don't know which array to move
//		forward, i or j?
