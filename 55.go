package main

import "math"

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
//
// Each element in the array represents your maximum jump length at that position.
//
// Determine if you are able to reach the last index.
//
//
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//
// Example 2:
//
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
//
//
//
// Constraints:
//
//     1 <= nums.length <= 3 * 10^4
//     0 <= nums[i][j] <= 10^5

func canJump(nums []int) bool {
	var i, farthest int
	for i, farthest = 0, 0; i < len(nums); i++ {
		if farthest < i {
			return false
		}
		farthest = max(farthest, i+nums[i])
	}

	return i == len(nums)
}

func canJump3(nums []int) bool {
	goal := len(nums) - 1

	for i := goal; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}

func canJump2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	var m, start, end int
	for end < len(nums)-1 {
		m = math.MinInt32
		for i := 0; start+i <= end; i++ {
			m = max(m, start+i+nums[start+i])
		}

		if m == end {
			return false
		}

		start, end = end, m

		if end >= len(nums)-1 {
			return true
		}
	}

	return end >= len(nums)-1
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func canJump1(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		if dp[i] {
			for j := 0; j <= nums[i]; j++ {
				if i+j == len(nums)-1 {
					return true
				} else if i+j < len(nums) {
					dp[i+j] = true
				}
			}
		}
	}

	return dp[len(nums)-1]
}

//	problems
//	1.	too slow, no need to traverse all numbers, I can use greedy way to
//		choose biggest number among jumping range

//	2.	use greedy algo to find maximum next jump in each jump

//	3.	boundary conditions when length == 1

//	4.	when jump is 0, then index is same

//	5.	wrong jump checking

//	6.	when doing recursive, it needs to check valid array index before
//		continue to looping, otherwise it terminates and return false

//	7.	add reference https://leetcode.com/problems/jump-game/discuss/20917/Linear-and-simple-solution-in-C%2B%2B

//		author uses more elegant way to traverse the jump, uses single
//		variable to store that largest value that can meet, and traverse
//		each index

//		if largest value < index, means it cannot go any further, return
//		false

//		it's really simple and elegant...

//	8.	add anoher reference https://leetcode.com/problems/jump-game/discuss/20907/1-6-lines-O(n)-time-O(1)-space

//		author provides another way to traverse backward, start from length-1
//		and if find any i + nums[i] >= goal, replace goal to i, then continue
//		to traverse backward

//	9.	wrong logic when traversing backward
