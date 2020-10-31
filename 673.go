package main

// Given an integer array nums, return the number of longest increasing subsequences.
//
//
//
// Example 1:
//
// Input: nums = [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
//
// Example 2:
//
// Input: nums = [2,2,2,2,2]
// Output: 5
// Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
//
//
//
// Constraints:
//
//     0 <= nums.length <= 2000
//     -106 <= nums[i] <= 106

// tc: O(n^2)
func findNumberOfLIS(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}

	lis, cnt := make([]int, size), make([]int, size)
	var longest, count int
	for i := range nums {
		lis[i], cnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lis[i] < lis[j]+1 {
					lis[i] = lis[j] + 1
					cnt[i] = cnt[j]
				} else if lis[i] == lis[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}

		if lis[i] > longest {
			count = cnt[i]
			longest = lis[i]
		} else if lis[i] == longest {
			count += cnt[i]
		}
	}

	return count
}

// tc: O(n^3), n: size of nums
func findNumberOfLIS1(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}

	// dp[i][j] means number of increasing subsequence length i,
	// last number ends at index j
	dp := make([][]int, size+1)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	var longestCount, count int
	for i := range dp[0] {
		dp[1][i] = 1
	}
	longestCount = size

	for i := 2; i < size+1; i++ {
		count = 0
		for j := i - 1; j < size; j++ {
			for k := j - 1; k >= i-2; k-- {
				if nums[j] > nums[k] {
					dp[i][j] += dp[i-1][k]
				}
			}
			count += dp[i][j]
		}

		if count > 0 {
			longestCount = count
		}
	}

	return longestCount
}

//	Notes
//	1.	inspired from longest increasing subsequence (LIS, problem 300), to
//		find count of LIS, use another dp array to store information.

//		what's important is longest subsequence can be formed at index i, as
//		index increases, longest subsequence also with chance to increase

//		I have done LIS problem roughly 10 days ago, once this pattern is
//		identified, this problem can be solved, that's the power of pattern!

//	2.	inspired from https://leetcode.
//	com/problems/number-of-longest-increasing-subsequence/discuss/107295/9ms-C%2B%2B-Explanation%3A-DP-%2B-Binary-search-%2B-prefix-sums-O(NlogN)-time-O(N)-space

//		tc is O(n log(n)), but it's too complicated, not implement

//		basic idea evolves from LIS dp, dp[i][j] = [x, y], means LIS length i,
//		jth LIS ends at x with count y
