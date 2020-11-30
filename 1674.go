package main

import "math"

//You are given an integer array nums of even length n and an integer limit. In one move, you can replace any integer from nums with another integer between 1 and limit, inclusive.
//
//The array nums is complementary if for all indices i (0-indexed), nums[i] + nums[n - 1 - i] equals the same number. For example, the array [1,2,3,4] is complementary because for all indices i, nums[i] + nums[n - 1 - i] = 5.
//
//Return the minimum number of moves required to make nums complementary.
//
//
//
//Example 1:
//
//Input: nums = [1,2,4,3], limit = 4
//Output: 1
//Explanation: In 1 move, you can change nums to [1,2,2,3] (underlined elements are changed).
//nums[0] + nums[3] = 1 + 3 = 4.
//nums[1] + nums[2] = 2 + 2 = 4.
//nums[2] + nums[1] = 2 + 2 = 4.
//nums[3] + nums[0] = 3 + 1 = 4.
//Therefore, nums[i] + nums[n-1-i] = 4 for every i, so nums is complementary.
//Example 2:
//
//Input: nums = [1,2,2,1], limit = 2
//Output: 2
//Explanation: In 2 moves, you can change nums to [2,2,2,2]. You cannot change any number to 3 since 3 > limit.
//Example 3:
//
//Input: nums = [1,2,1,2], limit = 2
//Output: 0
//Explanation: nums is already complementary.
//
//
//Constraints:
//
//n == nums.length
//2 <= n <= 105
//1 <= nums[i] <= limit <= 105
//n is even.

func minMoves(nums []int, limit int) int {
	size := len(nums)

	// dp[i]: needed moves for pair of sum to i
	dp := make([]int, limit*2+2)

	var m, n int
	for i := 0; i < size>>1; i++ {
		if nums[i] >= nums[size-1-i] {
			m, n = nums[i], nums[size-1-i]
		} else {
			m, n = nums[size-1-i], nums[i]
		}

		sum := m + n

		dp[2] += 2
		dp[n+1]--
		dp[sum]--
		dp[sum+1]++ // sum could be 2*limit, dp size should be increased
		dp[m+limit+1]++
	}

	minChange := math.MaxInt32
	var count int

	for i := 2; i <= limit*2; i++ {
		count += dp[i]
		if count < minChange {
			minChange = count
		}
	}

	return minChange
}

//	Notes
//	1.	1 <= number <= limit => 2 <= pair sum <= 2 * limit, for any given number,
//		max & min reachable numbers are fixed

//		e.g. number 3, limit = 5
//		another number can only ranges from 1 ~ 5, so pair sum for 3 is 4 ~ 8

//	2.	since possible sum is limited, it's possible to calculate moves for every
//		pair numbers, time complexity O(n*limit), n: size of array

//	3.	it's an accumulation process, store "events" to reduce computation
