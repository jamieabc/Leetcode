package main

// Given an unsorted array nums, reorder it in-place such that nums[0] <= nums[1] >= nums[2] <= nums[3]....
//
// Example:
//
// Input: nums = [3,5,2,1,6,4]
// Output: One possible answer is [3,5,1,6,2,4]

func wiggleSort(nums []int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		if (i&1 == 0 && nums[i] > nums[i-1]) || (i&1 == 1 && nums[i] < nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

//	problems
//	1.	too slow, because it's O(n log n), what I think of is to separate
//		numbers into a group of 3 elements

//		even start index, a <= b >= c, b is largest in 3, c is smallest in 3
//		odd start index, d >= e <= f, e is smallest in 3, f is largest in 3

//		but there might exist a problem, e.g. a, b, c, d
//		a <= b, b >= c, c >= d => switch c & d => a, b, d, c
//	2.	optimize, % takes divide operation, actually I only need a variable
//		that is flipping every round => bool
//	3.	it's slower...I should just check last bit
//	4.	refactor, what I do is swap, so two statements can be merged into one
