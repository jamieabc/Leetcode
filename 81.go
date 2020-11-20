package main

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
//
// You are given a target value to search. If found in the array return true, otherwise return false.
//
// Example 1:
//
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
//
// Example 2:
//
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false
//
// Follow up:
//
//     This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
//     Would this affect the run-time complexity? How and why?

func search(nums []int, target int) bool {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) bool {
	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == target {
			return true
		}

		// if left side boundary same as right side boundary, don't know which
		// side to go, need to advance one side to check
		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low++
			high--
		} else if nums[mid] >= nums[low] {
			// equal is crucial, if low == mid, it's still valid that left side
			// is ascending
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}

//	Notes
//	1.	duplicate numbers exist, it might cause binary separation not able to
//		check, e.g. nums = [1, 3, 1, 1, 1, 1], nums[0] = nums[2] = nums[5],
//		there's no way to check which side to go

//	2.	inspired from solution, when left & right & mid with same values, a
//		better way to do it is to advance one side and check again
