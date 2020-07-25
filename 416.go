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

func canPartition3(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	// memo[i][j] means start from index i, sum to j is possible or not
	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, sum/2+1)
		memo[i][0] = 1

		for j := 1; j < len(memo[0]); j++ {
			memo[i][j] = -1
		}
	}

	return dfs(nums, sum/2, 0, memo)
}

func dfs(nums []int, sum, idx int, memo [][]int) bool {
	if sum == 0 {
		return true
	}

	if idx >= len(nums) {
		return false
	}

	if memo[idx][sum] > -1 {
		return memo[idx][sum] == 1
	}

	var include, exclude bool
	if nums[idx] <= sum {
		include = dfs(nums, sum-nums[idx], idx+1, memo)
	}
	exclude = dfs(nums, sum, idx+1, memo)

	if exclude || include {
		memo[idx][sum] = 1
	} else {
		memo[idx][sum] = 0
	}

	return memo[idx][sum] == 1
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
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	// memo[i][j] == 1 means firsts i numbers can sum to j
	memo := make([][]bool, len(nums))
	for i := range memo {
		memo[i] = make([]bool, sum/2+1)
		memo[i][0] = true
	}

	if nums[0] < sum/2 {
		memo[0][nums[0]] = true
	}

	bfs(nums, sum/2, 1, memo)

	return memo[len(memo)-1][sum/2]
}

func bfs(nums []int, target, idx int, memo [][]bool) {
	if idx >= len(nums) {
		return
	}

	for i := 0; i < len(memo[0]); i++ {
		if memo[idx-1][i] {
			memo[idx][i] = true

			if nums[idx]+i <= target {
				memo[idx][nums[idx]+i] = true
			}
		}
	}

	bfs(nums, target, idx+1, memo)
}

//	problems
//	1.	too slow

//	2.	inspired from https://leetcode.com/problems/partition-equal-subset-sum/discuss/90592/01-knapsack-detailed-explanation

//		intuition: 2 subsets sum equal => each subset sums to (total sum)/2

//		intuition: if total sum is odd number, can never meet since /2 will
//		never separate evenly

//	3.	inspired from https://leetcode.com/problems/partition-equal-subset-sum/discuss/462699/Whiteboard-Editorial.-All-Approaches-explained.

//		very important summary

//	4.	for dfs, the goal is to find all possible values of subarray sum
