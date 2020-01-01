package main

//Given an array of integers nums, write a method that returns the "pivot" index of this array.
//
//We define the pivot index as the index where the sum of the numbers to the left of the index is equal to the sum of the numbers to the right of the index.
//
//If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.
//
//Example 1:
//
//Input:
//nums = [1, 7, 3, 6, 5, 6]
//Output: 3
//Explanation:
//The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
//Also, 3 is the first index where this occurs.
//
//
//
//Example 2:
//
//Input:
//nums = [1, 2, 3]
//Output: -1
//Explanation:
//There is no index that satisfies the conditions in the problem statement.
//
//
//
//Note:
//
//    The length of nums will be in the range [0, 10000].
//    Each element nums[i] will be an integer in the range [-1000, 1000].

func pivotIndex(nums []int) int {
	// start from left, right sum is total sum, then iterate all element from left, check if any equal
	length := len(nums)

	// no pivot
	if length == 0 {
		return -1
	}

	// 1
	if length == 1 {
		return 0
	}

	left := 0
	right := 0

	// get sum of right
	for i := 1; i < length; i++ {
		right += nums[i]
	}

	// check if pivot is index 0
	if left == right {
		return 0
	}

	for i := 1; i < length; i++ {
		left += nums[i-1]
		right -= nums[i]
		if left == right {
			return i
		}
	}

	return -1
}

// problems
// 1. when sum is smaller, check is next element makes which side larger/smaller
// 2. the condition choose 0 when there exist, because it won't change any result
// 3. some decision cannot be decided by only most recent element