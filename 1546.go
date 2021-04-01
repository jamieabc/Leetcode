package main

import "sort"

//Given an array nums and an integer target.
//
//Return the maximum number of non-empty non-overlapping subarrays such that the sum of values in each subarray is equal to target.
//
//
//
//Example 1:
//
//Input: nums = [1,1,1,1,1], target = 2
//Output: 2
//Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals to target(2).
//
//Example 2:
//
//Input: nums = [-1,3,5,1,4,2,-9], target = 6
//Output: 2
//Explanation: There are 3 subarrays with sum equal to 6.
//([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.
//
//Example 3:
//
//Input: nums = [-2,6,6,3,5,4,1,2,8], target = 10
//Output: 3
//
//Example 4:
//
//Input: nums = [0,0,0], target = 0
//Output: 3
//
//
//
//Constraints:
//
//    1 <= nums.length <= 10^5
//    -10^4 <= nums[i] <= 10^4
//    0 <= target <= 10^6

func maxNonOverlapping(nums []int, target int) int {
	counter := make(map[int]int)
	counter[0] = 0

	var sum, count int
	end := -1

	for i := range nums {
		sum += nums[i]

		if val, ok := counter[sum-target]; ok && val > end {
			count++
			end = i
		}

		counter[sum] = i + 1
	}

	return count
}

func maxNonOverlapping2(nums []int, target int) int {
	table := make(map[int]int)

	var sum, ans int
	prev := -1
	for i := range nums {
		if idx, ok := table[sum-target]; ok {
			if idx > prev {
				prev = i - 1
				ans++
			}
		}
		table[sum] = i
		sum += nums[i]
	}

	if idx, ok := table[sum-target]; ok {
		if idx > prev {
			ans++
		}
	}

	return ans
}

// TLE
func maxNonOverlapping1(nums []int, target int) int {
	intervals := make([][]int, 0)
	table := make(map[int][]int)

	var sum int
	for i := range nums {
		if arr, ok := table[sum-target]; ok {
			for j := range arr {
				intervals = append(intervals, []int{arr[j], i - 1})
			}
		}
		table[sum] = append(table[sum], i)
		sum += nums[i]
	}

	if arr, ok := table[sum-target]; ok {
		for i := range arr {
			intervals = append(intervals, []int{arr[i], len(nums) - 1})
		}
	}

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	prev := 0
	ans := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[prev][1] {
			prev = i
			ans++
		}
	}

	return ans
}

//	Notes
//	1.	since traversing is always increasing, no need to append all
//		index, just update latest one

//	2.	rewrite it, I thought it was about intervals, using prefix sum to
//		list all intervals that sums up to target.

//		with those intervals, it's about finding maximum non-overlapping
//		intervals, can be solved by sort by end time and greedy search

//		but listing all intervals is not necessary, because only care about
//		previous env time

//	3.	inspired from https://leetcode.com/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target/discuss/780882/Java-14-lines-Greedy-PrefixSum-with-line-by-line-explanation-easy-to-understand

//		my second attempt can be further improved by adding 1 to right-most index
