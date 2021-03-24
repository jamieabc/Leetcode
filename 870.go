package main

import "sort"

// Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].
//
// Return any permutation of A that maximizes its advantage with respect to B.
//
//
//
// Example 1:
//
// Input: A = [2,7,11,15], B = [1,10,4,11]
// Output: [2,11,7,15]
//
// Example 2:
//
// Input: A = [12,24,8,32], B = [13,25,32,11]
// Output: [24,32,8,12]
//
//
//
// Note:
//
// 1 <= A.length = B.length <= 10000
// 0 <= A[i] <= 10^9
// 0 <= B[i] <= 10^9

func advantageCount(A []int, B []int) []int {
	size := len(A)
	sortedA := make([]int, size)
	sortedB := make([]int, size)

	for i := range sortedA {
		sortedA[i] = i
	}
	sort.Slice(sortedA, func(i, j int) bool {
		return A[sortedA[i]] < A[sortedA[j]]
	})

	for i := range sortedB {
		sortedB[i] = i
	}
	sort.Slice(sortedB, func(i, j int) bool {
		return B[sortedB[i]] < B[sortedB[j]]
	})

	ans := make([]int, size)

	var i, j, last int
	for last = size - 1; i < len(A); i++ {
		if A[sortedA[i]] > B[sortedB[j]] {
			ans[sortedB[j]] = A[sortedA[i]]
			j++
		} else {
			ans[sortedB[last]] = A[sortedA[i]]
			last--
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/advantage-shuffle/discuss/190310/Explain-the-Algorithm-and-concise-Java-Implementation

//		author said something about generate Tian's horse racing...really
//		interesting, I have never thought of that...

//		also, max heap can be used to solve this problem
