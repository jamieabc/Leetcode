package main

import "sort"

// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
// Note:
//
// The solution set must not contain duplicate quadruplets.
//
// Example:
//
// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i <= size-4; i++ {
		for j := i + 1; j <= size-3; j++ {
			searchSum(nums, nums[i], nums[j], j+1, size-1, target-nums[i]-nums[j], &result)

			for j <= size-3 {
				if nums[j] == nums[j+1] {
					j++
				} else {
					break
				}
			}
		}

		for i <= size-4 {
			if nums[i] == nums[i+1] {
				i++
			} else {
				break
			}
		}
	}

	return result
}

func searchSum(nums []int, num1, num2, start, end, target int, result *[][]int) {
	for start < end {
		tmpSum := nums[start] + nums[end]

		if tmpSum == target {
			*result = append(*result, []int{num1, num2, nums[start], nums[end]})

			// find next non-duplicate number
			for start < len(nums)-1 && nums[start] == nums[start+1] {
				start++
			}
			start++

			for end > start && nums[end] == nums[end-1] {
				end--
			}
			end--
		} else if tmpSum > target {
			end--
		} else {
			start++
		}
	}
}

//	problems
//	1.	inspired from https://leetcode.com/problems/4sum/discuss/8565/Lower-bound-n3

//		minimum tc is O(n^3),
