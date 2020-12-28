package main

import "sort"

// Given an array A of integers, for each integer A[i] we need to choose either x = -K or x = K, and add x to A[i] (only once).
//
// After this process, we have some array B.
//
// Return the smallest possible difference between the maximum value of B and the minimum value of B.
//
//
//
// Example 1:
//
// Input: A = [1], K = 0
// Output: 0
// Explanation: B = [1]
//
// Example 2:
//
// Input: A = [0,10], K = 2
// Output: 6
// Explanation: B = [2,8]
//
// Example 3:
//
// Input: A = [1,3,6], K = 3
// Output: 3
// Explanation: B = [4,6,3]
//
//
//
// Note:
//
//     1 <= A.length <= 10000
//     0 <= A[i] <= 10000
//     0 <= K <= 10000

func smallestRangeII(A []int, K int) int {
	size := len(A)
	sort.Ints(A)

	if size <= 1 || A[size-1] == A[0] {
		return 0
	}

	ans := A[size-1] - A[0]

	for i := 1; i < size; i++ {
		ans = min(ans, max(A[size-1]-K, A[i-1]+K)-min(A[0]+K, A[i]-K))
	}

	return ans
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	consider final result, care only about max/min after transformation,
//		numbers within range are not considered

//	2.	For max/min in original array, it matters on average number & K.
//		To get minimum difference of transformed max & min, the strategy is add
//		K to all numbers < average, and minus K to all numbers > average

//		average = (max+min) / 2
//		min							average								max

//		(1) min+K < average-K
//		min		min+K   average-K	average		average+K	max-K		max

//		after transformation, smallest difference = max-K-(min+K)

//		(2) min < average-K < min+K
//		min		average-K	min+K	average		max-K		average+K	max

//		(3) average-K < min < max-K
//		average-K	min		max-K	average		min+K		max			average+K

//		after transformation, smallest difference either on average+K-(
//		number_just_larger_than_average-K) or
//		number_just_smaller_than_average+K-(average-K)

//	3.	it's not always true to have numbers +K before average, and -K after
//		average

//	4.	inspired from https://leetcode.com/problems/smallest-range-ii/discuss/173495/Actual-explanation-for-people-who-don't-understand-(I-hope)

//		it's complicated to consider relationships between K & average, another
//		way to view this problem is to know that max-K & min+K exists, so any
//		number breaks this range will become new max/min value in transformed
//		array

//	5.	inspired from https://leetcode.com/problems/smallest-range-ii/discuss/466198/5-lines-of-code-Best-Solution-with-Explanation

//		author provides a clue to prove if K >= (max - min), then do nothing.

//		K >= D (max - min)
//		max-K	min .... max min+K

//		new distance = (min+K) - (max-K) = min - max + 2K = -D + 2K >= D

//		if K < D, then there will be some point that shrink distance
//		+K, +K, +K, ..., +K, -K, -K, -K, ..., -K
