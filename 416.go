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

// tc: O(mn), n: array size, m: subset of sums
func canPartition(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	// dp[i][j]: sum to j from index i
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < len(dp[0]); j++ {
			// not select
			dp[i][j] = dp[i-1][j]

			// select
			if j >= nums[i] {
				// very brilliant bottom-up dp
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[len(dp)-1][sum/2]
}

func canPartition4(nums []int) bool {
	var total int
	for _, n := range nums {
		total += n
	}

	if total&1 > 0 {
		return false
	}
	target := total >> 1

	size := len(nums)
	// memo[i][j]: start from index i, possible sum to target
	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = 0
		}
	}

	return dfs(nums, memo, 0, target)
}

func dfs(nums []int, memo [][]int, idx, target int) bool {
	if target == 0 {
		return true
	}

	if target < 0 || idx == len(nums) {
		return false
	}

	if memo[idx][target] != 0 {
		return memo[idx][target] == 1
	}

	possible := dfs(nums, memo, idx+1, target) || dfs(nums, memo, idx+1, target-nums[idx])

	if possible {
		memo[idx][target] = 1
	} else {
		memo[idx][target] = -1
	}

	return memo[idx][target] == 1
}

func canPartition3(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	// memo[i][j]: does start from index i possible sum to j
	// use array is more efficient, make([]int, target+1)
	memo := make(map[int]map[int]bool)

	return dfs3(nums, sum/2, 0, memo)
}

func dfs3(nums []int, sum, idx int, memo map[int]map[int]bool) bool {
	if sum == 0 {
		return true
	}

	if idx >= len(nums) {
		return false
	}

	if _, ok := memo[idx][sum]; ok {
		return memo[idx][sum]
	}

	var include, exclude bool
	if nums[idx] <= sum {
		include = dfs3(nums, sum-nums[idx], idx+1, memo)
	}
	exclude = dfs3(nums, sum, idx+1, memo)

	if _, ok := memo[idx]; !ok {
		memo[idx] = make(map[int]bool)
	}
	memo[idx][sum] = include || exclude

	return memo[idx][sum]
}

// this is same as using hashmap, find all existing possible sums, add next val
// and check if able to reach target
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

// tc: O(mn), n: array size, m: # of subset sums
func canPartition1(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}
	target := sum >> 1

	// memo[i][j]: does first i numbers can sum to j
	memo := make([][]bool, len(nums))
	for i := range memo {
		memo[i] = make([]bool, target+1)
		memo[i][0] = true
	}

	if nums[0] <= target {
		memo[0][nums[0]] = true
	}

	bfs(nums, target, 1, memo)

	return memo[len(memo)-1][target]
}

// this is like bottom-up dp, start from every number, build relationship
func bfs(nums []int, target, idx int, memo [][]bool) {
	if idx == len(nums) {
		return
	}

	for i := 0; i < len(memo[0]); i++ {
		if memo[idx-1][i] {
			// not select nums[idx]
			memo[idx][i] = true

			// select nums[idx]
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

//	5.	for each number, select or not select, which means it's a binary tree

//	6.	inspired from solution, most important observation is the graph of
//		(not) selecting a number, which forms a binary tree

//		by that, it can be solved by bfs/dfs/dp

//	7.	to use dp/dfs, most important thing is to find recurring sub-problem

//		e.g. [1, 1, 2, 4, 6], sum goal = 14/2 = 7

//		-> not take index 0: [1, 2, 4, 6] sum to 7

//			-> not take index 1: [2, 4, 6] sum to 7
//			-> take index 1    : [2, 4, 6] sum to 6

//		-> take index 0	   : [1, 2, 4, 6] sum to 6

//			-> not take index 1: [2, 4, 6] sum to 7
//			-> take index 1    : [2, 4, 6] sum to 6

//		[2, 4, 6] sum to 7/6 both occurs twice, recurring sub-problem
