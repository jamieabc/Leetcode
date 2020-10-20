package main

import "math"

// Given an unsorted array of integers, find the length of longest increasing subsequence.
//
// Example:
//
// Input: [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//
// Note:
//
//     There may be more than one LIS combination, it is only necessary for you to return the length.
//     Your algorithm should run in O(n2) complexity.
//
// Follow up: Could you improve it to O(n log n) time complexity?

func lengthOfLIS(nums []int) int {
	// seq[i] means last number for longest increasing subsequence length i, each
	// lis sequence choose to have smallest last number
	// e.g. [1, 3, 5, 2, 8, 4, 6]
	// pick 1, length 1: [1]
	// pick 3, length 2: [1, 3], 3 > 1, increase length
	// pick 5, length 3: [1, 3, 5], 5 > 3, increase length
	// pick 2, length 2: [1, 2], 1 < 2 <= 3
	// pick 8, length 4: [1, 3, 5, 8], 8 > 5, increase length
	// pick 4, length 3: [1, 3, 4], 3 < 4 <= 5
	// pick 6, length 4: [1, 3, 5, 6], 5 < 6 <= 8
	// take last number from length 1, 2, 3, 4, sequence becomes 1, 2, 4, 6
	// which is a sorted array, can use binary search
	seq := make([]int, 0)

	for i := range nums {
		if len(seq) == 0 || nums[i] > seq[len(seq)-1] {
			// larger that existing largest, increase LIS length
			seq = append(seq, nums[i])
		} else {
			if nums[i] < seq[0] {
				// smaller than first number
				seq[0] = nums[i]
			} else {
				// binary search
				var tmp int
				for low, high := 0, len(seq)-1; low <= high; {
					mid := low + (high-low)/2

					if seq[mid] < nums[i] {
						low = mid + 1
					} else {
						tmp = mid
						high = mid - 1
					}
				}
				seq[tmp] = nums[i]
			}
		}
	}

	return len(seq)
}

// tc: O(n^2)
func lengthOfLIS4(nums []int) int {
	size := len(nums)

	// dp[i] means longest length of increasing seqeunce
	dp := make([]int, size)

	var longest int

	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		longest = max(longest, dp[i])
	}

	return longest
}

func lengthOfLIS3(nums []int) int {
	table := make([][]int, len(nums))
	for i := range table {
		table[i] = make([]int, len(nums))
		for j := range table {
			table[i][j] = -1
		}
	}
	return topDown(nums, len(nums)-1, math.MaxInt32, table)
}

// topDown(i) = 1 + max(topDown(j)), 0 < j < i
func topDown(nums []int, cur, prev int, table [][]int) int {
	// in case length == 0
	if cur < 0 {
		return 0
	}

	if prev != math.MaxInt32 && table[cur][prev] != -1 {
		return table[cur][prev]
	}

	var choose int
	if prev == math.MaxInt32 || nums[cur] < nums[prev] {
		choose = 1 + topDown(nums, cur-1, cur, table)
	}

	result := max(choose, topDown(nums, cur-1, prev, table))
	if prev != math.MaxInt32 {
		table[cur][prev] = result
	}

	return result
}

func lengthOfLIS2(nums []int) int {
	table := make([][]int, len(nums))
	for i := range table {
		table[i] = make([]int, len(nums))
		for j := range table {
			table[i][j] = -1
		}
	}

	return bottomUp(nums, 0, math.MinInt32, table)
}

func bottomUp(nums []int, cur, prev int, table [][]int) int {
	if len(nums) == cur {
		return 0
	}

	if prev >= 0 && table[prev][cur] != -1 {
		return table[prev][cur]
	}

	var choose int
	if prev < 0 || nums[cur] > nums[prev] {
		choose = 1 + bottomUp(nums, cur+1, cur, table)
	}

	result := max(choose, bottomUp(nums, cur+1, prev, table))

	if prev >= 0 && table[prev][cur] == -1 {
		table[prev][cur] = result
	}

	return result
}

func lengthOfLIS1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	// dp[i] - count for max seq length that is smaller than ith number
	dp := make([]int, size)
	maxL := 0

	for i := range nums {
		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		maxL = max(maxL, dp[i])
	}

	return maxL + 1
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//  problems
//  1.  longest subsequence, matters for choosing number. I cannot always
//      choose first number that is larger.
//      e.g. 1, 2, 100, 4, 5
//      longest is 1, 2, 4, 5, but if I choose 100, longest will be 1, 2, 100

//	2.	for longest increasing subsequence, it means a number has maximum
//		decreasing number ahead. so a problem is transformed into finding
//		maximum decreasing number ahead.

//		start from first number, whenever a number after is larger than self,
//		choose original value or self value +1.

//		tc: O(n^2)

//	3.	maximum subsequence may not alway at last item, need to update
//		for every number

//	4.	from solution, add brute force solution

//	5.	for dp, one important observation is that longest subsequence at
//		ith position is independent of later on positions

//	6.	inspired from https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/

//		bottom-up recursion

//	7.	add reference https://www.youtube.com/watch?v=S9oUiVYEq7E

//		tc: O(n log(n)), didn't implement it

//	8.	for any given position, longest subsequence comes before all number less
//		than self, so all previous calculated value can be reused

//	9.	inspired from https://leetcode.com/problems/longest-increasing-subsequence/discuss/74824/JavaPython-Binary-search-O(nlogn)-time-with-explanation

//		patience sorting https://www.cs.princeton.edu/courses/archive/spring13/cos423/lectures/LongestIncreasingSubsequence.pdf

//		patience sorting https://segmentfault.com/a/1190000003819886
