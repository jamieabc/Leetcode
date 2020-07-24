package main

// Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
//
// Note:
//
// Each of the array element will not exceed 100.
// The array size will not exceed 200.
//
//
//
// Example 1:
//
// Input: [1, 5, 11, 5]
//
// Output: true
//
// Explanation: The array can be partitioned as [1, 5, 5] and [11].
//
//
//
// Example 2:
//
// Input: [1, 2, 3, 5]
//
// Output: false
//
// Explanation: The array cannot be partitioned into equal sum subsets.
func canPartition(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	// dp[i][j] means sum j from index i items
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j]

			if j >= nums[i] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[len(dp)-1][sum/2]
}

func canPartition2(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	candidates := make(map[int]bool)
	candidates[0] = true
	comb(nums, candidates)

	return candidates[sum/2]
}

func comb(nums []int, candidates map[int]bool) {
	for i := range nums {
		keys := make([]int, 0)

		for key := range candidates {
			keys = append(keys, key)
		}

		for j := range keys {
			candidates[keys[j]+nums[i]] = true
		}
	}
}

func canPartition1(nums []int) bool {
	return dfs(nums, 0, 0, 0)
}

func dfs(nums []int, sum1, sum2, idx int) bool {
	if idx < len(nums) {
		return dfs(nums, sum1+nums[idx], sum2, idx+1) || dfs(nums, sum1, sum2+nums[idx], idx+1)
	}

	return sum1 == sum2
}

//	problems
//	1.	too slow

//	2.	inspired from https://leetcode.com/problems/partition-equal-subset-sum/discuss/90592/01-knapsack-detailed-explanation

//		intuition: 2 subsets sum equal => each subset sums to (total sum)/2

//		intuition: if total sum is odd number, can never meet since /2 will
//		never separate evenly

//	3.	inspired from https://leetcode.com/problems/partition-equal-subset-sum/discuss/462699/Whiteboard-Editorial.-All-Approaches-explained.

//		very important summary
