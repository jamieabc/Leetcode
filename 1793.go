package main

// You are given an array of integers nums (0-indexed) and an integer k.
//
// The score of a subarray (i, j) is defined as min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1). A good subarray is a subarray where i <= k <= j.
//
// Return the maximum possible score of a good subarray.
//
//
//
// Example 1:
//
// Input: nums = [1,4,3,7,4,5], k = 3
// Output: 15
// Explanation: The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15.
//
// Example 2:
//
// Input: nums = [5,5,4,5,4,1,1,1], k = 0
// Output: 20
// Explanation: The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.
//
//
//
// Constraints:
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 2 * 104
// 0 <= k < nums.length

func maximumScore(nums []int, k int) int {
	size := len(nums)

	// smallest from i ~ k if i <= k
	dp := make([]int, size)
	prev := nums[k]

	for i := k; i >= 0; i-- {
		prev = min(prev, nums[i])
		dp[i] = prev
	}

	// smallest from k ~ j if j >= k
	prev = nums[k+1]
	for i := k + 1; i < size; i++ {
		prev = min(prev, nums[i])
		dp[i] = prev
	}

	var ans int

	for i := 0; i <= k; i++ {
		smallest := dp[i]
		for j := k; j < size; j++ {
			smallest = min(smallest, dp[j])
			ans = max(ans, smallest*(j-i+1))
		}
	}

	return ans
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

//	Notes
//	1.	i way trying to minimize query of finding minimum numbers from i ~ j
//		i though it could be i ~ k ~ j
//		dp[a], a <= k finds minimum from a ~ k
//		dp[b], b > k finds minimum from k ~ b

//		e.g. nums = [1, 4, 3, 7, 4, 5], k = 3
//			   dp = [1, 3, 3, 7, 4, 5]
