package main

//Given an array nums and an integer target.
//
//Return the maximum number of non-empty non-overlapping subarrays such that the sum of values in each subarray is equal to target.
//
//
//
//Example 1:
//
//Input: nums = [1,1,1,1,1], target = 2
//Output: 2
//Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals to target(2).
//
//Example 2:
//
//Input: nums = [-1,3,5,1,4,2,-9], target = 6
//Output: 2
//Explanation: There are 3 subarrays with sum equal to 6.
//([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.
//
//Example 3:
//
//Input: nums = [-2,6,6,3,5,4,1,2,8], target = 10
//Output: 3
//
//Example 4:
//
//Input: nums = [0,0,0], target = 0
//Output: 3
//
//
//
//Constraints:
//
//    1 <= nums.length <= 10^5
//    -10^4 <= nums[i] <= 10^4
//    0 <= target <= 10^6

func maxNonOverlapping(nums []int, target int) int {
	counter := make(map[int]int)
	counter[0] = 0

	var sum, count int
	end := -1

	for i := range nums {
		sum += nums[i]

		if val, ok := counter[sum-target]; ok && val > end {
			count++
			end = i
		}

		counter[sum] = i + 1
	}

	return count
}

//	problems
//	1.	since traversing is always increasing, no need to append all
//		index, just update latest one
