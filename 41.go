package main

// Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1:
//
// Input: [1,2,0]
// Output: 3
// Example 2:
//
// Input: [3,4,-1,1]
// Output: 2
// Example 3:
//
// Input: [7,8,9,11,12]
// Output: 1
// Note:
//
// Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] < len(nums) && i != nums[i]-1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := range nums {
		if i != nums[i]-1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

//	problems
//	1.	boundary case when array is empty
