package main

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// You are given a target value to search. If found in the array return its index, otherwise return -1.
//
// You may assume no duplicate exists in the array.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// Example 1:
//
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
//
// Example 2:
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	} else if size == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	var low, high int

	for low, high = 0, len(nums)-1; low < high; {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		// decide which directions to go
		if nums[mid] >= nums[low] {
			// 3, 4, 5, 6, 1, 2
			// ^     ^        ^
			if target > nums[mid] || target < nums[low] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// 5, 6, 1, 2, 3, 4
			// ^        ^     ^
			if target < nums[low] && target > nums[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	if nums[low] == target {
		return low
	}
	return -1
}

func search3(nums []int, target int) int {
	for low, high := 0, len(nums)-1; low <= high; {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[high] < nums[mid] {
			// smallest at right side, largest also at right side
			if nums[low] > target || nums[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// smallest at left side, largest also at left side
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// tc: O(n log(n))
func search2(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[low] < nums[high] {
				high = mid - 1
			} else {
				if tmp := binarySearch(nums, target, low, mid-1); tmp == -1 {
					return binarySearch(nums, target, mid+1, high)
				} else {
					return tmp
				}
			}
		} else {
			if nums[low] < nums[high] {
				low = mid + 1
			} else {
				if tmp := binarySearch(nums, target, low, mid-1); tmp == -1 {
					return binarySearch(nums, target, mid+1, high)
				} else {
					return tmp
				}
			}
		}
	}

	return -1
}

func search1(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	if size == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	pivot := findPivot(nums)

	if pivot != 0 && target > nums[len(nums)-1] {
		return binarySearch(nums, 0, pivot-1, target)
	}

	return binarySearch(nums, pivot, len(nums)-1, target)
}

// pivot number is the smallest number in array
func findPivot(nums []int) int {
	var low, high int
	for low, high = 0, len(nums)-1; low < high; {
		mid := low + (high-low)/2

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func binarySearch(nums []int, start, end, target int) int {
	var low, high int

	for low, high = start, end; low < high; {
		mid := low + (high-low+1)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}

	if nums[low] == target {
		return low
	}

	return -1
}

//	Notes
//	1.	in later check, always array is not empty, so need to check it in
//		advance

//	2.	intuition, the pivot index is the smallest number in array

//	3.	inspired from solution, for every mid number not equal to target, decide
//		next half segment to search. so the problem becomes how to decide where
//		to search.

//		compare to target, if nums[mid] < target means need to search larger
//		part of interval, since array is shifted, larger part can be decided by
//		comparing nums[mid] to nums[start], if nums[mid] is larger, then search
//		right side of mid, if nums[mid] is smaller, than search left side of
//		mid

//	4.	the point is about pivot number (smallest number), the largest & smallest
//		number might be in the same side, so it needs two conditions to check

//		for smallest & largest at same side, check need to confirm these two
//		conditions at least one meet, so that binary search can go to the side

//		3 4 5 6 1 2
//		l   m     r, smallest at right side, largest at right side
