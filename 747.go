package main

//In a given integer array nums, there is always exactly one largest element.
//
//Find whether the largest element in the array is at least twice as much as every other number in the array.
//
//If it is, return the index of the largest element, otherwise return -1.
//
//Example 1:
//
//Input: nums = [3, 6, 1, 0]
//Output: 1
//Explanation: 6 is the largest integer, and for every other number in the array x,
//6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
//
//
//
//Example 2:
//
//Input: nums = [1, 2, 3, 4]
//Output: -1
//Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
//
//
//
//Note:
//
//    nums will have a length in the range [1, 50].
//    Every nums[i] will be an integer in the range [0, 99].

func dominantIndex(nums []int) int {
	length := len(nums)

	if length == 0 {
		return -1
	}

	if length == 1 {
		return 0
	}

	if length == 2 {
		if nums[0]*2 <= nums[1] {
			return 1
		}
		if nums[1]*2 <= nums[0] {
			return 0
		}
		return -1
	}

	max := 0
	max2 := 1
	if nums[0] < nums[1] {
		max = 1
		max2 = 0
	}

	for i := 2; i < length; i++ {
		if nums[max2] > nums[i] {
			continue
		}

		if nums[i] > nums[max] {
			max2 = max
			max = i
			continue
		}

		if nums[i] < nums[max] && nums[i] > nums[max2] {
			max2 = i
			continue
		}
	}

	if nums[max] >= nums[max2]*2 {
		return max
	}

	return -1
}

// problems
// 1. forget about condition that has only 1 element
// 2. separate length of 0 & 1
