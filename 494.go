package main

// You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.
//
// Find out how many ways to assign symbols to make sum of integers equal to target S.
//
// Example 1:
//
// Input: nums is [1, 1, 1, 1, 1], S is 3.
// Output: 5
// Explanation:
//
// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3
//
// There are 5 ways to assign symbols to make the sum of nums be target 3.
//
//
//
// Constraints:
//
//     The length of the given array is positive and will not exceed 20.
//     The sum of elements in the given array will not exceed 1000.
//     Your output answer is guaranteed to be fitted in a 32-bit integer.

func findTargetSumWays(nums []int, S int) int {
	// 1 - 2 + 3 - 4 = (1+3) - (2+4)
	// if sum of group1 = a, sum of group2 = b
	// a - b = S
	// a + b = total
	// a = (S+total)/2

	var sum int
	for _, n := range nums {
		sum += n
	}

	if (S+sum)&1 == 1 || sum < S {
		return 0
	}

	target := (S + sum) >> 1

	// dp[i][j] means count for first i number, sum to j
	dp := make([]int, target+1)
	dp[0] = 1

	for i := range nums {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[target]
}

func findTargetSumWays2(nums []int, S int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum < S {
		return 0
	}

	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, 2*(sum+1)+1)
		for j := range memo[0] {
			memo[i][j] = -1
		}
	}

	return permutation(nums, 0, sum, 0, S, memo)
}

func permutation(nums []int, idx, sum, cur, target int, memo [][]int) int {
	if idx == len(nums) {
		if cur == target {
			return 1
		}
		return 0
	}

	convertedIdx := abs(target - cur)
	if memo[idx][convertedIdx] != -1 {
		return memo[idx][convertedIdx]
	}

	plus := permutation(nums, idx+1, sum, cur+nums[idx], target, memo)
	minus := permutation(nums, idx+1, sum, cur-nums[idx], target, memo)

	memo[idx][convertedIdx] = plus + minus

	return memo[idx][convertedIdx]
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func findTargetSumWays1(nums []int, S int) int {
	var count int
	permutation(nums, 0, S, &count)

	return count
}

func permutation(nums []int, idx, target int, result *int) {
	if idx == len(nums) {
		if target == 0 {
			*result++
		}
		return
	}

	permutation(nums, idx+1, target+nums[idx], result)
	permutation(nums, idx+1, target-nums[idx], result)
}

//	problems
//	1.	at fist, I complicate the problem,  dp[i][j] means for index i, j is
//		summed value, which could range from -sum ~ sum, and since minus is
//		not able to represent in array, so I choose sum+1 as 0, -1 to be sum+1-1,
//		etc.

//		but if I store difference to target, then everything is positive

//	2.	intuition: with + & -, which means group numbers into two, one group
//		are positive, the other group is negative

//		e.g. 1 - 2 + 3 - 4 => (1+4) - (2+3)

//		this problem becomes separate number into two groups, difference of two
//		groups sum is some value

//		if sum of group1 is a, sum of group2 is b

//		a - b = S
//		a + b = total
//		a = (S+total)/2

//	3.	add reference https://leetcode.com/problems/target-sum/discuss/455024/DP-IS-EASY!-5-Steps-to-Think-Through-DP-Questions.

//		author provides a good explanation of thinking process
