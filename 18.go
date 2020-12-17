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
	sort.Ints(nums)
	ans := make([][]int, 0)
	size := len(nums)

	for a := 0; a < size-3; {
		for b := a + 1; b < size-2; {
			expected := target - nums[a] - nums[b]

			for c, d := b+1, size-1; c < size-1 && c < d; {
				tmp := nums[c] + nums[d]

				if tmp == expected {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})

					c++
					for c < size-1 && nums[c] == nums[c-1] {
						c++
					}

					d--
					for d > c && nums[d] == nums[d+1] {
						d--
					}
				} else if tmp < expected {
					c++
				} else {
					d--
				}
			}

			b++
			for b < size-2 && nums[b] == nums[b-1] {
				b++
			}
		}

		a++
		for a < size-3 && nums[a] == nums[a-1] {
			a++
		}
	}

	return ans
}

//	problems
//	1.	inspired from https://leetcode.com/problems/4sum/discuss/8565/Lower-bound-n3

//		minimum tc is O(n^3),
