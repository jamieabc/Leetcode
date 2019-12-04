package main

//Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.
//
//
//Example 1:
//
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
//
//Example 2:
//
//Input: nums = [-1,0,3,5,9,12], target = 2
//Output: -1
//Explanation: 2 does not exist in nums so return -1
//
//
//
//Note:
//
//    You may assume that all elements in nums are unique.
//    n will be in the range [1, 10000].
//    The value of each element in nums will be in the range [-9999, 9999].

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	if start == end && nums[start] != target {
		return -1
	}

	middle := (start + end) / 2

	if nums[middle] == target {
		return middle
	}

	if nums[middle] > target {
		return binarySearch(nums, start, middle-1, target)
	} else {
		return binarySearch(nums, middle+1, end, target)
	}
}
