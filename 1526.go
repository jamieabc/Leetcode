package main

import "math"

// Given an array of positive integers target and an array initial of same size with all zeros.
//
// Return the minimum number of operations to form a target array from initial if you are allowed to do the following operation:
//
// Choose any subarray from initial and increment each value by one.
//
// The answer is guaranteed to fit within the range of a 32-bit signed integer.
//
//
//
// Example 1:
//
// Input: target = [1,2,3,2,1]
// Output: 3
// Explanation: We need at least 3 operations to form the target array from the initial array.
// [0,0,0,0,0] increment 1 from index 0 to 4 (inclusive).
// [1,1,1,1,1] increment 1 from index 1 to 3 (inclusive).
// [1,2,2,2,1] increment 1 at index 2.
// [1,2,3,2,1] target array is formed.
//
// Example 2:
//
// Input: target = [3,1,1,2]
// Output: 4
// Explanation: (initial)[0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2] (target).
//
// Example 3:
//
// Input: target = [3,1,5,4,2]
// Output: 7
// Explanation: (initial)[0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1]
// -> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2] -> [3,1,5,4,2] (target).
//
// Example 4:
//
// Input: target = [1,1,1,1]
// Output: 1
//
//
//
// Constraints:
//
// 1 <= target.length <= 10^5
// 1 <= target[i] <= 10^5

func minNumberOperations(target []int) int {
	operations := target[0]
	prev := target[0]

	for i := 1; i < len(target); i++ {
		if target[i] > prev {
			operations += target[i] - prev
		}
		prev = target[i]
	}

	return operations
}

// TLE
func minNumberOperations1(target []int) int {
	var count int
	var i, j, lowest int
	changes := true

	for changes {
		changes = false
		lowest = math.MaxInt32

		for i, j = 0, 0; j < len(target); {
			// find a region separated by 0
			lowest = lowestHeight(target, &i, &j)

			if lowest != math.MaxInt32 {
				changes = true
				count += lowest
			}

			// update height
			for k := i; k < j; k++ {
				target[k] -= lowest
			}

			// update index
			i = j
		}
	}

	return count
}

func lowestHeight(target []int, i, j *int) int {
	lowest := math.MaxInt32
	size := len(target)

	for ; *i < size && target[*i] == 0; *i++ {
	}

	for *j = *i; *j < size && target[*j] > 0; *j++ {
		if lowest > target[*j] {
			lowest = target[*j]
		}
	}

	return lowest
}

//	Notes
//	1.	each interval will be decreased by lowest height, if there are n numbers,
//		worst case will be O(n^2)

//	2.	it numbers in non-increasing or non-decreasing, the cost is determined
//		by largest one

//		e.g. 1 2 3 4 5
//		for this interval, cost = 5

//		e.g. 5 4 3 2 1
//		for this interval, cost = 5

//		but how to deal with combined range?
//		e.g. 5, 2, 1, 3, 4

//	3.	inspired from https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/discuss/754623/Detailed-Explanation

//		my thinking didn't find the nature of problem, although I have thought
//		by diagram, it could be further optimized
