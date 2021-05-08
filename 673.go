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

	lis, count := make([]int, size), make([]int, size)
	var largest, total int

	for i := range nums {
		lis[i], count[i] = 1, 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if lis[i] < lis[j]+1 {
					count[i] = count[j]
					lis[i] = lis[j] + 1
				} else if lis[i] == lis[j]+1 {
					count[i] += count[j]
				}
			}
		}

		if lis[i] > largest {
			largest = lis[i]
			total = count[i]
		} else if lis[i] == largest {
			total += count[i]
		}
	}

	return total
}

// tc: O(n^2)
func findNumberOfLIS2(nums []int) int {
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

//	3.	inspired from https://leetcode.com/problems/number-of-longest-increasing-subsequence/discuss/107318/C%2B%2B-DP-with-explanation-O(n2)

//		count[i] means # of LIS ending at index i
//		to calculate this value, consider following situation:

//		index j		i
//		if nums[j] < nums[i], LIS could be increased

//		under this condition,
//			if lis[j]+1 == lis[i], count[i] = count[i] + count[j]

//			if lis[j] <<< lis[j], count[i] = count[j]

//		finally, sum all count[i] where lis[i] is the longest

//	4.	inspired from https://leetcode.com/problems/number-of-longest-increasing-subsequence/discuss/107295/9ms-C%2B%2B-Explanation%3A-DP-%2B-Binary-search-%2B-prefix-sums-O(NlogN)-time-O(N)-space

//		inspired from comment https://leetcode.com/problems/number-of-longest-increasing-subsequence/discuss/107295/9ms-C++-Explanation:-DP-+-Binary-search-+-prefix-sums-O(NlogN)-time-O(N)-space/257188

//		binary search uses additional hash to store LIS ends at index i

//	 	the goal is that if a sequence can be enlarge to k, it must come from
//		all previous sequence with length k-1, the goal is to store that
//		information

//		by patience sort, the index found is the integer that can be replaced,
//		but for length, need to know which index is concat

//		e.g. 1, 5, 4, 6, 3, 5, 7

//		stack: 1
//		length	seq
//		1 		[0: 1] (sequence ends at index 0, w/ count 1)

//		stack: 1, 5
//		length	seq
//		1		[0: 1]
//		2		[5: 1]

//		stack: 1, 4 (5 is replaced)
//		length	seq
//		1		[0: 1]
//		2		[5: 1], [4: 1] (4 comes from 5)

//		stack: 1, 4, 6
//		length	seq
//		1		[0: 1]
//		2		[5: 1], [4: 1]
//		3		[6: 2] (6 > 5 & 6 > 4)

//		stack: 1, 3, 6	(4 is replaced by 3)
//		length	seq
//		1		[0: 1]
//		2		[5: 1], [4: 1], [3: 1]
//		3		[6: 2]

//		stack: 1, 3, 5	(6 is replaced by 5)
//		length	seq
//		1		[0: 1]
//		2		[5: 1], [4: 1], [3: 1]
//		3		[6: 2], [5: 2] (5 > 4 & 5 > 3, count = 2)

//		stack: 1, 3, 5, 7
//		length	seq
//		1		[0: 1]
//		2		[5: 1], [4: 1], [3: 1]
//		3		[6: 2], [5: 2]
//		4		[7: 4] (7 > 6 & 7 > 5)

//		from observation, subarray of seq forms a decreasing array of number,
//		thus binary search ban be used

//		length 4: 1, 3, 5, 7
//				  1, 4, 5, 7
//				  1, 4, 6, 7
//				  1, 5, 6, 7

//		this is too complicated, not implemented
