package main

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
//
// Example 2:
//
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
//
// Example 3:
//
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
//
//
//
// Constraints:
//
//     1 <= nums.length <= 105
//     -231 <= nums[i] <= 231 - 1
//
//
// Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

// tc :O(n log(n)) sc: O(1)
func increasingTriplet(nums []int) bool {
	smallest, largest := -1, -1

	// LIS patience sort
	for i, n := range nums {
		if smallest == -1 || n <= nums[smallest] {
			smallest = i
		} else if n > nums[smallest] && (largest == -1 || n <= nums[largest]) {
			largest = i
		} else if largest > -1 && n > nums[largest] {
			return true
		}
	}

	return false
}

// tc: O(n), sc: O(n)
func increasingTriplet2(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}

	smaller := make([]bool, size)
	for i, smallest := 1, nums[0]; i < size; i++ {
		if smallest < nums[i] {
			smaller[i] = true
		} else {
			smallest = nums[i]
		}
	}

	for i, largest := size-2, nums[size-1]; i >= 0; i-- {
		if nums[i] < largest {
			if smaller[i] {
				return true
			}
		} else {
			largest = nums[i]
		}
	}

	return false
}

// tc: O(n^2)
func increasingTriplet1(nums []int) bool {
	size := len(nums)
	var smaller, larger bool

	for i := range nums {
		smaller, larger = false, false

		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				smaller = true
				break
			}
		}

		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] {
				larger = true
				break
			}
		}

		if smaller && larger {
			return true
		}
	}

	return false
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/increasing-triplet-subsequence/discuss/79053/My-way-to-approach-such-a-problem.-How-to-think-about-it-Explanation-of-my-think-flow.

//		author provides a very good way of viewing this problem: LIS with
//		patience sort
