package main

// A sequence of number is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.
//
// For example, these are arithmetic sequence:
//
// 1, 3, 5, 7, 9
// 7, 7, 7, 7
// 3, -1, -5, -9
//
// The following sequence is not arithmetic.
//
// 1, 1, 2, 5, 7
//
//
// A zero-indexed array A consisting of N numbers is given. A slice of that array is any pair of integers (P, Q) such that 0 <= P < Q < N.
//
// A slice (P, Q) of array A is called arithmetic if the sequence:
// A[P], A[p + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.
//
// The function should return the number of arithmetic slices in the array A.
//
// Example:
//
// A = [1, 2, 3, 4]
//
// return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.

func numberOfArithmeticSlices(A []int) int {
	var total int
	length := len(A)

	if length <= 2 {
		return total
	}

	for dp, i := 0, 2; i < length; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp++
			total += dp
		} else {
			dp = 0
		}
	}

	return total
}

func numberOfArithmeticSlices2(A []int) int {
	size := len(A)

	// dp[i]: # of ways for sequence ends at i
	dp := make([]int, size)
	var count int

	for i := 2; i < size; i++ {
		if A[i-2]-A[i-1] == A[i-1]-A[i] {
			dp[i] = dp[i-1] + 1
			count += dp[i]
		} else {
			dp[i] = 0
		}
	}

	return count
}

// tc: O(n)
func numberOfArithmeticSlices1(A []int) int {
	var count int
	size := len(A)
	if size < 3 {
		return 0
	}

	for low, high := 0, 0; low < size; {
		if high-low < 2 {
			high = low + 2
		}

		for ; high < size && A[low]-A[low+1] == A[high-1]-A[high]; high++ {
		}

		count += max(0, high-low-2)

		low++
	}

	return count
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	not considering boundary cases, such as array is empty, because
//		maxDiff relies on length - 1

//	2.	referenced from https://leetcode.com/problems/arithmetic-slices/discuss/90095/why-the-result-of-123456-is-10-not-12

//		the definition of slice is consecutive elements, for example of
//		1, 2, 3, 4, 5, ,6
//		although 2, 4, 6 are arithmetic numbers, but those numbers are not
//		consecutive in original array, so it's not a slice

//	3.	inspired from https://leetcode.com/problems/arithmetic-slices/discuss/90058/Simple-Java-solution-9-lines-2ms

//		author reduces dp, if differences of consecutive numbers are same,
//		increment dp. I focus on the situation that when differences exist,
//		reset, but this rule is already valid when checking consecutive
//		numbers

//	4.	inspired from https://leetcode.com/problems/arithmetic-slices/discuss/90093/3ms-C%2B%2B-Standard-DP-Solution-with-Very-Detailed-Explanation

//		array length must be larger or equals 3, otherwise will be 0
