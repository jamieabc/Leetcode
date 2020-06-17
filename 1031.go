package main

import "math"

// Given an array A of non-negative integers, return the maximum sum of elements in two non-overlapping (contiguous) subarrays, which have lengths L and M.  (For clarification, the L-length subarray could occur before or after the M-length subarray.)
//
// Formally, return the largest V for which V = (A[i] + A[i+1] + ... + A[i+L-1]) + (A[j] + A[j+1] + ... + A[j+M-1]) and either:
//
//     0 <= i < i + L - 1 < j < j + M - 1 < A.length, or
//     0 <= j < j + M - 1 < i < i + L - 1 < A.length.
//
//
//
// Example 1:
//
// Input: A = [0,6,5,2,2,5,1,9,4], L = 1, M = 2
// Output: 20
// Explanation: One choice of subarrays is [9] with length 1, and [6,5] with length 2.
//
// Example 2:
//
// Input: A = [3,8,1,3,2,1,8,9,0], L = 3, M = 2
// Output: 29
// Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
//
// Example 3:
//
// Input: A = [2,1,5,6,0,9,5,0,3,8], L = 4, M = 3
// Output: 31
// Explanation: One choice of subarrays is [5,6,0,9] with length 4, and [3,8] with length 3.
//
//
//
// Note:
//
//     L >= 1
//     M >= 1
//     L + M <= A.length <= 1000
//     0 <= A[i] <= 1000

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	size := len(A)
	if size == 0 {
		return 0
	}

	sums := make([]int, size)
	sums[0] = A[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + A[i]
	}

	maxSum := sums[L+M-1]
	maxL, maxM := sums[L-1], sums[M-1]

	// maxL then M
	for i := L + M; i < size; i++ {
		maxL = max(maxL, sums[i-M]-sums[i-M-L])
		maxSum = max(maxSum, maxL+sums[i]-sums[i-M])
	}

	// maxM then L
	for i := L + M; i < size; i++ {
		maxM = max(maxM, sums[i-L]-sums[i-M-L])
		maxSum = max(maxSum, maxM+sums[i]-sums[i-L])
	}

	return maxSum
}

func maxSumTwoNoOverlap1(A []int, L int, M int) int {
	size := len(A)
	if size == 0 {
		return 0
	}

	// running sum
	sums := make([]int, size)
	sums[0] = A[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + A[i]
	}

	longer, shorter := L, M
	if L < M {
		longer, shorter = M, L
	}

	// max sum after ith index
	maxSumAfter := make([]int, size)
	for i, sum := size-1, 0; i >= 0; i-- {
		if i >= size-shorter {
			sum += A[i]
			maxSumAfter[i] = sum
		} else {
			sum = sum - A[i+shorter] + A[i]
			if sum >= maxSumAfter[i+1] {
				maxSumAfter[i] = sum
			} else {
				maxSumAfter[i] = maxSumAfter[i+1]
			}
		}
	}

	// max sum before ith index
	maxSumBefore := make([]int, size)
	for i, sum := 0, 0; i < size; i++ {
		if i < shorter {
			sum += A[i]
			maxSumBefore[i] = sum
		} else {
			sum = sum - A[i-shorter] + A[i]
			if sum >= maxSumBefore[i-1] {
				maxSumBefore[i] = sum
			} else {
				maxSumBefore[i] = maxSumBefore[i-1]
			}
		}
	}

	maxSum := math.MinInt32

	var longerSum int
	for i := 0; i <= size-longer; i++ {
		if i == 0 {
			longerSum = sums[i+longer-1]
		} else {
			longerSum = sums[i+longer-1] - sums[i-1]
		}

		if i == 0 {
			maxSum = max(maxSum, longerSum+maxSumAfter[i+longer])
		} else if i == size-longer {
			maxSum = max(maxSum, longerSum+maxSumBefore[i-1])
		} else {
			maxSum = max(maxSum, longerSum+max(maxSumAfter[i+longer], maxSumBefore[i-1]))
		}
	}

	return maxSum
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	wrong range for max sum before

//	2.	miss a condition, if i == size - longer, maxSum only comes from
//		before

//	3.	tc: O(n), this problem is testing dp idea of storing sum

//	4.	inspired from https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/discuss/300029/Python-breaking-down-Lee215's-solution

//		max sum happens in 2 conditions:
//		- at ith position, sum prev M numbers, and choose maxL before i-M

//		  | - - - - - - - - | - - - |
//			     maxL           M    i

//		/ at ith position, sum prev L numbers, and choose maxM before i-L

//		  | - - - - - - - - - - - - | - - - - - |
// 				     maxM                L      i
