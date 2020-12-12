package main

// Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
//
// Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.
//
// Clarification:
//
// Confused why the returned value is an integer, but your answer is an array?
//
// Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller.
//
// Internally you can think of this:
//
// // nums is passed in by reference. (i.e., without making a copy)
// int len = removeDuplicates(nums);
//
// // any modification to nums in your function would be known by the caller.
// // using the length returned by your function, it prints the first len elements.
// for (int i = 0; i < len; i++) {
//     print(nums[i]);
// }
//
//
//
// Example 1:
//
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3]
// Explanation: Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively. It doesn't matter what you leave beyond the returned length.
//
// Example 2:
//
// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3]
// Explanation: Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively. It doesn't matter what values are set beyond the returned length.
//
//
//
// Constraints:
//
//     0 <= nums.length <= 3 * 104
//     -104 <= nums[i] <= 104
//     nums is sorted in ascending order.

func removeDuplicates(nums []int) int {
	var store int

	for i := range nums {
		if i < 2 || nums[i] > nums[store-2] {
			nums[store] = nums[i]
			store++
		}
	}

	return store
}

func removeDuplicates1(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return size
	}

	var store, idx, count int

	for store, idx = 0, 0; idx < size; idx++ {
		count = 1

		for ; idx+count < size; count++ {
			if nums[idx+count] != nums[idx] {
				break
			}
		}

		for j := 0; j < min(2, count); store, j = store+1, j+1 {
			nums[store] = nums[idx+count-1]
		}

		idx += count - 1
	}

	return store
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	sorted means numbers with same value are grouped together

//	2.	inspired from https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/discuss/27976/3-6-easy-lines-C%2B%2B-Java-Python-Ruby

//		author provides great insight to this problem only variable need to know
//		is where to	store

//		it's a brilliant solution, at most k numbers means to check previous k
//		numbers (sliding window)
