package main

// Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
//
// Example 1:
//
// Input:nums = [1,1,1], k = 2
// Output: 2
//
//
// Constraints:
//
// The length of the array is in range [1, 20,000].
// The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

func subarraySum(nums []int, k int) int {
	mapping := make(map[int]int)

	// sequence might starts from beginning, so 0 exist
	mapping[0] = 1

	var count, sum int
	for _, i := range nums {
		sum += i
		if s, ok := mapping[sum-k]; ok {
			count += s
		}
		mapping[sum]++
	}

	return count
}

func subarraySum1(nums []int, k int) int {
	// dp[i] means sum from index i to current position
	dp := make([]int, len(nums))
	var count int
	for i := range nums {
		for j := i; j >= 0; j-- {
			dp[j] += nums[i]
			if dp[j] == k {
				count++
			}
		}
	}

	return count
}
func subarraySum1(nums []int, k int) int {
	// dp[i] means sum from index i to current position
	dp := make([]int, len(nums))
	var count int
	for i := range nums {
		for j := i; j >= 0; j-- {
			dp[j] += nums[i]
			if dp[j] == k {
				count++
			}
		}
	}

	return count
}

//	problems
//	1.	too slow, inspired from https://leetcode.com/problems/subarray-sum-equals-k/discuss/190674/Python-O(n)-Based-on-%22running_sum%22-concept-of-%22Cracking-the-coding-interview%22-book

//	2.	it might exist multiple combinations to specific sum, use int instead
//		of bool
