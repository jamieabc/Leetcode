package main

//Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.
//
//You need to find the shortest such subarray and output its length.
//
//Example 1:
//
//Input: [2, 6, 4, 8, 10, 9, 15]
//Output: 5
//Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
//
//Note:
//
//    Then length of the input array is in range [1, 10,000].
//    The input array may contain duplicates, so ascending order here means <=.

// [2, 2, 1, 3, 4]
// [1, 2, 3, 4, 5]
// [1, 2, 4, 5, 3]
// [1, 3, 4, 2, 5]
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var left, right int

	for left = 0; left < len(nums)-1; left++ {
		// encounter problem
		if nums[left] > nums[left+1] {
			// if it's same number, traverse back to first number
			for left > 0 && nums[left] == nums[left-1] {
				left--
			}
			break
		}
	}

	for right = len(nums) - 1; right >= 1; right-- {
		// encounter problem
		if nums[right] < nums[right-1] {
			// if it's same number, traverse back to first number
			for right < len(nums)-1 && nums[right] == nums[right+1] {
				right++
			}
			break
		}
	}

	min := nums[left]
	max := nums[right]
	for i := left; i <= right; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	// for ascending, left-1 must be smaller than every number in range
	for left > 0 {
		if nums[left-1] > min {
			left--
		} else {
			break
		}
	}

	for right < len(nums)-1 {
		if nums[right+1] < max {
			right++
		} else {
			break
		}
	}

	if left > right {
		return 0
	}

	return right - left + 1
}
