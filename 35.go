package main

//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Example 1:
//
//Input: [1,3,5,6], 5
//Output: 2
//
//Example 2:
//
//Input: [1,3,5,6], 2
//Output: 1
//
//Example 3:
//
//Input: [1,3,5,6], 7
//Output: 4
//
//Example 4:
//
//Input: [1,3,5,6], 0
//Output: 0

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, start, end int) int {
	// only 2 elements
	if end < 0 {
		return 0
	}

	if start == len(nums) {
		return start
	}

	if start > end {
		if nums[start] < target {
			return start + 1
		}
		if nums[end] < target {
			return end + 1
		}
		return start
	}

	if start == end {
		if nums[start] < target {
			return start + 1
		} else {
			return start
		}
	}

	middle := (start + end) / 2

	if nums[middle] == target {
		return middle
	}

	if nums[middle] < target {
		return binarySearch(nums, target, middle+1, end)
	} else {
		return binarySearch(nums, target, start, middle-1)

	}
}
