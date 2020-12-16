package main

// Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it represented by array nums. You are asked to burst all the balloons. If the you burst balloon i you will get nums[left] * nums[i] * nums[right] coins. Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.
//
// Find the maximum coins you can collect by bursting the balloons wisely.
//
// Note:
//
// You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
// Example:
//
// Input: [3,1,5,8]
// Output: 167
// Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
// coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

func maxCoins(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	// dp[i][j]: max coins can get from i ~ j
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for d := 0; d < size; d++ {
		for L := 0; L+d < size; L++ {
			R := L + d

			for i := L; i <= R; i++ {
				var left, right int

				tmp := nums[i]
				if L > 0 {
					tmp *= nums[L-1]
				}

				if R < size-1 {
					tmp *= nums[R+1]
				}

				if i > L {
					left = dp[L][i-1]
				}

				if i+1 <= R {
					right = dp[i+1][R]
				}

				dp[L][R] = max(dp[L][R], tmp+left+right)
			}
		}
	}

	return dp[0][size-1]
}

// tc: O(n^3)
func maxCoins2(nums []int) int {
	size := len(nums)

	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, size)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dfs(nums, memo, 0, size-1)
}

func dfs(nums []int, memo [][]int, left, right int) int {
	size := len(nums)

	if left > right || left < 0 || right >= size {
		return 0
	}

	if memo[left][right] == -1 {
		var maximum int

		for i := left; i <= right; i++ {
			tmp := nums[i]

			// if within left boundary, the closest number
			if left > 0 {
				tmp *= nums[left-1]
			}

			// if with right boundary, the closest number
			if right < size-1 {
				tmp *= nums[right+1]
			}

			maximum = max(maximum, tmp+dfs(nums, memo, left, i-1)+dfs(nums, memo, i+1, right))
		}

		memo[left][right] = maximum
	}

	return memo[left][right]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func maxCoins1(nums []int) int {
	size := len(nums)
	used := make([]byte, size)
	for i := range used {
		used[i] = '0'
	}
	var maximum int

	backtracking(nums, used, 0, 0, &maximum)

	return maximum
}

// tc: O(n!)
func backtracking(nums []int, used []byte, length, current int, maximum *int) {
	size := len(nums)
	if size == length {
		*maximum = max(*maximum, current)
		return
	}

	for i := range used {
		if used[i] == '1' {
			continue
		}
		used[i] = '1'

		tmp := nums[i]
		for j := i - 1; j >= 0; j-- {
			if used[j] == '0' {
				tmp *= nums[j]
				break
			}
		}

		for j := i + 1; j < size; j++ {
			if used[j] == '0' {
				tmp *= nums[j]
				break
			}
		}

		backtracking(nums, used, length+1, current+tmp, maximum)

		used[i] = '0'
	}
}

//	Notes
//	1.	use backtracking, tc would be O(n!), too slow

//	2.	after a balloon is used, it has nothing to do with remaining balloons,
//		with the property, use 2D array to store maximum coins can get from range
//		i ~ j.

//	3.	For any balloon, if it's the first one to shoot, then coins can get
//		depends on what remaining balloons are which is complex.

//		The other way is to choose balloon as last one and increment it

//	4.	inspired from https://leetcode.com/problems/burst-balloons/discuss/76228/Share-some-analysis-and-explanations

//		use divide & conquer with tc O(n^3)

//	5.	inspired from https://leetcode.com/problems/burst-balloons/discuss/76229/For-anyone-that-is-still-confused-after-reading-all-kinds-of-explanations...

//		author reminds a good point: dp[i][j] means after burst all balloons in range
//		i ~ j

//	6.	inspired from https://leetcode.com/problems/burst-balloons/discuss/76241/Another-way-to-think-of-this-problem-(Matrix-chain-multiplication)

//		author provides another way of viewing this problem: matrix-chain multiplication
