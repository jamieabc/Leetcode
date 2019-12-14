package main

//Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.
//
//Example 1:
//
//Input: [2,3,-2,4]
//Output: 6
//Explanation: [2,3] has the largest product 6.
//
//Example 2:
//
//Input: [-2,0,-1]
//Output: 0
//Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	maxToCurrent := nums[0]
	minToCurrent := nums[0]
	result := nums[0]
	var tmp int

	// divide every number in 2 categories, total max to current number,
	// total min to current number, the max result could come from
	// current number, prev total max * current number,
	// prev total min * current number
	for i := 1; i < length; i++ {
		tmp = maxToCurrent
		maxToCurrent = max(nums[i], max(maxToCurrent*nums[i], minToCurrent*nums[i]))
		minToCurrent = min(nums[i], min(tmp*nums[i], minToCurrent*nums[i]))
		if maxToCurrent > result {
			result = maxToCurrent
		}
	}

	return result
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
