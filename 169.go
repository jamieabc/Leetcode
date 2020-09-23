package main

import (
	"math"
	"sort"
)

// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
//
// You may assume that the array is non-empty and the majority element always exist in the array.
//
// Example 1:
//
// Input: [3,2,3]
// Output: 3
//
// Example 2:
//
// Input: [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
	var majority, count int

	for _, num := range nums {
		if count == 0 {
			majority = num
			count++
		} else {
			if num == majority {
				count++
			} else {
				count--
			}
		}
	}

	return majority
}

func majorityElement1(nums []int) int {
	result := divide(nums)

	for _, nums := range result {
		if nums[1] >= len(nums)/2 {
			return nums[0]
		}
	}

	return 0
}

func divide(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 1 {
		result = append(result, []int{nums[0], 1})
	} else if len(nums) == 2 {
		if nums[0] == nums[1] {
			result = append(result, []int{nums[0], 2})
		} else {
			result = append(result, []int{nums[0], 1})
			result = append(result, []int{nums[1], 1})
		}
	} else {
		r1 := divide(nums[:len(nums)/2])
		r2 := divide(nums[len(nums)/2:])

		// merge
		r := append([][]int{}, r1...)

		for i := range r2 {
			if r2[i][0] == r1[0][0] {
				r[0][1] += r2[i][1]
			} else if len(r1) > 1 && r2[i][0] == r1[1][0] {
				r[1][1] += r2[i][1]
			} else {
				r = append(r, r2[i])
			}
		}
		result = append(result, conquer(r)...)
	}

	return result
}

func conquer(r [][]int) [][]int {
	sort.Slice(r, func(i, j int) bool {
		return r[i][1] > r[j][1]
	})

	if len(r) >= 2 {
		return r[:2]
	}

	return r
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, the way to calculate divide & conquer tc can be
//		seen as this: T(n) = 2T(n/2) + 2n, which yields to O(n log n)

//	2.	inspired from solution, boyer-moore algorithm is really brilliant, the
//		premise of this algorithm to work correctly is that: majority element
//		exist. Majority is defined as n/2 + 1

//	3.	inspired from https://leetcode.com/problems/majority-element/discuss/51612/C%2B%2B-6-Solutions

//		bit manipulation can also be used, key idea is that majority number exist
//		also means for every bit, majority determiend the number
