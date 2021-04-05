package main

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// You may assume no duplicate exists in the array.
//
// Example 1:
//
// Input: [3,4,5,1,2]
// Output: 1
//
// Example 2:
//
// Input: [4,5,6,7,0,1,2]
// Output: 0

func findMin(nums []int) int {
	size := len(nums)

	smallest := nums[0]

	// becareful about size == 1, because in the binary search condition,
	// if only one element, it won't fit into any condition, thus cause infinite loop
	// or change nums[mid] < nums[0] to nums[mid] <= nums[0]
	if nums[size-1] > nums[0] || size == 1 {
		return smallest
	}

	for low, high := 0, size-1; low <= high; {
		mid := low + (high-low)/2

		if nums[mid] > nums[size-1] {
			// 3,4,5,1,2
			low = mid + 1
		} else if nums[mid] < nums[0] {
			// 5,1,2,3,4
			smallest = nums[mid]
			high = mid - 1
		}
	}

	return smallest
}

func findMin1(nums []int) int {
	size := len(nums)

	// no rotation
	if nums[size-1] > nums[0] {
		return nums[0]
	}

	// find pivot index
	low, high := 0, size-1

	for low < high {
		mid := low + (high-low)/2

		if nums[high] > nums[mid] {
			high = mid
		} else {
			low = mid + 1
		}

		// if condition changes to nums[low] < nums[mid], need to deal with
		//	additional condition: mid == low
	}

	return nums[low]
}

//	Notes
//	1.	return number instead of index

//	2.	inspired from https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/48484/A-concise-solution-with-proof-in-the-comment

//		I am confuse why condition can only be nums[high] > nums[mid], and
//		nums[low] < nums[mid] not working.

//		the reason is in comment, left < right (no duplicates), and mid =
//		low + (high-low)/2, so mid will never be high, thus reduce one condition
//		to deal with same index

//		and since mid never equal to high, which means mid = high will always
//		makes interval shrinking

//	3.	another explanation https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/158940/Beat-100%3A-Very-Simple-(Python)-Very-Detailed-Explanation

//		very brilliant solution

//	4.	inspired from solution, there's more elegant way to solve it
