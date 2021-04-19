package main

// Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.
//
// The answer is guaranteed to fit in a 32-bit integer.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3], target = 4
// Output: 7
// Explanation:
// The possible combination ways are:
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// Note that different sequences are counted as different combinations.
//
// Example 2:
//
// Input: nums = [9], target = 3
// Output: 0
//
//
//
// Constraints:
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 1000
// All the elements of nums are unique.
// 1 <= target <= 1000
//
//
//
// Follow up: What if negative numbers are allowed in the given array? How does it change the problem? What limitation we need to add to the question to allow negative numbers?

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)

	if target < nums[0] {
		return 0
	}

	size := len(nums)
	memo := make([]int, max(target, nums[size-1])+1)
	for i := range memo {
		memo[i] = -1
	}

	return dfs(nums, target, memo)
}

func dfs(nums []int, target int, memo []int) int {
	if target == 0 {
		return 1
	}

	if target < 0 {
		return 0
	}

	if memo[target] == -1 {
		var count int

		for i := 0; i < len(nums) && nums[i] <= target; i++ {
			count += dfs(nums, target-nums[i], memo)
		}
		memo[target] = count
	}

	return memo[target]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	becareful about condition, although it says combination, but array
//		[1, 3] != [3, 1], it's more like permutation

//	2.	since every number is distinct, it further reduced the time complexity
//		distinct numbers makes count easier

//		[1, 2, 3], target = 4
//		first round: 1, target = 3
//					 2, target = 2
//					 3, target = 1

//		second round: 1, 1, target = 2
//					  1, 2, target = 1
//					  1, 3, target = 0

//					  2, 1 target = 1
//					  2, 2 target = 0

//					  3, 1 target = 0

//		since order matters, it's a recursion that each time reduce some value
//		and keep going

// 	2.	inspired from solution, originally I thought was O(n^2),
// 		n: target/smallest_number, but this is not true

// 		because each possible outcome will only be computed once, there are
// 		n numbers, and all possible sum will only range from 1 ~ target,
// 		overall tc will be O(target * n)

//	3.	this problem is dp, not backtracking

//	4.	for the follow-up, since sum is not always increasing, it could be
//		infinite sequence
