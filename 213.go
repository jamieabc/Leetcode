package main

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given a list of non-negative integers nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
//
//
// Example 1:
//
// Input: nums = [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
//
// Example 2:
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
//
// Example 3:
//
// Input: nums = [0]
// Output: 0
//
//
//
// Constraints:
//
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 1000

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	} else if size == 2 {
		return max(nums[0], nums[1])
	}

	return max(robSimple(nums, 0, size-2), robSimple(nums, 1, size-1))
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func robSimple(nums []int, start, end int) int {
	prev1, prev2 := max(nums[start], nums[start+1]), nums[start]

	for i := start + 2; i <= end; i++ {
		prev2 = max(prev1, prev2+nums[i])
		prev1, prev2 = prev2, prev1
	}

	return max(prev1, prev2)
}

//	Notes
//	1.	I feel like two arrays are not necessary, can it be merged into one?

//		two arrays with different initial value, and different conditions for
//		last value

//	2.	inspired from solution, cannot reduce, but array is not necessary, only
//		to find max is sufficient
